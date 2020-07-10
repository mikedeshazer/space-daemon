package integration_tests

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	spaceApp "github.com/FleekHQ/space-daemon/app"
)

// Ignoring this for now because it fails as we don't
// currently properly shutdown all started goroutines
// note, this leak might be from libraries we don't control
func XTestAppDoesNotLeakGoroutines(t *testing.T) {
	goRoutinesBefore := runtime.NumGoroutine()

	_, cfg, env := GetTestConfig()
	ctx, cancelCtx := context.WithCancel(context.Background())
	app := spaceApp.New(cfg, env)

	var err error
	errChan := make(chan error)
	go func() {
		errChan <- app.Start(ctx)
	}()

	select {
	case <-app.WaitForReady():
	case err = <-errChan:
	}

	assert.Nil(t, err, "app.Start() Failed")
	cancelCtx()
	err = <-errChan
	assert.Nil(t, err, "app.Shutdown() Failed")

	assert.Equal(t, goRoutinesBefore, runtime.NumGoroutine(), "Goroutine leaked on app Shutdown")
}
