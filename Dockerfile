FROM golang:1.20

WORKDIR /SHIFTSYNC

COPY . .

RUN make build

RUN make test

CMD make run