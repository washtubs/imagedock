package imagedock

import (
	"sync/atomic"
)

var atomicRequestId int64 = -1

func nextRequestId() int64 {
	return atomic.AddInt64(&atomicRequestId, 1)
}

var contextKeys = struct {
	userId    string
	requestId string
}{
	userId:    "userId",
	requestId: "requestId",
}
