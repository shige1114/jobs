import unittest
import sys
from repository.company import CompanyRepository
from pathlib import Path

class TestCompany(unittest.TestCase):
    def test_insert(self):
        print("__________________________")
        print("_________insert_________")
        try:
            CompanyRepository.create("test","test1","http:")
        except Exception as e:
            print(e)
        t =CompanyRepository.get("test","test")
        
        print("__________________________")
        print("test:{}".format(t))
        print("__________________________")
        pass

        

       


if __name__ == "__main__":

    unittest.main()