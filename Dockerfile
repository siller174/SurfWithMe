FROM        golang:1.13-alpine3.10 as build

WORKDIR     /meetingbuild

ENV         GOBIN=/meetingbuild/bin

ADD         . /meetingbuild

RUN         apk update && apk add --no-cache git

RUN         go install -mod=vendor /meetingbuild/cmd/...

FROM		golang:1.13-alpine3.10

WORKDIR     /opt/manager

COPY        --from=build /meetingbuild/bin/meetingHelper /opt/manager
COPY        --from=build /meetingbuild/configs/meetingHelper/meetingHelper.properties /opt/manager

ENTRYPOINT ["/opt/manager/meetingHelper"]

EXPOSE      8080
