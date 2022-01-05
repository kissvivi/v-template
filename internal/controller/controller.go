package controller

import (
	"v-template/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewController)

type Controller struct {
	testService *service.TestUseCase
	demoService *service.DemoUseCase
}

func NewController(testService *service.TestUseCase, demoService *service.DemoUseCase) *Controller {
	return &Controller{testService: testService, demoService: demoService}
}
