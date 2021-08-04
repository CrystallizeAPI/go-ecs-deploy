# Go ECS Deploy

A simple script for redeploying services on AWS.

## Prerequisites

- Valid AWS credentials as well as an `AWS_DEPLOY_ROLE_ARN` environment variable to assume.
- An existing task definition in ECS - this won't be created for you.

## Installation

```sh
go get github.com/CrystallizeAPI/go-ecs-deploy
```

## Usage

```sh
go-ecs-deploy -cluster <cluster-name> -service <service-name> -region <aws-region>
```
