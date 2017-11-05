package main

import "flag"

var defaultFlagSet *flag.FlagSet

func init() {
	defaultFlagSet = flag.NewFlagSet("default", flag.ExitOnError)
	defaultFlagSet.String("listen-addr", "0.0.0.0:9094", "")
	defaultFlagSet.String("db-addr", "", "")
	defaultFlagSet.String("db-user", "", "")
	defaultFlagSet.String("db-passwd", "", "")
	defaultFlagSet.String("db-passwd-file", "", "")
	defaultFlagSet.String("minio-addr", "localhost", "")
	defaultFlagSet.String("minio-default-bucket", "defaultbucket", "")
	defaultFlagSet.String("minio-access-key", "", "")
	defaultFlagSet.String("minio-access-key-file", "", "")
	defaultFlagSet.String("minio-secret-key", "", "")
	defaultFlagSet.String("minio-secret-key-file", "", "")
}
