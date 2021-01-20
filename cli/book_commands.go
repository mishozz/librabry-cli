package cli

import (
	"fmt"

	"github.com/mishozz/library-cli/client"
	"github.com/spf13/cobra"
)

func NewGetBooksCmd(client client.BookClient) *cobra.Command {
	return &cobra.Command{
		Use:   "get-all",
		Short: "Get books from the library",
		Long:  `Get specific book or all books from the library`,
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")

			respString, err := client.GetAllBooks(token)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to fetch books from library")
			}
			fmt.Fprintf(cmd.OutOrStdout(), respString)
		},
	}
}

func NewGetBookCmd(client client.BookClient) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get specific book from the library",
		Long:  "Get specific book from the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			isbn, _ := cmd.Flags().GetString("isbn")

			respString, err := client.GetBook(token, isbn)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to fetch book with isbn %s from library", isbn)
			}
			fmt.Fprintf(cmd.OutOrStdout(), respString)
		},
	}
}

func NewSaveBookCmd(client client.BookClient) *cobra.Command {
	return &cobra.Command{
		Use:   "save",
		Short: "Save book",
		Long:  "Save book in the library with the provided properties (isbn,title,author,units)",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			title, _ := cmd.Flags().GetString("title")
			author, _ := cmd.Flags().GetString("author")
			units, _ := cmd.Flags().GetInt("units")
			isbn, _ := cmd.Flags().GetString("isbn")

			respString, err := client.SaveBook(token, isbn, title, author, uint(units))
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Unable to save book with isbn %s", isbn)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), respString)
			}
		},
	}
}

func NewDeleteBookCmd(bookClient client.BookClient) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete specific book from the library",
		Long:  "Delete specific book from the library",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			isbn, _ := cmd.Flags().GetString("isbn")

			err := bookClient.Delete(token, isbn)
			if err != nil {
				if err == client.UnauthorizedErr {
					fmt.Fprintf(cmd.OutOrStdout(), "You need to be authorized to access this route")
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Unable to delete book with isbn %s", isbn)
				}
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "Book with isbn %s successfully deleted", isbn)
			}
		},
	}
}

func init() {
	getBooksCmd := NewGetBooksCmd(client.Books)
	getBookCmd := NewGetBookCmd(client.Books)
	saveBookCmd := NewSaveBookCmd(client.Books)
	deleteBookCmd := NewDeleteBookCmd(client.Books)

	rootCmd.AddCommand(getBooksCmd)
	rootCmd.AddCommand(getBookCmd)
	rootCmd.AddCommand(saveBookCmd)
	rootCmd.AddCommand(deleteBookCmd)

	getBooksCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getBooksCmd.MarkFlagRequired("token")

	getBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	getBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getBookCmd.MarkFlagRequired("isbn")
	getBookCmd.MarkFlagRequired("token")

	deleteBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	deleteBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	deleteBookCmd.MarkFlagRequired("isbn")
	deleteBookCmd.MarkFlagRequired("token")

	saveBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	saveBookCmd.Flags().StringP("title", "n", "", "Title of the book")
	saveBookCmd.Flags().StringP("author", "a", "", "Author of the book")
	saveBookCmd.Flags().IntP("units", "u", 0, "Available units")
	saveBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	saveBookCmd.MarkFlagRequired("isbn")
	saveBookCmd.MarkFlagRequired("token")
	saveBookCmd.MarkFlagRequired("title")
	saveBookCmd.MarkFlagRequired("author")
	saveBookCmd.MarkFlagRequired("units")

}
