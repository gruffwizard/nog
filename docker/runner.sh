#!/bin/bash
echo "CMD: $1"
echo "PWD:$PWD"
echo "M2:$M2_HOME"

echo "v2"
cd /home/nog/tools/theia && yarn start /home/nog/src --hostname 0.0.0.0 --port 8081 &

echo "Hello"
mvn compile quarkus:dev
