FROM ubuntu:latest

WORKDIR /srv

ENTRYPOINT ["/srv/bin/alphabet"]

ENV APP_NAME=alphabet-api

RUN mkdir ./bin

COPY ./main ./bin/alphabet

EXPOSE 8080