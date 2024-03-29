FROM golang:1.18-alpine AS buildStage

WORKDIR /usr/app

# Update all packages and add make
RUN apk add --update make

# Copy everything in the directory
COPY . .

# Build the application using makefile
RUN make build



FROM alpine:latest

WORKDIR /usr/app

COPY --from=buildStage /usr/app/bin/main .
COPY .env .

ENV GIN_MODE=release

EXPOSE 3000
CMD ["/usr/app/main"]
