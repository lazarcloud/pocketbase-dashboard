#!/bin/bash

docker build . -t monsieurlazar/lazarbase
docker run -d -p 8081:80 -e ORIGIN=http://localhost:8081 --name lazar-dash -v /var/run/docker.sock:/var/run/docker.sock -v /home/pocketbase/metadata:/data --network=lazar-static monsieurlazar/lazarbase