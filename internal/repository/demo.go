package repository

import (
	"v-template/internal/service"
	"v-template/pkg/request"
	"fmt"
)

var _ service.DemoRepo = (*demoRepo)(nil)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) service.DemoRepo {
	return &demoRepo{data: data}
}

func (d demoRepo) DemoOk(request *request.TestRequest) string {
	fmt.Println("demo ok !!!!!!")

	return request.Name
}
