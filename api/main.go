package main

import (
	"flag"

	"webxam/config"
	"webxam/db"
	"webxam/server"
)

func main() {
	e := flag.String("e", "development", "環境変数の設定オプション")
	flag.Parse()

	config.Init(*e)
	db.Init()
	server.Init()
}
