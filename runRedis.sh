#!/bin/sh

docker run  --rm -d \
            -p 6379:6379 \
	        -e REDIS_PASSWORD=hardpassword \
    	    -v $PWD/redis-data:/bitnami/redis/data \
    	    --name redis \
    	    bitnami/redis:latest