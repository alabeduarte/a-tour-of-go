FROM golang:1.8

COPY . /app
WORKDIR /app

RUN go get golang.org/x/tour/gotour
CMD ["go", "run", "./moretypes/exercises/exercise-maps.go"]
