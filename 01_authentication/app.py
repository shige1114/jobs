from flask import Flask,request,make_response,jsonify
from sqlalchemy.exc import DataError,ProgrammingError,IntegrityError
from db.user import UserApi
from authenticator import get_access_token,verify_jwt
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.backends import default_backend
app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route("/login", methods=["POST"])
def signIn():
    """
    ユーザのサインインを行う処理
    """
    email = request.json.get("email", '')        
    print(email)
    password = request.json.get("password", "")
    print(password)
    
    try:
        # ユーザーを取得
        user = UserApi.get(email=email, password=password)
        if user is None:
            return jsonify({"error": "Invalid email or password"}), 404
        
        # 秘密鍵を読み込む
        with open("./private_key.pem", "rb") as key_file:
            private_key = serialization.load_pem_private_key(
                key_file.read(),
                password=None,
                backend=default_backend()
            )
        
        # JWTを作成
        jwt_token = get_access_token(user.id, private_key)

        # JWTを返す
        return jsonify({"jwt": jwt_token}), 200
    
    except Exception as e:
        print(f"Error occurred: {e}")
        return jsonify({"error": "An error occurred during login"}), 500


@app.route("/authentication",methods=["POST"])
def singUp():
    """
    ユーザのサインアップを行う処理
    """
    try:
        print(request.form)
        UserApi.create(**request.form)
    except IntegrityError as e:
        return "invalid email",404
    except DataError as e:
        return "Invalid value",404
    except ProgrammingError as e:
        return "Invalid Error",404
    except Exception as e:
        print(e)
        return "Error",404
    user = UserApi.get(email=request.form["email"],password=request.form["password"])
    print("User:{}".format(user))
    return jsonify({"jwt":get_access_token(user.id)}),200

@app.route("/authentication",methods=["DELETE"])
def signOut():
    return

@app.route("/user_id", methods=["GET"])
def get_user():
    """
    ユーザが存在するか確認する処理
    """
    user_id = request.args.get("userId", '')
    if not user_id:
        return "User ID is required", 400

    try:
        user = UserApi.get_by_id(user_id)
    except Exception as e:
        print(e)
        return "Internal Server Error", 500
    
    if user:
        return jsonify({"message": "User exists"}), 200
    else:
        return jsonify({"message": "User not found"}), 404

@app.route("/verify_token", methods=["POST"])
def verify_token():
    token = request.json.get("token")
    if not token:
        return jsonify({"error": "Token is missing"}), 400

    try:
        with open("./public_key.pem", "rb") as key_file:
            public_key = serialization.load_pem_public_key(
                key_file.read(),
                backend=default_backend()
            )
        payload = verify_jwt(token, public_key)
        return jsonify({"user_id":payload}), 200
    except Exception as e:
        print(f"Error occurred: {e}")
        return jsonify({"error": "Invalid token"}), 400
    

@app.route("/authentication",methods=["PUT"])
def changeInfo():
    return
if __name__ == "__main__":
    app.run(host="0.0.0.0",port=1000,debug=True)