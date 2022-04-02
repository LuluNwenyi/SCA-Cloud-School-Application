package test

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/docker"

)

func TestSCATask(t *testing.T) {

	var tag string = "sca-task-test/v0"
	var container_name string = "sca-task-test"

	// Build the Dockerfile
	buildOptions := &docker.BuildOptions{Tags: []string{tag}}
	docker.Build(t, "/Users/lulu/Documents/DevProjects/SCA_Task", buildOptions)
	fmt.Println("Build Successful")

	// Run the container
	opts := &docker.RunOptions{Entrypoint: "", Command: []string{"run", "-p", "5001"}, Name: container_name, Remove: true}
	docker.Run(t, tag, opts)
	fmt.Println("Container is running")
}
