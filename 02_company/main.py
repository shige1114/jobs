from repository.init_db import reset_db,init_db
from flask import Flask,request,make_response,jsonify
from sqlalchemy.exc import DataError,ProgrammingError,IntegrityError
from repository.company import CompanyRepository

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route("/company",methods=["GET"])
def get_company():
    """
    """
    name = request.args.get("name","")
    kana = request.args.get("name","")

    try:
        companies = CompanyRepository.get(name,kana)
    except Exception as e:
        print(e)
        return "bad email or password",404
    
    if not companies:
        return jsonify({"companies":[]})
    
    else:
        return jsonify({"companies": [company.to_dict() for company in companies]}),200



@app.route("/company_id", methods=["GET"])
def get_user():
    """
    ユーザが存在するか確認する処理
    """
    company_id = request.args.get("companyId", '')

    try:
        company = CompanyRepository.get_by_id(company_id)
    except Exception as e:
        print(e)
        return "Internal Server Error", 500
    
    if company:
        return jsonify({"message": "compnay exists"}), 200
    else:
        return jsonify({"message": "company not found"}), 404

@app.route("/authentication",methods=["PUT"])
def changeInfo():
    return
if __name__ == "__main__":
    app.run(host="0.0.0.0",port=1010,debug=True)