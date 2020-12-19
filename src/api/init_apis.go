//+build wireinject

package api

import (
	"github.com/google/wire"
	"tasquest-server/src/api/v1"
	"tasquest-server/src/common/errorhandler"
	"tasquest-server/src/infra/web"
	"tasquest-server/src/security"
)

func InitAuthApi() *v1.AuthAPI {
	wire.Build(v1.ProvideAuthAPI, web.ProvideWebserver, security.SecurityProvider, errorhandler.ErrorHandlerProvider)
	return &v1.AuthAPI{}
}
