# build stage
FROM golang:alpine AS build-env
RUN apk update; apk add git; mkdir -p /go/src/goapp
ADD . /go/src/goapp/
RUN cd /go/src/goapp/ && go-wrapper download && go-wrapper install;
RUN cd /go/src/goapp/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static" -w -s' -o /go/src/goapp/goapp

# final stage
FROM damdo/sscratch
MAINTAINER https://github.com/damdo
WORKDIR /app
COPY --from=build-env /go/src/goapp/goapp /app/
CMD ["/app/goapp"]
