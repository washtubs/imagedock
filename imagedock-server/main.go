package main

import (
	"flag"
	"io/ioutil"
	l "log"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	logging "github.com/op/go-logging"
	"github.com/xenogt/imagedock"
)

var log *logging.Logger = imagedock.GetLogger()

type LaunchConfig struct {
	ListenAddr   string
	PostgresOpts *pg.Options
	MinioOpts    *imagedock.MinioConfig
}

func readPassOrTokenFromFile(filename string, maxSize int) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %v... %v", filename, err)
	}
	f.Close()

	passwd, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read from password file %s. An error occurred: %v",
			filename, err)
	}

	if len(passwd) > maxSize {
		log.Fatalf("Password from %s exceeds max pass size %v > %v", filename, len(passwd), maxSize)
	}

	return string(passwd)
}

func argsToLaunchConfig([]string) *LaunchConfig {

	defaultFlagSet.Parse(os.Args[1:])

	lconfig := new(LaunchConfig)
	lconfig.PostgresOpts = new(pg.Options)
	lconfig.MinioOpts = new(imagedock.MinioConfig)

	defaultFlagSet.VisitAll(func(flag *flag.Flag) {
		log.Debug(flag.Name)
		switch flag.Name {
		case "listen-addr":
			lconfig.ListenAddr = flag.Value.String()
		case "db-addr":
			lconfig.PostgresOpts.Addr = flag.Value.String()
		case "db-user":
			lconfig.PostgresOpts.User = flag.Value.String()
		case "db-passwd":
			lconfig.PostgresOpts.Password = flag.Value.String()
		case "db-passwd-file":
			if flag.Value.String() != "" {
				lconfig.PostgresOpts.Password = readPassOrTokenFromFile(flag.Value.String(), 256)
			}
		case "minio-addr":
			lconfig.MinioOpts.Addr = flag.Value.String()
		case "minio-default-bucket":
			lconfig.MinioOpts.DefaultBucket = flag.Value.String()
		case "minio-access-key":
			lconfig.MinioOpts.AccessKey = flag.Value.String()
		case "minio-access-key-file":
			if flag.Value.String() != "" {
				lconfig.MinioOpts.AccessKey = readPassOrTokenFromFile(flag.Value.String(), 100)
			}
		case "minio-secret-key":
			lconfig.MinioOpts.SecretKey = flag.Value.String()
		case "minio-secret-key-file":
			if flag.Value.String() != "" {
				lconfig.MinioOpts.SecretKey = readPassOrTokenFromFile(flag.Value.String(), 100)
			}
		}
	})
	return lconfig
}

func main() {
	l.SetFlags(l.Flags() | l.Llongfile)

	lconfig := argsToLaunchConfig(os.Args[1:])

	log.Debugf("Minio %#v, Postgres %#v", lconfig.MinioOpts, lconfig.PostgresOpts)
	//_ := postgres.Options{}

	http.ListenAndServe(lconfig.ListenAddr, imagedock.NewHttpHandler())
}
