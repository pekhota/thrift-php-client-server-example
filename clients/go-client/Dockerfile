FROM golang:1.16-alpine

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build

CMD exec /bin/sh -c "trap : TERM INT; sleep infinity & wait"
#CMD ["./go-client"]