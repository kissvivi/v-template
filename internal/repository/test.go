package repository

import (
	"v-template/internal/model"
	"v-template/internal/service"
	"v-template/pkg/request"
)

var _ service.TestRepo = (*testRepo)(nil)

type testRepo struct {
	data *Data
}

func NewTestRepo(data *Data) service.TestRepo {
	return &testRepo{
		data: data,
	}
}

func (t testRepo) TestOk(request *request.TestRequest) string {
	//orm.DB.Create(model.Test{Name: request.Name})
	var ss = request.Name
	return ss
}

func (t testRepo) CreateTest(request *request.TestRequest) (model.Test, error) {

	test := model.Test{Name: request.Name}
	result := t.data.Mysql.Create(&test)
	return test, result.Error
}
