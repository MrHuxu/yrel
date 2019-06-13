FROM node:11.15.0 AS node-builder

WORKDIR /work
COPY ./client /work/client
COPY ./package.json /work/
COPY ./package-lock.json /work/
COPY ./webpack.config.prd.js /work/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config webpack.config.prd.js

FROM golang:latest AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/MrHuxu/yrel/website
COPY ./main.go /go/src/github.com/MrHuxu/yrel/website/
COPY ./go.mod /go/src/github.com/MrHuxu/yrel/website/
COPY ./go.sum /go/src/github.com/MrHuxu/yrel/website/

RUN go mod download
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./templates /output/templates
COPY --from=node-builder /work/built/bundle.js /output/built/
COPY --from=go-builder /go/src/github.com/MrHuxu/yrel/website/main /output/

EXPOSE 8082
ENTRYPOINT [ "./main" ]