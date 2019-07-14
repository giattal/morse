#!/bin/bash
docker build -t ip-to-morse:0.0.1 .
docker run -d --privileged -p 5555:5555 ip-to-morse:0.0.1
