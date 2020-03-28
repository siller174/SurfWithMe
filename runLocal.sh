#!/bin/sh
set -e
export GOBIN=`pwd`/bin
go install ./cmd/...
unset GOBIN
echo "Your files in /bin/"
./bin/meetingHelper --config-path ./configs/meetingHelper/meetingHelper.properties
