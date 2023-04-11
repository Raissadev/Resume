#!/bin/bash

apt-get update
apt-get install -y docker.io

systemctl start docker
systemctl enable docker

docker run -d -p 80:8000 raissageek/geek:v1.0