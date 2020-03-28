FROM        golang:1.13-alpine3.10 as build

WORKDIR     /meetingHelper

ENV         GOBIN=/meetingHelper

RUN         apk update && apk add --no-cache git

COPY        /configs/meetingHelper .

ADD         . .

RUN         go install -mod=vendor /meetingHelper/cmd/...

ENTRYPOINT  ["./meetingHelper"]

EXPOSE      8080
