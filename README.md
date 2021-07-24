# Renew Kubernetes Docker secrets for AWS ECR

AWS Elastic Container Registry (ECR) provides a cost-effective private registry for your Docker containers.
However, ECR Docker credentials
[expire every 12 hours](https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login.html).

To work around this, I created this small tool to automatically refresh the secret in Kubernetes.
It deploys as a cron job and ensures that your Kubernetes cluster
will always be able to pull Docker images from ECR.

## Docker Images

The latest images are:

- `itzmanish/ecr-token-renew:latest`

## Environment Variables

The tool is mainly configured through environment variables. These are:

- AWS_ACCESS_KEY_ID (required): AWS access key used to create the Docker credentials.
- AWS_SECRET_ACCESS_KEY (required): AWS secret needed to fetch Docker credentials from AWS.
- AWS_REGION (required): The AWS region where your ECR instance is created.
- DOCKER_SECRET_NAME (required): The name of the Kubernetes secret where the Docker credentials are stored.
- TARGET_NAMESPACE (optional): Comma-separated list of namespaces.
  A Docker secret is created in each of these.
  If this environment variable is not set, a value of `default` is assumed.
- DOCKER_REGISTRIES (optional): Comma-separated list of registry URL.
  If none is provided, the default URL returned from AWS is used.
  - Example: `DOCKER_REGISTRIES=https://321321.dkr.ecr.us-west-2.amazonaws.com,https://123123.dkr.ecr.us-east-2.amazonaws.com`

# Installing instruction

```bash
$ go install github.com/itzmanish/ecr-token-renew@latest
```
