from sqlalchemy import create_engine,text
from sqlalchemy.orm import declarative_base, sessionmaker,scoped_session
import os
Base = declarative_base()

engine = create_engine(
    "postgresql://{user}:{password}@{host}/{dbname}".format(**{
        'user': os.getenv('POSTGRES_USER', 'postgres'),
        'password': os.getenv('POSTGRES_PASSWORD', 'password'),
        'host': os.getenv('POSTGRES_HOST', ''),
        'dbname':os.getenv('POSTGRES_DB',"")
        }),echo=True
)
session_factory = sessionmaker(
    bind=engine,
    autoflush=True,
    autocommit=False
)
Session = scoped_session(session_factory)


