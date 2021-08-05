FROM golang:1.15.13-alpine3.12 as builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE on
ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk update \
    && apk add git \
    gcc \
    musl-dev \
    upx \
    && apk add --no-cache ca-certificates && update-ca-certificates \
    && go get -u github.com/cosmtrek/air@latest \
    && go get -u bitbucket.org/liamstask/goose/cmd/goose \
    && go get github.com/pwaller/goupx \
    && go get github.com/go-bindata/go-bindata/... \
    && go get github.com/elazarl/go-bindata-assetfs/...

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY ./app ${ROOT}

EXPOSE 8080
