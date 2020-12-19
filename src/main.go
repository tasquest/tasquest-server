package main

import (
	"tasquest-server/src/api"
	"tasquest-server/src/common"
	"tasquest-server/src/infra/web"
)

func main() {
	common.PrintLogo()

	api.InitAuthApi()

	router := web.ProvideWebserver()
	_ = router.Run()
}
