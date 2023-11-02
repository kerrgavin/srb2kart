#!/bin/bash

apt-get update
apt-get -y install sudo
useradd -m docker && echo "docker:docker" | chpasswd && adduser docker sudo
echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
