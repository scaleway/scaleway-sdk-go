package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	containerSweeper "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1/sweepers"
	instanceSweeper "github.com/scaleway/scaleway-sdk-go/api/instance/v1/sweepers"
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
		log.Fatalf("Cannot create Scaleway client: %v", err)
	}

	errs := []error(nil)

	err = containerSweeper.SweepAllLocalities(client, false)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping container: %w", err))
	}

	err = instanceSweeper.SweepAllLocalities(client, false)
	if err != nil {
		errs = append(errs, fmt.Errorf("error sweeping instance: %w", err))
	}

	if len(errs) > 0 {
		log.Fatal(errors.Join(errs...).Error())

		return -1
	}

	return 0
}
