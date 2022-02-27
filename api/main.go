package main

import (
	"flag"

	"webxam/config"
	"webxam/database"
	"webxam/server"
)

func main() {
	e := flag.String("e", "development", "環境変数の設定オプション")
	flag.Parse()

	config.Init(*e)
	database.Init()
	server.Init()
}
