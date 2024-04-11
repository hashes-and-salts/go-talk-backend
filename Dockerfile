FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .

RUN go build -o main .
CMD [ "./main" ]