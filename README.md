# Fixed Delay Middleware for Hoverfly

Simple [Hoverfly](https://hoverfly.io) middleware webserver for deployment on Pivotal Cloud Foundry.

The middleware webserver injects a fixed delay of five seconds before each response in a Hoverfly simulation is returned.

This application is part of the [hoverfly-cloud-foundry-experiment](https://github.com/tjcunliffe/hoverfly-cloud-foundry-experiment).

## Dependencies

- [Cloud Foundry CLI](https://github.com/cloudfoundry/cli)
- [Go 1.9](https://golang.org/dl/)

## Deployment

```
cf p
```