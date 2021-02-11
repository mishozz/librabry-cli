package cli

import (
	"fmt"

	"github.com/mishozz/library-cli/client"
	"github.com/spf13/cobra"
)

// NewLoginCmd return cobra command for login
func NewLoginCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login with username and password",
		Long:  "Login with username and password",
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")

			respString, err := client.Login(email, password)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to login. Check your username and password")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

// NewLogoutCmd return cobra command for logout
func NewLogoutCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout",
		Long:  "Logout from your account",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			respString, err := client.Logout(token)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to logout. Check you token!")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

// NewTakeBookCmd return cobra command for taking a book
func NewTakeBookCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "take",
		Short: "Take book",
		Long:  "Take book from the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			email, _ := cmd.Flags().GetString("email")
			isbn, _ := cmd.Flags().GetString("isbn")

			respString, err := client.TakeBook(token, email, isbn)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to take book from the library")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

// NewReturnBookCmd return cobra command for returning a book
func NewReturnBookCmd(userClient client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "return",
		Short: "Return book",
		Long:  "Return book in the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			email, _ := cmd.Flags().GetString("email")
			isbn, _ := cmd.Flags().GetString("isbn")

			err := userClient.ReturnBook(token, email, isbn)
			if err != nil {
				if err == client.UnauthorizedErr {
					fmt.Fprintf(cmd.OutOrStdout(), "You need to be authorized to access this route")
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Unable to return your book")
				}
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "Successfully returned your book")
			}
		},
	}
}

// NewGetUsersCmd return cobra command for getting all of the users
func NewGetUsersCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "get-all-users",
		Short: "Get all users",
		Long:  "Get all users of the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")

			respString, err := client.GetAllUsers(token)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to fetch users")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

// NewGetUserCmd return cobra command for getting a user
func NewGetUserCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "get-user",
		Short: "Get user of the library",
		Long:  "Get user of the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			email, _ := cmd.Flags().GetString("email")

			respString, err := client.GetUser(token, email)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to fetch user with email %s", email)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

// NewRegisterCmd return cobra command for registering a user
func NewRegisterCmd(client client.UserClient) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Register user in the library",
		Long:  "Register user in the library",
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")

			respString, err := client.Register(email, password)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to register. Try again!")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}
func init() {
	loginCmd := NewLoginCmd(client.User)
	logoutCmd := NewLogoutCmd(client.User)
	takeBookCmd := NewTakeBookCmd(client.User)
	returnBookCmd := NewReturnBookCmd(client.User)
	getUsersCmd := NewGetUsersCmd(client.User)
	getUserCmd := NewGetUserCmd(client.User)
	registerCmd := NewRegisterCmd(client.User)

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(takeBookCmd)
	rootCmd.AddCommand(returnBookCmd)
	rootCmd.AddCommand(getUsersCmd)
	rootCmd.AddCommand(getUserCmd)
	rootCmd.AddCommand(registerCmd)

	loginCmd.Flags().StringP("email", "e", "", "Set your email")
	loginCmd.MarkFlagRequired("email")
	loginCmd.Flags().StringP("password", "p", "", "Enter you password")
	loginCmd.MarkFlagRequired("password")

	logoutCmd.Flags().StringP("token", "t", "", "Your jwt token")
	logoutCmd.MarkFlagRequired("token")

	takeBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	takeBookCmd.Flags().StringP("email", "e", "", "Set your email")
	takeBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	takeBookCmd.MarkFlagRequired("token")
	takeBookCmd.MarkFlagRequired("email")
	takeBookCmd.MarkFlagRequired("isbn")

	returnBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	returnBookCmd.Flags().StringP("email", "e", "", "Set your email")
	returnBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	returnBookCmd.MarkFlagRequired("token")
	returnBookCmd.MarkFlagRequired("email")
	returnBookCmd.MarkFlagRequired("isbn")

	getUsersCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getUsersCmd.MarkFlagRequired("token")

	getUserCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getUserCmd.Flags().StringP("email", "e", "", "Set your email")
	getUserCmd.MarkFlagRequired("token")
	getUserCmd.MarkFlagRequired("email")

	registerCmd.Flags().StringP("email", "e", "", "Set your email")
	registerCmd.MarkFlagRequired("email")
	registerCmd.Flags().StringP("password", "p", "", "Set you password")
	registerCmd.MarkFlagRequired("password")
}
