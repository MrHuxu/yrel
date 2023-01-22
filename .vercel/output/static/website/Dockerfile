FROM node:11.15.0-alpine AS node-builder

ENV NODE_ENV production

WORKDIR /work
COPY ./client /work/client
COPY ./package.json /work/
COPY ./package-lock.json /work/
COPY ./webpack.config.prd.js /work/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config webpack.config.prd.js

FROM golang:alpine AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./main.go /work/
COPY ./go.mod /work/
COPY ./go.sum /work/

RUN go mod download
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./templates /output/templates
COPY --from=node-builder /work/built/bundle.js /output/built/
COPY --from=go-builder /work/main /output/

EXPOSE 8082
ENTRYPOINT [ "./main" ]