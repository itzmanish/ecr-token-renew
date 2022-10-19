# Renew Kubernetes Docker secrets for AWS ECR

AWS Elastic Container Registry (ECR) provides a cost-effective private registry for your Docker containers.
However, ECR Docker credentials
[expire every 12 hours](https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login.html).

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

## Helm Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

  helm repo add ecr-token-renew https://itzmanish.github.io/ecr-token-renew

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
ecr-token-renew` to see the charts.

To install the ecr-token-renew chart:

    helm install my-ecr-token-renew ecr-token-renew/chart
    
Note: you need to override `configMapKeyRef` key with correct configmap value
where your AWS credentials are present. you can always append `-f override.yaml`
on the above command.

To uninstall the chart:

    helm uninstall  my-ecr-token-renew

