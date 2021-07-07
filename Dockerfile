ARG BASE=1.16.5-buster
FROM golang:${BASE} as build

WORKDIR /usr/src/app

COPY engine.json ./engine.json.template
RUN apt-get update && apt-get install -y jq
RUN export go_version=$(go version | cut -d ' ' -f 3) && \
    cat engine.json.template | jq '.version = .version + "/" + env.go_version' > ./engine.json

COPY codeclimate-govet.go go.mod go.sum ./
# RUN apk add --no-cache git
RUN go build -o codeclimate-govet .

FROM golang:${BASE}

LABEL maintainer="Code Climate <hello@codeclimate.com>"

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/engine.json /
COPY --from=build /usr/src/app/codeclimate-govet ./

VOLUME /code
WORKDIR /code

CMD ["/usr/src/app/codeclimate-govet"]
