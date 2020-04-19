#!/bin/sh

docker image build -t meetingunit:0.2 . ;
docker save meetingunit:0.2 | ssh root@178.128.165.236 docker load
ssh root root@178.128.165.236 docker stop meetingunit
ssh root root@178.128.165.236 docker rm meetingunit
ssh root root@178.128.165.236 docker run -d -p 9091:8080 --link=redis_cont --name meetingunit meetingunit


