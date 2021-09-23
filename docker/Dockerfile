FROM golang:alpine

RUN apk add --purge --no-cache --update xpdf poppler-utils inotify-tools git wv libc-dev gcc

RUN mkdir -p /go/src/github.com/mihnealun/commentix

#&& go get -u github.com/go-delve/delve/cmd/dlv
#RUN go get -u github.com/DATA-DOG/godog/cmd/godog

WORKDIR /go/src/github.com/mihnealun/commentix

CMD ./docker/autobuild.sh;