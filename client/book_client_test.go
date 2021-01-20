package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHttpClient struct {
	mock.Mock
}

func (m *mockHttpClient) SendRequest(req *http.Request) (str string, err error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockHttpClient) Do(req *http.Request) (res *http.Response, err error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

func Test_BookClient_GetAllBook(t *testing.T) {
	tests := []struct {
		name           string
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("all books", nil)
			return m
		},
		expectedString: "all books",
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
			b := &bookClient{
				client: tt.mockHttpClient(m),
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
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
			m.On("SendRequest", mock.Anything).Return("book", nil)
			return m
		},
		expectedString: "book",
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
			b := &bookClient{
				client: tt.mockHttpClient(m),
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
		mockHttpClient func(m *mockHttpClient) *mockHttpClient
		expectedString string
		isbn           string
		title          string
		author         string
		units          uint
		err            error
	}{{
		name: "success",
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
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
		mockHttpClient: func(m *mockHttpClient) *mockHttpClient {
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
			m := &mockHttpClient{}
			b := &bookClient{
				client: tt.mockHttpClient(m),
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
		err: errors.New("Unable to delete book"),
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockHttpClient{}
			b := &bookClient{
				client: tt.mockHttpClient(m),
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
