// +build integration

package main

import "testing"

func TestPostgresReadOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Read can be cancelled
	// TODO: bonus points if we can show that the resource intensive operation is also cancelled
}

func TestPostgresWriteOpRespectsCancellationViaContext(t *testing.T) {
	// TODO: performing a long running Write can be cancelled
	// TODO: the the transaction should be rolled back
}
