package cli

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookClient struct {
	mock.Mock
}

func (m *mockBookClient) GetAllBooks(token string) (str string, err error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockBookClient) GetBook(token, isbn string) (str string, err error) {
	args := m.Called(token, isbn)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockBookClient) SaveBook(token, isbn, title, author string, availableUnits uint) (str string, err error) {
	args := m.Called(token, isbn, title, author, availableUnits)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockBookClient) Delete(token, isbn string) error {
	args := m.Called(token, isbn)
	return args.Error(0)
}

func Test_GetBooksCmd(t *testing.T) {
	tests := []struct {
		name           string
		mockBookClient func(m *mockBookClient) *mockBookClient
		expectedOutput string
	}{{
		name: "success",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("GetAllBooks", mock.Anything).Return("books", nil)
			return m
		},
		expectedOutput: "books",
	}, {
		name: "error while fetching books",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("GetAllBooks", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to fetch books from library",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockBookClient{}
			getAllBooksCmd := NewGetBooksCmd(tt.mockBookClient(m))
			b := bytes.NewBufferString("")
			getAllBooksCmd.SetOut(b)
			getAllBooksCmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedOutput, string(out))
		})
	}
}

func Test_GetBookCmd(t *testing.T) {
	tests := []struct {
		name           string
		mockBookClient func(m *mockBookClient) *mockBookClient
		expectedOutput string
	}{{
		name: "success",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("GetBook", mock.Anything, mock.Anything).Return("book", nil)
			return m
		},
		expectedOutput: "book",
	}, {
		name: "error while fetching books",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("GetBook", mock.Anything, mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to fetch book with isbn  from library",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockBookClient{}
			getBookCmd := NewGetBookCmd(tt.mockBookClient(m))
			b := bytes.NewBufferString("")
			getBookCmd.SetOut(b)
			getBookCmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedOutput, string(out))
		})
	}
}

func Test_SaveBookCmd(t *testing.T) {
	tests := []struct {
		name           string
		mockBookClient func(m *mockBookClient) *mockBookClient
		expectedOutput string
	}{{
		name: "success",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("SaveBook", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("success", nil)
			return m
		},
		expectedOutput: "success",
	}, {
		name: "error while fetching books",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("SaveBook", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to save book with isbn ",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockBookClient{}
			saveBookCmd := NewSaveBookCmd(tt.mockBookClient(m))
			b := bytes.NewBufferString("")
			saveBookCmd.SetOut(b)
			saveBookCmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedOutput, string(out))
		})
	}
}

func Test_DeleteBookCmd(t *testing.T) {
	tests := []struct {
		name           string
		mockBookClient func(m *mockBookClient) *mockBookClient
		expectedOutput string
	}{{
		name: "success",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("Delete", mock.Anything, mock.Anything).Return(nil)
			return m
		},
		expectedOutput: "Book with isbn  successfully deleted",
	}, {
		name: "error while fetching books",
		mockBookClient: func(m *mockBookClient) *mockBookClient {
			m.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error"))
			return m
		},
		expectedOutput: "Unable to delete book with isbn ",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockBookClient{}
			deleteBookCmd := NewDeleteBookCmd(tt.mockBookClient(m))
			b := bytes.NewBufferString("")
			deleteBookCmd.SetOut(b)
			deleteBookCmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedOutput, string(out))
		})
	}
}
