// +build integration

package main

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	minio "github.com/minio/minio-go"
	"github.com/washtubs/imagedock"
)

func TestMinioBasic(t *testing.T) {
	lconfig := getDockerComposeLaunchConfig(t)
	minioOpts := lconfig.MinioOpts

	mc, err := imagedock.CreateMinioClient(minioOpts)
	if err != nil {
		t.Logf("Error creating minio client: %v", err)
		t.FailNow()
	}

	err = mc.MakeBucket(minioOpts.DefaultBucket, minioOpts.DefaultLocation)
	if err != nil {
		t.Log("Error making bucket", err)
	}

	t.Log("Made bucket")

	b, err := mc.BucketExists(minioOpts.DefaultBucket)
	if err != nil {
		t.Log("Error checking if bucket exists", err)
		t.FailNow()
	}

	if b {
		t.Log("bucket exists!")
	}

	s := "objectcontents"
	reader := strings.NewReader(s)
	//mc.PutObjectWithContext(
	_, err = mc.PutObjectWithContext(context.Background(),
		minioOpts.DefaultBucket,
		"blah",
		reader,
		-1, //For size input as -1 PutObject does a multipart Put operation until input stream reaches EOF.
		minio.PutObjectOptions{})

	if err != nil {
		t.Log("error putting object", err)
		t.FailNow()
	}

	obj, err := mc.GetObjectWithContext(context.Background(),
		minioOpts.DefaultBucket,
		"blah",
		minio.GetObjectOptions{})

	if err != nil {
		t.Log("error getting object", err)
		t.FailNow()
	}

	buf, err := ioutil.ReadAll(obj)
	if err != nil {
		t.Log("Error reading object stream", err)
		t.FailNow()
	}

	t.Logf("Got object back: %#v", string(buf))

}

func TestMinioReadOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Read can be cancelled
	// TODO: bonus points if we can show that the resource intensive operation is also cancelled
}

func TestMinioWriteOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Write can be cancelled
	// the file should not be there if it's cancelled
}
