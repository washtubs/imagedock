// +build integration

package imagedock

import (
	"testing"

	minio "github.com/minio/minio-go"
	"github.com/washtubs/imagedock"
)

func TestMinioBasic(t *testing.T) {
	lconfig := argsToLaunchConfig(getDockerComposeLaunchConfig(t))
	var mc *minio.Client
	mc = imagedock.CreateMinioClient(lconfig.MinioOpts)

}

func TestMinioReadOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Read can be cancelled
	// TODO: bonus points if we can show that the resource intensive operation is also cancelled
}

func TestMinioWriteOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Write can be cancelled
	// the file should not be there if it's cancelled
}
