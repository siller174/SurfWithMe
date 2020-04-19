#!/bin/sh

docker image build -t meetingunit:0.2 . ;
docker save meetingunit:0.2 | ssh root@178.128.165.236 docker load