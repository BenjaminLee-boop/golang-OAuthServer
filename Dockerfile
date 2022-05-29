FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./src .

RUN go build -o /authServer

EXPOSE 9001

CMD [ "/authServer" ]