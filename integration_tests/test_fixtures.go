package integration_tests

import (
	"github.com/FleekHQ/space-daemon/config"
	spaceEnv "github.com/FleekHQ/space-daemon/core/env"
	"github.com/namsral/flag"
)

var (
	devMode        = flag.Bool("dev", true, "defaulting to true for testing")
	ipfsaddr       = flag.String("ipfs_addr", "", "Set IPFS_ADDR for this value")
	mongousr       = flag.String("mongo_usr", "", "Required. Set MONGO_USR for this value")
	mongopw        = flag.String("mongo_pw", "", "Required. MONGO_PW")
	mongohost      = flag.String("mongo_host", "", "Required")
	mongorepset    = flag.String("mongo_replica_set", "", "Required")
	spaceapi       = flag.String("service_api_url", "", "Required")
	spacehubauth   = flag.String("service_hub_auth_url", "", "Required")
	textilehub     = flag.String("txl_hub_target", "", "Required. Set TXL_HUB_TARGET")
	textilehubma   = flag.String("txl_hub_ma", "", "Required")
	textilethreads = flag.String("txl_threads_target", "", "Required")
)

// GetTestConfig returns a ConfigMap instance instantiated using the env variables
func GetTestConfig() (*config.Flags, config.Config, spaceEnv.SpaceEnv) {
	flags := &config.Flags{
		Ipfsaddr:             *ipfsaddr,
		Mongousr:             *mongousr,
		Mongopw:              *mongopw,
		Mongohost:            *mongohost,
		Mongorepset:          *mongorepset,
		ServicesAPIURL:       *spaceapi,
		ServicesHubAuthURL:   *spacehubauth,
		DevMode:              *devMode == true,
		TextileHubTarget:     *textilehub,
		TextileHubMa:         *textilehubma,
		TextileThreadsTarget: *textilethreads,
	}

	// env
	env := spaceEnv.New()

	// load configs
	return flags, config.NewMap(flags), env
}

type TextileFixture struct{}

type AppFixture struct{}

func (a *AppFixture) Setup() {
	// TODO: Setup rpeated fixture logic, etc
}
