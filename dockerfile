FROM alpine:latest

RUN apk add --update --no-cache tzdata ca-certificates && update-ca-certificates
ENV TZ=Asia/Almaty

COPY ./bin/binary /binary
RUN mkdir /config
ENTRYPOINT ["/binary"]
