## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

  helm repo add itzmanish-gitops https://itzmanish.github.io/ecr-token-renew

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
ecr-token-renew` to see the charts.

To install the ecr-token-renew chart:

    helm install my-ecr-token-renew ecr-token-renew/ecr-token-renew

To uninstall the chart:

    helm delete my-ecr-token-renew