package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHTTPClient struct {
	mock.Mock
}

func (m *mockHTTPClient) SendRequest(req *http.Request) (str string, err error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockHTTPClient) Do(req *http.Request) (res *http.Response, err error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

func Test_BookClient_GetAllBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("all books", nil)
			return m
		},
		expectedString: "all books",
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
			b := &bookClient{
				client: tt.mockHTTPClient(m),
			}
			respStr, err := b.GetAllBooks("test")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_BookClient_GetBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("book", nil)
			return m
		},
		expectedString: "book",
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
			b := &bookClient{
				client: tt.mockHTTPClient(m),
			}
			respStr, err := b.GetBook("test", "isbn")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_BookClient_SaveBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHTTPClient func(m *mockHTTPClient) *mockHTTPClient
		expectedString string
		isbn           string
		title          string
		author         string
		units          uint
		err            error
	}{{
		name: "success",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("success", nil)
			return m
		},
		expectedString: "success",
		isbn:           "test",
		title:          "test",
		author:         "test",
		units:          12,
		err:            nil,
	}, {
		name: "error while sending request",
		mockHTTPClient: func(m *mockHTTPClient) *mockHTTPClient {
			m.On("SendRequest", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedString: "",
		isbn:           "test",
		title:          "test",
		author:         "test",
		units:          12,
		err:            errors.New("error"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			b := &bookClient{
				client: tt.mockHTTPClient(m),
			}
			respStr, err := b.SaveBook("test", tt.isbn, tt.title, tt.author, tt.units)
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
			assert.Equal(t, tt.expectedString, respStr)
		})
	}
}

func Test_BookClient_DeleteBook(t *testing.T) {
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
		err: errors.New("Unable to delete book"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHTTPClient{}
			b := &bookClient{
				client: tt.mockHTTPClient(m),
			}
			err := b.Delete("test", "isbn")
			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
