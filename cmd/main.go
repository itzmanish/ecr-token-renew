package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/itzmanish/ecr-token-renew/pkg/aws"
	"github.com/itzmanish/ecr-token-renew/pkg/k8s"
)

const (
	envVarDockerSecretName = "DOCKER_SECRET_NAME"
	envVarTargetNamespace  = "TARGET_NAMESPACE"
	envVarRegistries       = "DOCKER_REGISTRIES"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Running at " + time.Now().UTC().String())

	name := os.Getenv(envVarDockerSecretName)
	if name == "" {
		panic(fmt.Sprintf("Environment variable %s is required", envVarDockerSecretName))
	}

	namespaceList := os.Getenv(envVarTargetNamespace)
	if namespaceList == "" {
		namespaceList = "default"
	}

	fmt.Println("Fetching auth data from AWS... ")
	credentials, err := aws.GetDockerCredentials()
	checkErr(err)

	servers := getServerList(credentials.Server)
	fmt.Printf("Docker Registries: %s\n", strings.Join(servers, ","))

	namespaces, err := k8s.GetNamespaces(namespaceList)
	checkErr(err)
	fmt.Printf("Updating kubernetes secret [%s] in %d namespaces\n", name, len(namespaces))

	failed := false
	for _, ns := range namespaces {
		fmt.Printf("Updating secret in namespace [%s]... ", ns)
		err = k8s.UpdatePassword(ns, name, credentials.Username, credentials.Password, servers)
		if nil != err {
			fmt.Printf("failed: %s\n", err)
			failed = true
		} else {
			fmt.Println("success")
		}
	}

	if failed {
		panic(errors.New("failed to create one of more Docker login secrets"))
	}

	fmt.Println("Job complete.")
}

func getServerList(defaultServer string) []string {
	addedServersSetting := os.Getenv(envVarRegistries)

	if addedServersSetting == "" {
		return []string{defaultServer}
	}

	return strings.Split(addedServersSetting, ",")
}
