############################
# STEP 1 build executable binary
# This is an intermediate docker image, will be deleted.
############################

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ADD . $GOPATH/src/app/
WORKDIR $GOPATH/src/app/

# Fetch dependencies.
# Using go get.
RUN go mod tidy
# Build the binary.
RUN go build -o /go/bin/echo_server

############################
# STEP 2 build a small image
############################

FROM alpine

RUN mkdir /app
WORKDIR /app

# Copy our static executable.
COPY --from=builder /go/bin/echo_server /app/echo_server

# Add app user
RUN addgroup -S app && adduser -S -G app app

# Change to app user
RUN chown -R app:app /app/

# Run as app user
USER app

EXPOSE 1323

CMD ["./echo_server"]
