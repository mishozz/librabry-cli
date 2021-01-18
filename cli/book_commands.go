package cli

import (
	"fmt"

	"github.com/mishozz/library-cli/client"
	"github.com/spf13/cobra"
)

var getBooksCmd = &cobra.Command{
	Use:   "get-all",
	Short: "Get books from the library",
	Long:  `Get specific book or all books from the library`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")

		books, err := client.Books.GetAllBooks(token)
		if err != nil {
			fmt.Printf("Unable to fetch books from library")

		}
		fmt.Printf(books)
	},
}

var getBookCmd = &cobra.Command{
	Use:   "get",
	Short: "Get specific book from the library",
	Long:  "Get specific book from the library",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		isbn, _ := cmd.Flags().GetString("isbn")

		book, err := client.Books.GetBook(token, isbn)
		if err != nil {
			fmt.Printf("Unable to fetch book with isbn %s from library", isbn)

		}
		fmt.Printf(book)
	},
}

var saveBookCmd = &cobra.Command{
	Use:   "save",
	Short: "Save book",
	Long:  "Save book in the library with the provided properties (isbn,title,author,units)",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		title, _ := cmd.Flags().GetString("title")
		author, _ := cmd.Flags().GetString("author")
		units, _ := cmd.Flags().GetInt("units")
		isbn, _ := cmd.Flags().GetString("isbn")

		book, err := client.Books.SaveBook(token, isbn, title, author, units)
		if err != nil {
			fmt.Printf("Unable to save book with isbn %s", isbn)

		}
		fmt.Printf(book)
	},
}

func init() {
	rootCmd.AddCommand(getBooksCmd)
	rootCmd.AddCommand(getBookCmd)
	rootCmd.AddCommand(saveBookCmd)

	getBooksCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getBooksCmd.MarkFlagRequired("token")

	getBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	getBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getBookCmd.MarkFlagRequired("isbn")
	getBookCmd.MarkFlagRequired("token")

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
