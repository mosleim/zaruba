FROM ubuntu

ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get upgrade -y

RUN apt-get install make wget curl git python3 python3-pip nodejs npm golang docker.io -y

RUN npm install -g typescript

RUN alias python='python3'
RUN alias pip='pip3'