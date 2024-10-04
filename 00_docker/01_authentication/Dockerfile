FROM python:3-alpine

RUN apk add --no-cache git
RUN pip install --upgrade pip
COPY . .
RUN pip install -r  requirements.txt
EXPOSE 5000

CMD [ "python","app.py" ]