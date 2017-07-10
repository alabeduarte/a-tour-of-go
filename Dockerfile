FROM golang:1.8

COPY . /app
WORKDIR /app

CMD ["go", "help"]
