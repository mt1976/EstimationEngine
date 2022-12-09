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
- __design__
   - __catalog__
     - [catalog.md](catalog/catalog.md)
     - [confidence.md](catalog/confidence.md)
     - [credentials.md](catalog/credentials.md)
     - [credentialsAction.md](catalog/credentialsAction.md)
     - [docType.md](catalog/docType.md)
     - [estimationSession.md](catalog/estimationSession.md)
     - [estimationSessionAction.md](catalog/estimationSessionAction.md)
     - [estimationState.md](catalog/estimationState.md)
     - [externalMessage.md](catalog/externalMessage.md)
     - [feature.md](catalog/feature.md)
     - [featureNew.md](catalog/featureNew.md)
     - [inbox.md](catalog/inbox.md)
     - [origin.md](catalog/origin.md)
     - [originState.md](catalog/originState.md)
     - [profile.md](catalog/profile.md)
     - [project.md](catalog/project.md)
     - [projectAction.md](catalog/projectAction.md)
     - [projectState.md](catalog/projectState.md)
     - [resource.md](catalog/resource.md)
     - [schedule.md](catalog/schedule.md)
     - [session.md](catalog/session.md)
     - [translation.md](catalog/translation.md)
     - [userRole.md](catalog/userRole.md)

---
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![Go](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml)
[![Docker - CI](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml)
[![Docker - Publish](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
