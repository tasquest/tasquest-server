package main

import (
	"tasquest.com/server/commons"
	"tasquest.com/server/infra/web"
)

func main() {
	commons.PrintLogo()

	Bootstrap()

	router := web.ProvideWebServer()
	_ = router.Run()
}
