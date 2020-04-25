#!/bin/sh


host=104.248.163.157

docker image build -t meetingunit:0.2 . ;
docker save meetingunit:0.2 | ssh root@$host docker load
ssh root@$host docker stop meetingunit
ssh root@$host docker rm meetingunit
ssh root@$host docker run -d -p 9091:8080 --link=redis --name meetingunit -v /root/sert:/meetingbuild meetingunit:0.2


