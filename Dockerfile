FROM golang:1.19

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o GP2K5-api

CMD ["./GP2K5-api"]