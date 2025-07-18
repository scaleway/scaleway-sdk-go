package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	accountSweeper "github.com/scaleway/scaleway-sdk-go/api/account/v3/sweepers"
	applesiliconSweeper "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1/sweepers"
	baremetalSweeper "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1/sweepers"
	blockSweeper "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1/sweepers"
	cockpitSweeper "github.com/scaleway/scaleway-sdk-go/api/cockpit/v1/sweepers"
	containerSweeper "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1/sweepers"
	flexibleipSweeper "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1/sweepers"
	functionSweeper "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1/sweepers"
	iamSweeper "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1/sweepers"
	inferenceSweeper "github.com/scaleway/scaleway-sdk-go/api/inference/v1/sweepers"
	inferenceSweeperBeta "github.com/scaleway/scaleway-sdk-go/api/inference/v1beta1/sweepers"
	instanceSweeper "github.com/scaleway/scaleway-sdk-go/api/instance/v1/sweepers"
	iotSweeper "github.com/scaleway/scaleway-sdk-go/api/iot/v1/sweepers"
	ipamSweeper "github.com/scaleway/scaleway-sdk-go/api/ipam/v1/sweepers"
	jobsSweeper "github.com/scaleway/scaleway-sdk-go/api/jobs/v1alpha1/sweepers"
	k8sSweeper "github.com/scaleway/scaleway-sdk-go/api/k8s/v1/sweepers"
	lbSweeper "github.com/scaleway/scaleway-sdk-go/api/lb/v1/sweepers"
	mnqSweeper "github.com/scaleway/scaleway-sdk-go/api/mnq/v1beta1/sweepers"
	mongodbSweeper "github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1/sweepers"
	rdbSweeper "github.com/scaleway/scaleway-sdk-go/api/rdb/v1/sweepers"
	redisSweeper "github.com/scaleway/scaleway-sdk-go/api/redis/v1/sweepers"
	registrySweeper "github.com/scaleway/scaleway-sdk-go/api/registry/v1/sweepers"
	secretSweeper "github.com/scaleway/scaleway-sdk-go/api/secret/v1beta1/sweepers"
	sdbSweeper "github.com/scaleway/scaleway-sdk-go/api/serverless_sqldb/v1alpha1/sweepers"
	vpcSweeper "github.com/scaleway/scaleway-sdk-go/api/vpc/v2/sweepers"
	vpcgwSweeper "github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1/sweepers"
	webhostingSweeper "github.com/scaleway/scaleway-sdk-go/api/webhosting/v1/sweepers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func main() {
	exitCode := mainNoExit()
	os.Exit(exitCode)
}

func getConfigProfile() *scw.Profile {
	config, err := scw.LoadConfig()
	if err != nil {
		return &scw.Profile{}
	}
	profile, err := config.GetActiveProfile()
	if err != nil {
		return &scw.Profile{}
	}

	return profile
}

func mainNoExit() int {
	configProfile := getConfigProfile()
	envProfile := scw.LoadEnvProfile()
	profile := scw.MergeProfiles(configProfile, envProfile)

	client, err := scw.NewClient(
		scw.WithProfile(profile),
		scw.WithUserAgent("sdk-sweeper"),
		scw.WithEnv(),
	)
	if err != nil {
		log.Fatalf("Cannot create Scaleway client: %w", err)
	}

	errs := []error(nil)

	err = accountSweeper.SweepAll(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping account: %w", err))
	}

	err = applesiliconSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping applesilicon: %w", err))
	}

	err = baremetalSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping baremetal: %w", err))
	}

	err = cockpitSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping cockpit: %w", err))
	}

	err = containerSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping container: %w", err))
	}

	err = flexibleipSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping flexibleip: %w", err))
	}

	err = functionSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping function: %w", err))
	}

	err = iamSweeper.SweepSSHKey(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping iam: %w", err))
	}

	err = inferenceSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping inference v1: %w", err))
	}

	err = inferenceSweeperBeta.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping inference v1beta1: %w", err))
	}

	err = instanceSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping instance: %w", err))
	}

	// Instance servers need to be swept before volumes and snapshots can be swept
	// because volumes and snapshots are attached to servers.
	err = blockSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping block: %w", err))
	}

	err = iotSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping iot: %w", err))
	}

	err = jobsSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping jobs: %w", err))
	}

	err = k8sSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping k8s: %w", err))
	}

	err = lbSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping lb: %w", err))
	}

	err = mongodbSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping mongodb: %w", err))
	}

	err = mnqSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping mnq: %w", err))
	}

	err = rdbSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping rdb: %w", err))
	}

	err = redisSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping redis: %w", err))
	}

	err = registrySweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping registry: %w", err))
	}

	err = secretSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping secret: %w", err))
	}

	err = sdbSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping sdb: %w", err))
	}

	//err = temSweeper.SweepDomain(client, region?, skippedDomain?)
	//if err != nil {
	//	errs = append(errs, fmt.Errorf("error sweeping tem: %w", err))
	//}

	err = vpcSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping vpc: %w", err))
	}

	err = vpcgwSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping vpcgw: %w", err))
	}

	err = webhostingSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping webhosting: %w", err))
	}

	// IPAM IPs need to be swept in the end because we need to be sure
	// that every resource with an attached ip is destroyed before executing it.
	err = ipamSweeper.SweepAllLocalities(client)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping ipam: %w", err))

	}

	if len(errs) > 0 {
		log.Fatalf(errors.Join(errs...).Error())

		return -1
	}

	return 0
}
