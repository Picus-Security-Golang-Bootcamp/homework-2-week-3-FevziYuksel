package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Pages     int    `json:"pages"`
	Stocks    int    `json:"stocks"`
	Price     int    `json:"price"`
	StockCode string `json:"stockCode"`
	ISBN      string `json:"ISBN"`
	Author    Author `json:"author"`
	IsInStock bool   `json:"isInStock"`
}
type BookSlice struct {
	Books []Book `json:"BookList"`
}

var (
	Command          = flag.String("command", "", "List all books")
	SearchWord       = flag.String("name", "", "Enter name of the code")
	RequestId        = flag.Int("ID", 0, "Enter id of the code")
	PurchaseQuantity = flag.Int("quantity", 0, "Enter amount that you would like to purchase")
)

func main() {
	bookList := BookSlice{}
	bookList.fillBookList()
	flag.Parse()

	switch *Command {
	case "list":
		bookList.listTheBooks()
	case "search":
		bookList.searchByName(SearchWord)
	case "get":
		bookList.findBookID(RequestId)
	case "delete":
		bookList.deleteBook(RequestId)
	case "buy":
		bookList.buyBook(RequestId, PurchaseQuantity)
	default:
		multiLine := "Valid commands are: \n" +
			"list => list all books \n" +
			"search => search and fetch books by name  \n" +
			"get => fetch a single book by id \n" +
			"delete => delete a single book by id \n" +
			"buy => buy a single book by requested quantity \n"
		fmt.Printf("%s", multiLine)
	}

}
func (bookList *BookSlice) fillBookList() {
	file, err := os.Open("BookList.json")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()
	jsonFile, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal([]byte(jsonFile), bookList)
}

func (bookList *BookSlice) listTheBooks() {
	soldBooks := 0
	for _, book := range bookList.Books {
		if book.IsInStock {
			fmt.Println(book.Name)
		} else {
			soldBooks++
		}
	}
	if soldBooks == len(bookList.Books) {
		fmt.Println("Everything is sold out!")
	}
}
func (bookList *BookSlice) searchByName(name *string) {
	for _, book := range bookList.Books {
		bookName := strings.ToLower(book.Name)
		nameLower := strings.ToLower(*name)
		if strings.Contains(bookName, nameLower) && book.IsInStock {
			fmt.Println(book.Name)
		}
	}
}
func (bookList *BookSlice) findBookID(ID *int) {
	if *ID > len(bookList.Books) {
		fmt.Println("EOF error")
		return
	}
	for _, book := range bookList.Books {
		if (book.Id == *ID) && (book.IsInStock) {
			fmt.Println(book.Name)
		}
	}
}

//Destroy Structs in Slice when deleted to completely empty list
func (bookList *BookSlice) deleteBook(ID *int) {
	if *ID > len(bookList.Books) {
		fmt.Print("EOF error")
		return
	}
	for i, book := range bookList.Books {
		if book.Id == *ID {
			(&bookList.Books[i]).IsInStock = false // == (bookList.Books[i]).IsInStock = false
			bookList.UpdateJson()
			fmt.Printf("%s is deleted\n ", book.Name)
			break
		}
	}
}
func (bookList *BookSlice) buyBook(ID *int, quantity *int) {
	for i, book := range bookList.Books {
		if book.Id == *ID {
			bookList.Books[i].Stocks -= *quantity
			bookList.UpdateJson()
			if book.Stocks == 0 {
				(&bookList.Books[i]).IsInStock = false
			} else if book.Stocks < 0 {
				fmt.Println("You cannot buy more than stock")
				continue
			}
			fmt.Printf("There are %d %s left in the stock\n", book.Stocks, book.Name)
		}
	}
}

func (bookList *BookSlice) UpdateJson() {
	byteValue, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("BookList.json", byteValue, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
