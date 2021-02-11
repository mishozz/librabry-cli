package cli

import "github.com/stretchr/testify/mock"

func ExampleNewSaveBookCmd() {
	mock := func(m *mockBookClient) *mockBookClient {
		m.On("SaveBook", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockBookClient{}
	saveBookCmd := NewSaveBookCmd(mock(m))
	saveBookCmd.Execute()
	// Output:
	// success
}

func ExampleNewDeleteBookCmd() {
	mock := func(m *mockBookClient) *mockBookClient {
		m.On("Delete", mock.Anything, mock.Anything).Return(nil)
		return m
	}
	m := &mockBookClient{}
	deleteBookCmd := NewDeleteBookCmd(mock(m))
	deleteBookCmd.Execute()
	// Output:
	// Book with isbn  successfully deleted
}

func ExampleNewGetBookCmd() {
	mock := func(m *mockBookClient) *mockBookClient {
		m.On("GetBook", mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockBookClient{}
	getBookCmd := NewGetBookCmd(mock(m))
	getBookCmd.Execute()
	// Output:
	// success
}

func ExampleNewGetUserCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("GetUser", mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	getUserCmd := NewGetUserCmd(mock(m))
	getUserCmd.Execute()
	// Output:
	// success
}

func ExampleNewGetUsersCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("GetAllUsers", mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	getUserCmd := NewGetUsersCmd(mock(m))
	getUserCmd.Execute()
	// Output:
	// success
}

func ExampleNewLoginCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("Login", mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	loginCmd := NewLoginCmd(mock(m))
	loginCmd.Execute()
	// Output:
	// success
}

func ExampleNewLogoutCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("Logout", mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	loginCmd := NewLogoutCmd(mock(m))
	loginCmd.Execute()
	// Output:
	// success
}

func ExampleNewTakeBookCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("TakeBook", mock.Anything, mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	loginCmd := NewTakeBookCmd(mock(m))
	loginCmd.Execute()
	// Output:
	// success
}

func ExampleNewReturnBookCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("ReturnBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		return m
	}
	m := &mockUserClient{}
	loginCmd := NewReturnBookCmd(mock(m))
	loginCmd.Execute()
	// Output:
	// Successfully returned your book
}

func ExampleNewRegisterCmd() {
	mock := func(m *mockUserClient) *mockUserClient {
		m.On("Register", mock.Anything, mock.Anything).Return("success", nil)
		return m
	}
	m := &mockUserClient{}
	loginCmd := NewRegisterCmd(mock(m))
	loginCmd.Execute()
	// Output:
	// success
}
