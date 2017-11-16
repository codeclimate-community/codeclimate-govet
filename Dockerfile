ARG BASE=1.9.3-alpine3.7
FROM golang:${BASE} as build

WORKDIR /usr/src/app

COPY engine.json ./engine.json.template
RUN apk add --no-cache jq
RUN export go_version=$(go version | cut -d ' ' -f 3) && \
    cat engine.json.template | jq '.version = .version + "/" + env.go_version' > ./engine.json

COPY codeclimate-govet.go ./
RUN apk add --no-cache git
RUN go get -t -d -v ./...
RUN go build -o codeclimate-govet .


FROM golang:${BASE}

LABEL maintainer="Code Climate <hello@codeclimate.com>"

WORKDIR /usr/src/app

RUN adduser -u 9000 -D app

COPY --from=build /usr/src/app/engine.json /
COPY --from=build /usr/src/app/codeclimate-govet ./

USER app

VOLUME /code

CMD ["/usr/src/app/codeclimate-govet"]
