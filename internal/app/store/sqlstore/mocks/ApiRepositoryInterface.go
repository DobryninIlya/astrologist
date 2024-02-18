// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"
	firebase "main/internal/app/firebase"

	logrus "github.com/sirupsen/logrus"

	mock "github.com/stretchr/testify/mock"

	model "main/internal/app/model"

	openai "main/internal/app/openai"
)

// ApiRepositoryInterface is an autogenerated mock type for the ApiRepositoryInterface type
type ApiRepositoryInterface struct {
	mock.Mock
}

// AddAuthor provides a mock function with given fields: groupId
func (_m *ApiRepositoryInterface) AddAuthor(groupId int) bool {
	ret := _m.Called(groupId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(groupId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CheckSecret provides a mock function with given fields: secret
func (_m *ApiRepositoryInterface) CheckSecret(secret string) (bool, error, int) {
	ret := _m.Called(secret)

	var r0 bool
	var r1 error
	var r2 int
	if rf, ok := ret.Get(0).(func(string) (bool, error, int)); ok {
		return rf(secret)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(secret)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(secret)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(string) int); ok {
		r2 = rf(secret)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// CheckToken provides a mock function with given fields: tokenStr
func (_m *ApiRepositoryInterface) CheckToken(tokenStr string) (model.ApiClient, error, int) {
	ret := _m.Called(tokenStr)

	var r0 model.ApiClient
	var r1 error
	var r2 int
	if rf, ok := ret.Get(0).(func(string) (model.ApiClient, error, int)); ok {
		return rf(tokenStr)
	}
	if rf, ok := ret.Get(0).(func(string) model.ApiClient); ok {
		r0 = rf(tokenStr)
	} else {
		r0 = ret.Get(0).(model.ApiClient)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(string) int); ok {
		r2 = rf(tokenStr)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GetConfirmationCode provides a mock function with given fields:
func (_m *ApiRepositoryInterface) GetConfirmationCode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNewsById provides a mock function with given fields: id
func (_m *ApiRepositoryInterface) GetNewsById(id int) (model.News, error) {
	ret := _m.Called(id)

	var r0 model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.News, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.News); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.News)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewsPreviews provides a mock function with given fields: count, offset
func (_m *ApiRepositoryInterface) GetNewsPreviews(count int, offset int) ([]model.News, error) {
	ret := _m.Called(count, offset)

	var r0 []model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]model.News, error)); ok {
		return rf(count, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int) []model.News); ok {
		r0 = rf(count, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(count, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenInfo provides a mock function with given fields: tokenStr
func (_m *ApiRepositoryInterface) GetTokenInfo(tokenStr string) (model.ApiClient, error) {
	ret := _m.Called(tokenStr)

	var r0 model.ApiClient
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.ApiClient, error)); ok {
		return rf(tokenStr)
	}
	if rf, ok := ret.Get(0).(func(string) model.ApiClient); ok {
		r0 = rf(tokenStr)
	} else {
		r0 = ret.Get(0).(model.ApiClient)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeNews provides a mock function with given fields: news
func (_m *ApiRepositoryInterface) MakeNews(news model.News) (int, error) {
	ret := _m.Called(news)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(model.News) (int, error)); ok {
		return rf(news)
	}
	if rf, ok := ret.Get(0).(func(model.News) int); ok {
		r0 = rf(news)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(model.News) error); ok {
		r1 = rf(news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseNews provides a mock function with given fields: update, log, _a2
func (_m *ApiRepositoryInterface) ParseNews(update model.VKUpdate, log *logrus.Logger, _a2 *openai.ChatGPT) error {
	ret := _m.Called(update, log, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.VKUpdate, *logrus.Logger, *openai.ChatGPT) error); ok {
		r0 = rf(update, log, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistrationToken provides a mock function with given fields: ctx, client, _a2
func (_m *ApiRepositoryInterface) RegistrationToken(ctx context.Context, client *model.ApiClient, _a2 firebase.FirebaseAPIInterface) (string, error) {
	ret := _m.Called(ctx, client, _a2)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ApiClient, firebase.FirebaseAPIInterface) (string, error)); ok {
		return rf(ctx, client, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.ApiClient, firebase.FirebaseAPIInterface) string); ok {
		r0 = rf(ctx, client, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.ApiClient, firebase.FirebaseAPIInterface) error); ok {
		r1 = rf(ctx, client, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMobileUserInfo provides a mock function with given fields: user
func (_m *ApiRepositoryInterface) SaveMobileUserInfo(user model.ApiClient) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.ApiClient) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetConfirmationCode provides a mock function with given fields: code
func (_m *ApiRepositoryInterface) SetConfirmationCode(code string) {
	_m.Called(code)
}

type mockConstructorTestingTNewApiRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewApiRepositoryInterface creates a new instance of ApiRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApiRepositoryInterface(t mockConstructorTestingTNewApiRepositoryInterface) *ApiRepositoryInterface {
	mock := &ApiRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}