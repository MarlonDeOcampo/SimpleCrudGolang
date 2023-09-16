FROM golang:1.18.1

WORKDIR /app

COPY . /app/

COPY .env .

RUN go mod tidy

RUN go build -o main .

CMD [ "/app/main" ]

EXPOSE 4005

