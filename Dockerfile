FROM golang:1.15.2-alpine3.12

ARG SERVICE_NAME

WORKDIR /app

COPY . .

RUN go build -o server ./${SERVICE_NAME}/server.go

CMD [ "/app/server" ]
