import unittest
from authenticator import get_access_token,hash_password,verify_password

class AuthTest(unittest.TestCase):
    
    def test(self):
        passw = hash_password("pass")

        print(passw)

        print("suc:{}".format(verify_password("pass","pass")))
        print("fail:{}".format(verify_password("pass","pas")))
        
    def test2(self):
        jwt = get_access_token("803ad68c-6a43-4636-b457-4abfba44574c")

        print(jwt)