// +build integration

package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func getSuperProjectRootOrDie(t *testing.T) string {

	envvar := os.Getenv("COMPOSE_ROOT")
	if envvar != "" {
		return envvar
	}

	buf, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	if err != nil {
		t.Fatal("Error getting project root.", err, string(buf))
	}
	return strings.TrimSpace(string(buf))
}

func getDockerComposeLaunchConfig(t *testing.T) *LaunchConfig {
	root := getSuperProjectRootOrDie(t)
	args := []string{
		"-minio-addr", "192.168.99.100:9000",
		"-minio-default-bucket", "testbuck",
		"-minio-default-location", "us-east-1",
		"-minio-access-key-file", root + "/minio_access_key",
		"-minio-secret-key-file", root + "/minio_secret_key",
	}
	return argsToLaunchConfig(args)
}

func reinitPostgresTables(t *testing.T) {
	cmd := exec.Command("scripts/postgres_init.sh")
	cmd.Dir = getSuperProjectRootOrDie(t)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("Error re-initializing postgres db.", err, string(buf))
	}
}

func TestMinio(t *testing.T) {
}

func TestReinitPostgresScript(t *testing.T) {

	reinitPostgresTables(t)

	//t.Fail()
}
