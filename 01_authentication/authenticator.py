from cryptography.hazmat.primitives.asymmetric import rsa
import datetime
import os
import base64
import json
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric import padding

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
def _base64_to_json(b64):
    decoded_bytes = base64.urlsafe_b64decode(b64 + '==')
    decoded_str = decoded_bytes.decode('utf-8')
    json_obj = json.loads(decoded_str)
    return json_obj
    
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


def get_access_token(id: str, private_key) -> str:
    """
    JWTの作成 (RSAによる署名)
    """
    exp = datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(hours=EXP_TIME)
    header = {
        "alg": "RS256",  # RSA + SHA256アルゴリズム
        "typ": "jwt"
    }
    payload = {
        "iss": os.getenv("HOST"),  # 発行者
        "sub": id,  # 識別子
        "exp": exp.isoformat()  # 有効期限
    }

    header_base64 = _json_to_base64(header)
    payload_base64 = _json_to_base64(payload)

    # トークンヘッダーとペイロードを署名
    message = f"{header_base64}.{payload_base64}".encode("utf-8")
    signature = private_key.sign(
        message,
        padding.PKCS1v15(),  # 署名のためのパディング
        hashes.SHA256()  # SHA256で署名
    )
    
    # トークンを作成
    signature_base64 = base64.urlsafe_b64encode(signature).decode('utf-8')
    return f"{header_base64}.{payload_base64}.{signature_base64}"


def verify_jwt(token: str, public_key) -> bool:
    """
    JWTトークンの署名検証
    """
    try:
        # トークンを3つに分割 (header, payload, signature)
        header_base64, payload_base64, signature_base64 = token.split('.')

        # 署名のデコード
        signature = base64.urlsafe_b64decode(signature_base64 + '==')  # JWT署名はURLセーフエンコード

        # 検証用のメッセージ
        message = f"{header_base64}.{payload_base64}".encode("utf-8")

        # 公開鍵で署名を検証
        public_key.verify(
            signature,
            message,
            padding.PKCS1v15(),  # 署名のためのパディング
            hashes.SHA256()  # 使用されたハッシュアルゴリズム
        )
        # ペイロードをデコードしてJSONに変換
        payload_json = base64.urlsafe_b64decode(payload_base64 + '==').decode('utf-8')
        payload = json.loads(payload_json)

        # ユーザーIDを抽出
        user_id = payload.get("sub")

        if user_id:
            return user_id
        else:
            raise ValueError("UserID not found in payload")

    except Exception as e:
        print("Invalid token:", e)
        return None