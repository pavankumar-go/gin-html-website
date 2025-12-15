FROM golang:1.17-alpine AS builder
WORKDIR $GOPATH/src/github.com/gin-html-website
COPY . .
RUN apk add git ca-certificates g++
RUN go build -ldflags '-extldflags "-static"' -o website
RUN mkdir requirements \
  && cp -R static requirements/ \
  && cp -R scripts requirements/ \
  && cp -R templates requirements/ \
  && cp sitemap.xml requirements/

RUN cp website requirements/

FROM alpine:3.21
RUN apk add curl bash ca-certificates
WORKDIR /opt/app/
COPY --from=builder go/src/github.com/gin-html-website/requirements/ /opt/app/
ENV APP_PATH=/opt/app
CMD ["./website"]
