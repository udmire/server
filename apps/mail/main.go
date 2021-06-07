package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"

	"e.coding.net/mmstudio/blade/server/services/mail"
	"e.coding.net/mmstudio/blade/server/utils"
	log "github.com/rs/zerolog/log"

	// micro plugins
	_ "github.com/micro/go-plugins/broker/nsq/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/store/consul/v2"
	_ "github.com/micro/go-plugins/transport/grpc/v2"
)

var (
	BinaryVersion string
	GoVersion     string
	GitLastLog    string
)

func version() {
	fmt.Println("BinaryVersion:", BinaryVersion)
	fmt.Println("GoVersion:", GoVersion)
	fmt.Println("GitLastLog:", GitLastLog)
	os.Exit(0)
}

func help() {
	fmt.Println("The commands are:")
	fmt.Println("version       see all versions")
	os.Exit(0)
}

func main() {
	utils.LDFlagsCheck(os.Args, version, help)

	m := mail.New()
	if err := m.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("mail run failed")
	}

	m.Stop()
}
