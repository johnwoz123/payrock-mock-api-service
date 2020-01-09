FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/user-auth-service

COPY . .
RUN go mod download

RUN CGO_ENABLED=0  go build -a -installsuffix cgo -o user-auth-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
#COPY --from=builder /go/src/github.com/johnwoz123/payrock-mock-api-service/user-auth-service/user-auth-service .
COPY --from=builder /app/user-auth-service/user-auth-service .

CMD ["./user-auth-service"]
