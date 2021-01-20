package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserClient_GetAllUsers(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			respStr, err := u.GetAllUsers("test")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_UserClient_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			respStr, err := u.GetUser("test", "email")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_UserClient_Login(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			respStr, err := u.Login("email", "password")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_UserClient_Logout(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			respStr, err := u.Logout("test")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_UserClient_TakeBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("success", nil)
			return m
		},
		expectedString: "success",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			respStr, err := u.TakeBook("test", "email", "isbn")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_BookClient_ReturnBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 204}, nil)
			return m
		},
		err: nil,
	}, {
		name: "Unauthorized",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 401}, nil)
			return m
		},
		err: errors.New("Unauthorized"),
	}, {
		name: "unable to delete book",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 500}, nil)
			return m
		},
		err: errors.New("Unable to return book"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			u := &userClient{
				client: tt.mockHttpClient(m),
			}
			err := u.ReturnBook("test", "email", "isbn")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
