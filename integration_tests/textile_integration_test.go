package integration_tests

import (
	"bytes"
	"context"
	"flag"
	"os"
	"testing"

	"github.com/FleekHQ/space-daemon/config"
	"github.com/FleekHQ/space-daemon/core/env"
	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/FleekHQ/space-daemon/core/textile"
	"github.com/stretchr/testify/assert"
)

var (
	devMode        = flag.Bool("dev", true, "defaulting to true for testing")
	ipfsaddr       string
	mongousr       string
	mongopw        string
	mongohost      string
	mongorepset    string
	spaceapi       string
	spacehubauth   string
	textilehub     string
	textilehubma   string
	textilethreads string
	textileClient  textile.Client
	ctx            context.Context
)

func TestMain(m *testing.M) {
	if os.Getenv("INT_TESTS") != "true" {
		return
	}

	ctx = context.Background()

	// setup
	cf := &config.Flags{
		Ipfsaddr:             ipfsaddr,
		Mongousr:             mongousr,
		Mongopw:              mongopw,
		Mongohost:            mongohost,
		Mongorepset:          mongorepset,
		ServicesAPIURL:       spaceapi,
		ServicesHubAuthURL:   spacehubauth,
		DevMode:              *devMode == true,
		TextileHubTarget:     textilehub,
		TextileHubMa:         textilehubma,
		TextileThreadsTarget: textilethreads,
	}

	env := env.New()
	cfg := config.NewMap(env, cf)
	st := store.New(
		store.WithPath("/tmp"),
	)
	if err := st.Open(); err != nil {
		return
	}

	// setup local buckets
	buckd := textile.NewBuckd(cfg)
	err := buckd.Start(ctx)
	if err != nil {
		return
	}

	// setup textile client
	textileClient = textile.NewClient(st)
	go func() {
		textileClient.StartAndBootstrap(ctx, cfg)
	}()
	<-textileClient.WaitForReady()

	// run
	os.Exit(m.Run())
}

func TestAddAndRetreiveFile(t *testing.T) {
	b, _ := textileClient.CreateBucket(ctx, "testbucket1")

	assert.NotNil(t, b)

	f, _ := os.Open("./test_files/file1")
	p1, p2, _ := b.UploadFile(ctx, "file1", f)
	assert.NotNil(t, p1)
	assert.NotNil(t, p2)

	var bb bytes.Buffer
	b.GetFile(ctx, "file1", &bb)
	assert.Equal(t, bb.String(), "test data")
}
