package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const search = "search"
const list = "list"
const buy = "buy"
const delete = "delete"

type Deletable interface {
	DeleteBook(bookList []Book, bookId int) []Book
}

type Book struct {
	Id          int
	StockNumber int
	PageNumber  int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	Author
}

type Author struct {
	Id   int
	Name string
}

// Constructor methods
func InitBook(id, stockNumber, pageNumber int, price float64, bookName, stockCode, isbn, authorName string) Book {
	book := Book{
		Id:          id,
		StockNumber: stockNumber,
		Price:       price,
		Name:        bookName,
		PageNumber:  pageNumber,
		StockCode:   stockCode,
		Isbn:        isbn,
		Author:      InitAuthor(id, authorName),
	}

	return book
}

func InitAuthor(id int, name string) Author {
	author := Author{
		Id:   id,
		Name: name,
	}
	return author
}

// Method for perform action according to user input on book list.
func PerformAction(inputList []string, bookList []Book) {
	if len(bookList) == 0 {
		fmt.Printf("\nBook list cannot found!\n\n")
		return
	}

	if len(inputList) == 1 {
		PrintMessagesToConsole()
		return
	}

	firstInput := inputList[1]

	// list all books
	if firstInput == list {
		ListBooks(bookList)
		return
	}

	// search book
	if firstInput == search {
		if len(inputList) == 2 {
			fmt.Printf("\nPlease enter the book name you want to search!\n\n")
			return
		}
		searchedBooks := SearchBook(bookList, inputList[2:])
		PrintList(searchedBooks)
		return
	}

	// buy book
	if firstInput == buy {
		if len(inputList) == 2 {
			fmt.Printf("\nPlease enter the book id and the number of books you want to buy!\n\n")
			return
		}
		PurchaseBook(bookList, inputList[2:])
		return
	}

	// delete book
	if firstInput == delete {
		if len(inputList) == 2 {
			fmt.Printf("\nPlease enter the book id you want to delete!\n\n")
			return
		}

		if len(inputList) > 3 {
			fmt.Printf("\nToo many arguments!\n\n")
			return
		}
		bookList = DeleteBook(bookList, inputList[2])
		PrintList(bookList)
		return
	}

	PrintMessagesToConsole()
}

// Method for print the books in the book list.
func ListBooks(books []Book) {
	PrintList(books)
}

// Method for search books according to given input and prints the matched books
func SearchBook(books []Book, inputList []string) []Book {
	var matchedBooks []Book
	searchedBook := strings.Join(inputList, " ")

	for _, book := range books {
		if Contains(matchedBooks, book.Id) {
			continue
		}

		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(searchedBook)) ||
			strings.Contains(strings.ToLower(book.Author.Name), strings.ToLower(searchedBook)) ||
			strings.Contains(strings.ToLower(book.StockCode), strings.ToLower(searchedBook)) {
			matchedBooks = append(matchedBooks, book)
		}
	}

	if len(matchedBooks) == 0 {
		fmt.Printf("\nCannot find any books!\n\n")
		return []Book{}
	}

	return matchedBooks
}

// Method for purchase book
func PurchaseBook(books []Book, inputList []string) {
	if len(inputList) == 1 {
		fmt.Printf("\nPlease specify the number of books you want to buy\n\n")
		return
	}

	if len(inputList) > 2 {
		fmt.Printf("\nToo many arguments\n\n")
		return
	}

	bookId := ConvertStringToInt(inputList[0])
	numberOfPurchasedBooks := ConvertStringToInt(inputList[1])

	if bookId == -1 || numberOfPurchasedBooks == -1 {
		fmt.Printf("\nInvalid Input. Please give a number!\n\n")
		return
	}

	isBookFound := false
	for _, book := range books {
		if book.Id == bookId {
			isBookFound = true
			(&book).DecreaseStockNumber(numberOfPurchasedBooks)
			break
		}
	}

	if !isBookFound {
		fmt.Printf("\nCannot find the book you want to buy\n\n")
	}
}

// Method for decrease stock number
func (book *Book) DecreaseStockNumber(stockNumber int) {
	if book.StockNumber >= stockNumber {
		fmt.Printf("\nPurchased successfully\n\n")
		book.StockNumber -= stockNumber
		return
	}
	fmt.Printf("\nStock number not enough\n\n")
}

// Method for delete a book
func DeleteBook(bookList []Book, bookId string) []Book {
	deletedBookId := ConvertStringToInt(bookId)
	if deletedBookId == -1 {
		fmt.Printf("\nInvalid Input. Please give a number!\n\n")
		return bookList
	}

	i := -1
	for index, book := range bookList {
		if book.Id == deletedBookId {
			i = index
			break
		}
	}

	if i == -1 {
		fmt.Println("Deleted before")
		return bookList
	}

	fmt.Printf("\nThe book is deleted successfully\n\n")
	return append(bookList[:i], bookList[i+1:]...)
}

// UTILS
// Method for print a list.
func PrintList(bookList []Book) {
	if len(bookList) == 0 {
		return
	}

	fmt.Println()
	for _, value := range bookList {
		fmt.Println(value.Name)
	}
	fmt.Println()
}

// Method for check if array contains the book or not.
func Contains(bookList []Book, id int) bool {
	for _, book := range bookList {
		if book.Id == id {
			return true
		}
	}
	return false
}

// Method for converting string input to int
// if it is a number then return the number, otherwise return -1
func ConvertStringToInt(input string) int {
	if result, err := strconv.Atoi(input); err == nil {
		return result
	}
	return -1
}

func ConvertStringToFloat64(input string) float64 {
	if result, err := strconv.ParseFloat(input, 64); err == nil {
		return result
	}
	return -1
}

func GenerateRandomInt(max int64) int {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(max))
	bigStr := fmt.Sprint(bigInt)
	result := ConvertStringToInt(bigStr)
	if result != -1 {
		return result
	}
	fmt.Printf("\nUnexpected error")
	return 0
}

func GenerateRandomFloat(max int64) float64 {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(max))
	bigStr := fmt.Sprint(bigInt)
	result := ConvertStringToFloat64(bigStr)
	if result != -1 {
		return result
	}
	fmt.Printf("\nUnexpected error")
	return 0
}

func GenerateRandomCode(length int) (string, error) {
	seed := "012345679"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}

// Method for print messages to user in the console.
func PrintMessagesToConsole() {
	fmt.Printf("\n--Invalid Input--\n\n")
	fmt.Println("You can use the methods below to make some actions on book list")
	fmt.Println("list: Lists the books")
	fmt.Println("search \"bookname\": searches the bookname given in the book list")
	fmt.Println("buy: you can buy books")
	fmt.Printf("delete: you can delete a book from book list\n\n")
}
