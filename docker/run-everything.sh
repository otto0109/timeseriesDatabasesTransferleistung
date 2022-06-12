#!/bin/bash

docker-compose -f influxDB/docker-compose.yml up -d
docker-compose -f postgre/docker-compose.yml up -d
docker-compose -f questDB/docker-compose.yml up -d
docker-compose -f timescaleDB/docker-compose.yml up -d