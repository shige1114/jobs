FROM python:3-alpine

RUN apk add --no-cache git
RUN pip install --upgrade pip
COPY . /app
WORKDIR /app
RUN pip install -r  requirements.txt
EXPOSE 5000
