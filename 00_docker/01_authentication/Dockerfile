FROM python:3-alpine

RUN apk add --no-cache git
RUN pip install --upgrade pip
COPY . /jobs
WORKDIR /jobs
RUN pip install -r  requirements.txt
EXPOSE 5000
