FROM        golang:1.13-alpine3.10 as build

WORKDIR     /meetingHelperbuild

ENV         GOBIN=/meetingHelperbuild/bin

ADD         . /meetingHelperbuild

RUN         apk update && apk add --no-cache git

RUN         go install -mod=vendor /meetingHelperbuild/cmd/...

ENTRYPOINT  ["./meetingHelperbuild/bin/meetingHelper --config-path /meetingHelperbuild/configs/meetingHelper/meetingHelper.properties"]

EXPOSE      8080
