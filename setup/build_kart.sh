#!/bin/bash

git clone https://github.com/STJr/Kart-Public.git
cd Kart-Public/src
git checkout master
LIBGME_CFLAGS= LIBGME_LDFLAGS=-lgme make
