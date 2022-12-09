# ![Estimation Engine](assets/favicon-32x32.png) Estimation Engine
The Estimation Engine is a small Go tool designed to generate reliable estimates for work undertaken by Eurobase
# Software Requirements
* Go v1.19
* MS SQL Server

# Docker Compose
```yaml
---
version: '3.3'
services:
  estimation_engine:
    image: mt1976/estimation_engine:latest
    container_name: estimation_engine
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
      - CONTEXT_PATH=url-base
    ports:
      - 5050:5050
    networks:
      - estEngNetwork
    restart: unless-stopped
networks:
  estEngNetwork:{}
logging:
  options:
    max-size: 1g
```
---
# Data Model
Object Information can be found [here](design/catalog) 
---
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![Go](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml)
[![Docker - CI](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml)
[![Docker - Publish](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
