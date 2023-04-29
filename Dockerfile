FROM golang:alpine


WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o up-echo

EXPOSE 3001

ENTRYPOINT ["/app/up-echo"]