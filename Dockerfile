FROM        golang:1.13-alpine3.10 as build

WORKDIR     /meetingbuild

ENV         GOBIN=/meetingbuild/bin

ADD         . /meetingbuild

RUN         apk update && apk add --no-cache git

RUN         go install -mod=vendor /meetingbuild/cmd/...

FROM		golang:1.13-alpine3.10

WORKDIR     /opt/

COPY        --from=build /meetingbuild/bin/meetingHelper /opt/
COPY        --from=build /npsbuild/configs/meetingHelper/meetingHelper.properties /opt

ENTRYPOINT ["./opt/meetingHelper --c"]

EXPOSE      8080
