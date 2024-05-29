import os
import base64
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes,serialization
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric import padding
import datetime
import json

EXP_TIME = 1
def _pbkd2hmac(salt:str):
    return PBKDF2HMAC(
        algorithm=hashes.SHA256(),
        length=32,
        salt=salt,
        iterations=100,
        backend=default_backend()
    )

def _json_to_base64(jsn):
    jsn_str = json.dumps(jsn)
    jsn_byte = jsn_str.encode()
    jsn_base64 = base64.urlsafe_b64encode(jsn_byte)

    return jsn_base64.decode() 

def hash_password(password:str) -> str:
    """
    パスワードのハッシュ化とソルトのエンコード
    """
    salt = os.urandom(16)
    kdf = _pbkd2hmac(salt)
    key = kdf.derive(password.encode())

    return base64.urlsafe_b64encode(salt+key).decode()

def verify_password(stored_password:str,password:str) -> bool:
    """
    パスワードのチェック
    """
    decode = base64.urlsafe_b64decode(stored_password.encode())
    salt = decode[:16]
    stored_key = decode[16:]
    kdf = _pbkd2hmac(salt)
    try:
        kdf.verify(password.encode(),stored_key)
        return True
    except:
        print("password error")
        return False


def get_access_token(id:str) -> str:
    """
    JWTの作成
    """
    with open("./public_key.pem", "rb") as key_file:
        public_key = serialization.load_pem_public_key(
            key_file.read(),
            backend=default_backend()
        )
    exp = datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(hours=EXP_TIME)
    header = {
        "alg":"SHA256",
        "typ":"jwt"
    }
    payload = {
        "iss":os.getenv("HOST"), #発行サービス
        "sub":id, #識別子
        "exp":exp.__str__()#有効期限
    }
    header_base64 = _json_to_base64(header)
    payload_base64 = _json_to_base64(payload)
    
    signature = public_key.encrypt(
        f"{header_base64}.{payload_base64}".encode("utf-8"),
        padding.OAEP(
            mgf=padding.MGF1(algorithm=hashes.SHA256()),
            algorithm=hashes.SHA256(),
            label=None
        )
    ) 
    return "{}.{}.{}".format(header_base64,payload_base64,signature)