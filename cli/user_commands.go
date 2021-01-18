package cli

import (
	"fmt"

	"github.com/mishozz/library-cli/client"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with username and password",
	Long:  "Login with username and password",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")

		token, err := client.User.Login(email, "")
		if err != nil {
			fmt.Printf("Unable to login. Check your username and password")
		}

		fmt.Println("Login succesful.")
		fmt.Printf("Your token is: %s", token)
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  "Logout from your account",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		msg, err := client.User.Logout(token)
		if err != nil {
			fmt.Printf("Unable to logout. Check you token!")
		}
		fmt.Printf(msg)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)

	loginCmd.Flags().StringP("email", "e", "", "Set your email")
	loginCmd.MarkFlagRequired("email")

	logoutCmd.Flags().StringP("token", "t", "", "Your jwt token")
	logoutCmd.MarkFlagRequired("token")
}
