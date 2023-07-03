FROM golang:1.20

WORKDIR /SHIFTSYNC

COPY . .

RUN make build

CMD make run