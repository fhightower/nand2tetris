FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get -y update && \
    apt-get install -y less tree vim sudo wget openjdk-8-jre build-essential git

RUN useradd -g users -G sudo -m -s /bin/bash ubuntu && \
    echo 'ubuntu:foobar' | chpasswd
RUN echo 'Defaults visiblepw'            >> /etc/sudoers
RUN echo 'ubuntu ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER ubuntu

WORKDIR /home/ubuntu/
