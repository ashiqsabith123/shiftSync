FROM golang:1.20 AS builder

WORKDIR /SHIFTSYNC

COPY . .

RUN make build

# Second stage for running the application
FROM alpine:latest

WORKDIR /SHIFTSYNC

COPY --from=builder /SHIFTSYNC .

CMD make run
