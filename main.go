package main

import (
	"tasquest.com/server/common"
	"tasquest.com/server/infra/web"
)

func main() {
	common.PrintLogo()

	Bootstrap()

	router := web.ProvideWebServer()
	_ = router.Run()
}
