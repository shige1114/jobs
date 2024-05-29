import unittest
from db.user import UserApi


class TestUserApi(unittest.TestCase):
    

    def test_insert(self):
        email = "example@gmail.com"
        password = "pass"
        print("__________________________")
        print("_________insert_________")
        try:
            UserApi.create(email,password)
        except Exception as e:
            print(e)
        user  = UserApi.get(email,password)
        
        print(user)
        pass

    def test_get(self):
        print("__________________________")
        print("_________select_________")
        email = "example@gmail.com"
        password = "pass"

        user = UserApi.get(email=email,password=password)
        print("パスワード認証成功")
        print(user)
        user = UserApi.get(email=email,password="pas")
        print("パスワード認証失敗")
        print(user)
        
        print("nothing email")
        user = UserApi.get(email="g@gmail.com",password="pas")
        print(user)

        

       


if __name__ == "__main__":
    unittest.main()