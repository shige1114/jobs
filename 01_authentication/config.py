class Config(object):
    TESTING = False
    

class DevConfig(Config):
    DATABASE_URI = ""

class TestConfig(Config):
    DATABASE_URI = ""
    TESTING = True