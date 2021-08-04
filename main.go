package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var (
	clusterName = flag.String("cluster", "", "Name of cluster for deployment")
	serviceName = flag.String("service", "", "Name of service to deploy")
)

func fail(s string) {
	fmt.Println(s)
	os.Exit(2)
}

func main() {
	flag.Parse()

	if *clusterName == "" || *serviceName == "" {
		flag.Usage()
		fail("cluster and service names must be provided")
	}

	roleArn := os.Getenv("AWS_DEPLOY_ROLE_ARN")
	if roleArn == "" {
		fail("role arn must be provided")
	}

	forceNewDeployment := true
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, roleArn)

	ecs.New(sess, &aws.Config{
		Credentials: creds,
	}).UpdateService(&ecs.UpdateServiceInput{
		Cluster:            clusterName,
		Service:            serviceName,
		ForceNewDeployment: &forceNewDeployment,
	})
}
