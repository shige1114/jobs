FROM python:3-alpine

RUN apk add --no-cache git
RUN pip install --upgrade pip
COPY . .
RUN pip install -r  requirements.txt
EXPOSE 5000
CMD [ "python","main.py" ]


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
