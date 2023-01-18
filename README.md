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

# Data Model
Object Information can be found [here](design/catalog)
- [catalog.md](design/catalog/catalog.md)
- [confidence.md](design/catalog/confidence.md)
- [credentials.md](design/catalog/credentials.md)
- [credentialsAction.md](design/catalog/credentialsAction.md)
- [docType.md](design/catalog/docType.md)
- [estimationSession.md](design/catalog/estimationSession.md)
- [estimationSessionAction.md](design/catalog/estimationSessionAction.md)
- [estimationState.md](design/catalog/estimationState.md)
- [externalMessage.md](design/catalog/externalMessage.md)
- [feature.md](design/catalog/feature.md)
- [featureNew.md](design/catalog/featureNew.md)
- [inbox.md](design/catalog/inbox.md)
- [origin.md](design/catalog/origin.md)
- [originState.md](design/catalog/originState.md)
- [profile.md](design/catalog/profile.md)
- [project.md](design/catalog/project.md)
- [projectAction.md](design/catalog/projectAction.md)
- [projectState.md](design/catalog/projectState.md)
- [resource.md](design/catalog/resource.md)
- [schedule.md](design/catalog/schedule.md)
- [session.md](design/catalog/session.md)
- [translation.md](design/catalog/translation.md)
- [userRole.md](design/catalog/userRole.md)

## Badges
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)
[![Go](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/go.yml)
[![Docker - CI](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker-image.yml)
[![Docker - Publish](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml/badge.svg)](https://github.com/mt1976/ebEstimates/actions/workflows/docker_push.yml)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
