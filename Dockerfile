FROM golang:1.18-alpine3.16 as build
WORKDIR /app
COPY . .
RUN GOPROXY="https://goproxy.io,direct" go build -o bin/toolbox cmd/*.go

FROM alpine:3.16 as run
RUN apk add alpine-conf && /sbin/setup-timezone -z Asia/Shanghai && apk del alpine-conf
WORKDIR /app
EXPOSE 9000 31000
COPY --from=build /app/bin/toolbox .
ENTRYPOINT [ "/app/toolbox" ]
