from sqlalchemy import Column,String,Date
from db import engine,Base
from datetime import datetime 
from db import Session
from authenticator import hash_password,verify_password
from uuid import uuid4
from sqlalchemy.exc import IntegrityError

class User(Base):
    __tablename__ = 'users'
    id = Column(String,primary_key=True)
    user_name = Column(String)
    email = Column(String,nullable=False,unique=True)
    password = Column(String,nullable=False)
    created_date = Column(Date,nullable=False)

class UserApi():
    
    def create(email,password):
        new_user = User(
            id=uuid4(),
            created_date=datetime.today(),
            email=email,
            password=hash_password(password)
        )
        with Session() as s :
            try:
                s.add(new_user)
                s.commit()
            except IntegrityError as e:
                s.rollback()  # ロールバック
                print("invalid email")
                raise e  # IntegrityErrorを再度raiseしてエラーを返す
            except Exception as e:
                print(e)
                raise Exception("Failed to add user") from e

    def get(email,password) -> User:
        with Session() as s:
            user = s.query(User).filter_by(email=email).first()
            
            if user and verify_password(user.password,password):
                return user
            return None
            
    def put(id):


        pass