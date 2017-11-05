// +build integration

package main

import (
	"os/exec"
	"strings"
	"testing"
)

func getProjectRootOrDie(t *testing.T) string {
	buf, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	if err != nil {
		t.Fatal("Error getting project root.", err, string(buf))
	}
	return strings.TrimSpace(string(buf))
}

func getDockerComposeLaunchConfig(t *testing.T) LaunchConfig {
	root := getProjectRootOrDie(t)
	args := []string{
		"-minio-addr", "localhost"
		"-minio-default-bucket", "testbuck"
		"-minio-access-key-file", root+"/minio_access_key"
		"-minio-secret-key-file", root+"/minio_secret_key"
	}
	return argsToLaunchConfig(args)
}

func reinitPostgresTables(t *testing.T) {
	cmd := exec.Command("scripts/postgres_init.sh")
	cmd.Dir = getProjectRootOrDie(t)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("Error getting project root.", err, string(buf))
	}
}

func TestMinio(t *testing.T) {
	lconfig := argsToLaunchConfig(getDockerComposeLaunchConfig(t))

}

func TestDatabase(t *testing.T) {

	reinitPostgresTables(t)

	//t.Fail()
}
