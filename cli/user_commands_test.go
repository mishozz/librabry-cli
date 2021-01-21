package cli

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserClient struct {
	mock.Mock
}

func (m *mockUserClient) Login(username, password string) (str string, err error) {
	args := m.Called(username, password)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockUserClient) Logout(token string) (str string, err error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockUserClient) TakeBook(token, email, isbn string) (str string, err error) {
	args := m.Called(token, email, isbn)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockUserClient) ReturnBook(token, email, isbn string) (err error) {
	args := m.Called(token, email, isbn)
	if args.Get(0) == nil {
		return
	}
	return args.Error(0)
}

func (m *mockUserClient) GetAllUsers(token string) (str string, err error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *mockUserClient) GetUser(token, email string) (str string, err error) {
	args := m.Called(token, email)
	if args.Get(0) == nil {
		return
	}
	return args.Get(0).(string), args.Error(1)
}

func Test_GetAllUsers(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("GetAllUsers", mock.Anything).Return("users", nil)
			return m
		},
		expectedOutput: "users",
	}, {
		name: "error while fetching users",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("GetAllUsers", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to fetch users",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewGetUsersCmd(tt.mockUserClient(m))
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

func Test_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("GetUser", mock.Anything, mock.Anything).Return("user", nil)
			return m
		},
		expectedOutput: "user",
	}, {
		name: "error while fetching user",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("GetUser", mock.Anything, mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to fetch user with email ",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewGetUserCmd(tt.mockUserClient(m))
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

func Test_Login(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("Login", mock.Anything, mock.Anything).Return("testToken", nil)
			return m
		},
		expectedOutput: "Login succesful. Your token is: testToken",
	}, {
		name: "wrong credentials",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("Login", mock.Anything, mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to login. Check your username and password",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewLoginCmd(tt.mockUserClient(m))
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

func Test_Logout(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("Logout", mock.Anything).Return("success", nil)
			return m
		},
		expectedOutput: "success",
	}, {
		name: "wrong token",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("Logout", mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to logout. Check you token!",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewLogoutCmd(tt.mockUserClient(m))
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

func Test_TakeBook(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("TakeBook", mock.Anything, mock.Anything, mock.Anything).Return("success", nil)
			return m
		},
		expectedOutput: "success",
	}, {
		name: "error while taking book",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("TakeBook", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("error"))
			return m
		},
		expectedOutput: "Unable to take book from the library",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewTakeBookCmd(tt.mockUserClient(m))
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

func Test_ReturnBook(t *testing.T) {
	tests := []struct {
		name           string
		mockUserClient func(m *mockUserClient) *mockUserClient
		expectedOutput string
	}{{
		name: "success",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("ReturnBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			return m
		},
		expectedOutput: "Successfully returned you book",
	}, {
		name: "error while taking book",
		mockUserClient: func(m *mockUserClient) *mockUserClient {
			m.On("ReturnBook", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
			return m
		},
		expectedOutput: "Unable to return your book",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &mockUserClient{}
			getAllBooksCmd := NewReturnBookCmd(tt.mockUserClient(m))
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
