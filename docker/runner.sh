#!/bin/bash
echo "CMD: $1"
echo "PWD:$PWD"
echo "M2:$M2_HOME"

echo "v2"
cd /home/qlauncher/tools/thea && yarn start /home/qlauncher/src --hostname 0.0.0.0 --port 8081 &

#mvn compile quarkus:dev
bash
