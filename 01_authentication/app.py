from flask import Flask,request,make_response,jsonify
from sqlalchemy.exc import DataError,ProgrammingError,IntegrityError
from db.user import UserApi
from authenticator import get_access_token

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route("/authentication",methods=["GET"])
def signIn():
    """
    ユーザのサインインを行う処理
    """
    
    email = request.args.get("email",'')        
    print(email)
    password = request.args.get("password","")
    print(password)
    try:
        user = UserApi.get(email=email,password=password)
        if user == None:
            return "bad password ",404 
    except Exception as e:
        print(e)
        return "bad email or password",404
    
    return jsonify({"jwt":get_access_token(user.id)}),200

@app.route("/authentication",methods=["POST"])
def singUp():
    """
    ユーザのサインアップを行う処理
    """
    try:
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

@app.route("/authentication",methods=["PUT"])
def changeInfo():
    return
if __name__ == "__main__":
    app.run(host="0.0.0.0",port=1000,debug=True)