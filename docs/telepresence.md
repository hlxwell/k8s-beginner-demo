# Concepts

It will save your time to:
- Create dockerfile
- Push to docker image
- Create Deployment & Service for your docker
- Deploy to kubernetes
- Repeat above steps when you frequently update your application at local.

## Installation
https://www.telepresence.io/reference/install

## Swap Deployment
If you want to shutdown a running service and use local dev version, then you can use swap deployment.

```
telepresence --swap-deployment web-service --expose 8000 --run python3 -m http.server 8000 &
```

## New Deployment for telepresence

```
telepresence --new-deployment DEPLOY-SERVICE-NAME --run-shell # Create a new deployment and link local env to remote proxy.
telepresence --run curl http://service_name:8080 # Create random deployment name.
```

