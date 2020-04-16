FROM        golang:1.13-alpine3.10 as build

WORKDIR     /meetingbuild

ENV         GOBIN=/meetingbuild/bin

ADD         . /build

RUN         apk update && apk add --no-cache git

RUN         go install -mod=vendor /meetingbuild/cmd/...

WORKDIR     /opt/
FROM		golang:1.13-alpine3.10

COPY        --from=build /meetingbuild/bin/meetingHelper /opt/meetingHelper
COPY        --from=build /npsbuild/configs/meetingHelper/meetingHelper.properties /opt/meetingHelper

ENTRYPOINT ["./opt/meetingHelper --c"]

EXPOSE      8080
