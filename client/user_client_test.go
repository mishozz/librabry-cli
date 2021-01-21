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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("users", nil)
			return m
		},
		expectedString: "users",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("success", nil)
			return m
		},
		expectedString: "success",
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 204}, nil)
			return m
		},
		err: nil,
	}, {
		name: "Unauthorized",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 401}, nil)
			return m
		},
		err: errors.New("Unauthorized"),
	}, {
		name: "unable to delete book",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("Do", mock.Anything).Return(&http.Response{StatusCode: 500}, nil)
			return m
		},
		err: errors.New("Unable to return book"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			u := &userClient{
				client: tt.mockHTTPClient(m),
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
