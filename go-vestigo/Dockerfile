FROM golang:1.18

WORKDIR /myapp

COPY . /myapp

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" .

FROM scratch

COPY --from=0 /myapp/server /go/bin/server

ENTRYPOINT ["/go/bin/server"]
