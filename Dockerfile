FROM golang:1.20.10-alpine3.17

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /build ./cmd/main.go

EXPOSE 8781

CMD [ "/build" ]
