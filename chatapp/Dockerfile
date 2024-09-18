FROM golang:alpine 

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o chatapp 

EXPOSE 8080

ENTRYPOINT [ "./chatapp" ]
