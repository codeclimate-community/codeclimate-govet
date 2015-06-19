FROM alpine:edge

WORKDIR /usr/src/app
COPY bin/ /usr/src/app

RUN apk --update add go git build-base && \
    GOPATH=/ go get golang.org/x/tools/cmd/vet && \
    apk del build-base && rm -fr /usr/share/ri

CMD ["/usr/src/app/codeclimate-govet"]
