# Build 
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /gojo
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w" -gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -o main .

# Run 
FROM alpine:3.18
WORKDIR /gojo
COPY --from=builder /gojo/main .
COPY gojo.env .
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/gojo/main" ]