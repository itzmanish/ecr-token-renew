## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

  helm repo add ecr-token-renew https://itzmanish.github.io/ecr-token-renew

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
ecr-token-renew` to see the charts.

To install the ecr-token-renew chart:

    helm install my-ecr-token-renew ecr-token-renew/chart
    
> Note: you need to override `configMapKeyRef` key with correct configmap value
where your AWS credentials are present. you can always append `-f override.yaml`
on the above command. You can check charts/deployment/values.yaml for default values.

To uninstall the chart:

    helm uninstall  my-ecr-token-renew
