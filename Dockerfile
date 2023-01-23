FROM node:11.15.0-alpine AS node-builder

ENV NODE_ENV production

WORKDIR /work
COPY ./website/client /work/client
COPY ./website/package.json /work/
COPY ./website/package-lock.json /work/
COPY ./website/webpack.config.prd.js /work/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config webpack.config.prd.js

FROM golang:alpine AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./website/main.go /work/website/
COPY ./lexer /work/lexer
COPY ./parser /work/parser
COPY ./go.mod /work/
COPY ./go.sum /work/

RUN go mod download
RUN go build -o main website/main.go

FROM alpine

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./website/templates /output/templates
COPY ./website/client/assets /output/client/assets
COPY --from=node-builder /work/built/bundle.js /output/built/
COPY --from=go-builder /work/main /output/

EXPOSE 8082
ENTRYPOINT [ "./main" ]