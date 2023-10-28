# Build
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /gojo
COPY . .

# The following line compiles your Go application with specific build flags.
RUN GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w" -gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -o main .

# Run
FROM alpine:3.18
WORKDIR /gojo

# Copy the compiled binary from the builder stage to the final image.
COPY --from=builder /gojo/main .

# Copy the environment file.
COPY gojo.env .

# Specify the command to run your application when the container starts.
CMD [ "./main" ]