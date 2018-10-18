FROM golang:1.11.1

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/midorigreen/sample/Go/linebot-pwg

COPY ./ ./
RUN dep ensure

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./pwg-server .

FROM alpine:latest

WORKDIR /opt/
COPY --from=0 /go/src/github.com/midorigreen/sample/Go/linebot-pwg/pwg-server .

ENTRYPOINT [ "./pwg-server" ]
CMD ["-p", "80"]
