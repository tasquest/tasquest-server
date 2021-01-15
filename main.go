package main

import (
	"tasquest.com/server"
	"tasquest.com/server/infra/web"
)

func main() {
	server.PrintLogo()

	Bootstrap()

	router := web.ProvideWebServer()
	_ = router.Run()
}
