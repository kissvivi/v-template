//go:build wireinject
// +build wireinject

package cmd

import (
	"v-template/config"
	"v-template/internal/controller"
	"v-template/internal/repository"
	"v-template/internal/service"
	"v-template/router"
	"github.com/google/wire"
	"net/http"
)

func InitApp(cfg *config.Config) (*http.Server, func(), error) {

	panic(wire.Build(repository.RepoSet, service.ProviderSet, controller.ProviderSet, router.InitRouterSet, InitServerSet))
}
