FROM golang:1.14.2-alpine AS builder

EXPOSE 9000:9000/tcp

ENV GO_DOMAIN="github.com" \
    GO_GROUP="fredericorecsky" \
    GO_PROJECT="yatodo"

ENV APP_DIR="${GOPATH}/src/${GO_DOMAIN}/${GO_GROUP}/${GO_PROJECT}"

RUN apk --update add git make gcc libc-dev

RUN mkdir -v -p ${APP_DIR}
WORKDIR ${APP_DIR}

COPY . ${APP_DIR}

RUN make clean
RUN make dep
RUN make build

# run

FROM golang:1.14.2-alpine

ENV GO_DOMAIN="github.com" \
    GO_GROUP="fredericorecsky" \
    GO_PROJECT="yatodo"

ENV APP_DIR="${GOPATH}/src/${GO_DOMAIN}/${GO_GROUP}/${GO_PROJECT}"

COPY --from=builder  ${APP_DIR}/build/${GO_PROJECT} /usr/local/bin/${GO_PROJECT}
CMD ["/usr/local/bin/yatodo"]


