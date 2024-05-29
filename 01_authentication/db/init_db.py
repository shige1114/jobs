from db import Base,engine
from db.user import User
def init_db():
    Base.metadata.create_all(bind=engine)

def reset_db():
    Base.metadata.drop_all(bind=engine)
    init_db()


if __name__ == "__main__":
    reset_db()