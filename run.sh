#!/bin/bash

./addonHandler /app/configs/addon_config.json /app/addons/
cd .srb2kart
./lsdl2srb2kart -dedicated -server -home /app -file /app/addons/misc/* /app/addons/maps/* /app/addons/characters/*
