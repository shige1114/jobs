from sqlalchemy import Column,String,or_,Boolean
from repository import engine,Base,Session
from datetime import datetime 
from uuid import uuid4
from sqlalchemy.exc import IntegrityError

class Company(Base):
    __tablename__ = 'companies'
    id = Column(String,primary_key=True)
    name = Column(String,nullable=False,unique=True)
    kana = Column(String,nullable=False)
    url = Column(String,nullable=True)
    active = Column(Boolean,default=True)

class CompanyRepository():
    
    def create(name,kana,url):
        company = Company(
            id=uuid4(),
            name=name,
            kana=kana,
            url=url
        )
        with Session() as s :
            try:
                s.add(company)
                s.commit()
            except IntegrityError as e:
                s.rollback()  # ロールバック
                print("invalid email")
                raise e  # IntegrityErrorを再度raiseしてエラーを返す
            except Exception as e:
                print(e)
                raise Exception("Failed to add user") from e

    def get(name,kana) -> Company:
        pattern = "%{}%".format(name)
        kana_pattern = "%{}%".format(kana)
        with Session() as s:
            company = s.query(Company).filter(or_(Company.name.like(pattern),Company.kana.like(kana_pattern))).all()
            return company
            
    def put(id):


        pass