FROM golang:1.17-alpine AS builder
WORKDIR $GOPATH/src/github.com/gin-html-website
COPY . .
RUN apk add --no-cache git ca-certificates g++
RUN GOOS=linux GOARCH=arm64 go build -ldflags '-extldflags "-static"' -o website
RUN mkdir requirements \
  && cp -R static requirements/ \
  && cp -R templates requirements/ \
  && cp website requirements/

FROM alpine:3.14
RUN apk add --no-cache curl bash ca-certificates
WORKDIR /opt/app/
COPY --from=builder go/src/github.com/gin-html-website/requirements/ /opt/app/
ARG APP_PATH
ENV APP_PATH=/opt/app
CMD ["./website"]
