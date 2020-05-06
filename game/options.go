package game

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewFlags() []cli.Flag {
	return []cli.Flag{
		// game settings
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "debug", Usage: "debug mode"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "game_id", Usage: "game server unique id(0 - 1024)"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "account_connect_max", Usage: "how many account connections can be dealwith"}),

		// ip and port
		altsrc.NewStringFlag(&cli.StringFlag{Name: "public_ip", Usage: "public ip for clients connecting"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "https_listen_addr", Usage: "https listen address"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "tcp_listen_addr", Usage: "tcp listen address"}),

		// cert
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_debug", Usage: "debug tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_debug", Usage: "debug tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_release", Usage: "release tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_release", Usage: "release tls server_key path"}),

		// db
		altsrc.NewStringFlag(&cli.StringFlag{Name: "db_dsn", Usage: "db data source name"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "database", Usage: "database name"}),

		// micro service
		altsrc.NewStringFlag(&cli.StringFlag{Name: "registry", Usage: "micro service registry"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "broker", Usage: "micro service broker"}),
		&cli.StringFlag{
			Name:  "config_file",
			Value: "../../config/game/config.toml",
		},
	}
}
