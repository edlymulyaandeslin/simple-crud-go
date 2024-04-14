package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var books []Book

var fileName string = "data.csv"

func main() {
	// createFile(fileName)
	err := loadDataFromCSV(fileName)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 14), "Book Data Management", strings.Repeat("=", 14))
		fmt.Println("1. View All Books")
		fmt.Println("2. Add New Book")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Exit")
		fmt.Println(strings.Repeat("=", 50))

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your choice : ")
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			err := viewAllBooks()
			printError(err)
		case 2:
			err := addNewBook()
			printError(err)
		case 3:
			err := updateBook()
			printError(err)
		case 4:
			err := deleteBook()
			printError(err)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}

func printError(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

// func createFile(fileName string) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("file", fileName, "berhasil dibuat.")
// 	}
// 	defer file.Close()
// }

func addNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Book Details :")
	fmt.Print("Book Id : ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title : ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Release Year : ")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Book Pages : ")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Are you sure want to add this book (y/n)? ")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		_, err := findBookById(newBook.Id)
		if err != nil {
			books = append(books, newBook)
		} else {
			return fmt.Errorf("Book with id: %d already exist", newBook.Id)
		}

		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}
		fmt.Println("Book added successfully.")

	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not saved")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		addNewBook()
	}

	return nil
}

func saveDataToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"
		file.WriteString(row)
	}

	return nil
}

func loadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		// fmt.Println(record)
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			ReleaseYear: record[3],
			Pages:       pages,
		}
		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}

	return nil
}

func viewAllBooks() error {
	if len(books) == 0 {
		return fmt.Errorf("no books available")
	}

	for i, book := range books {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Book -", i+1)
		fmt.Println("Book Id :", book.Id)
		fmt.Println("Book Title :", book.Title)
		fmt.Println("Book Author :", book.Author)
		fmt.Println("Release Year :", book.ReleaseYear)
		fmt.Println("Pages :", book.Pages)
		fmt.Println(strings.Repeat("=", 50))
	}

	return nil
}

func updateBook() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book Id to Update : ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	book, err := findBookById(bookId)
	if err != nil {
		return err
	}

	var updatedBook Book

	fmt.Println("Enter Updated Book Details :")
	fmt.Print("Book Title : ")
	scanner.Scan()
	updatedBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	updatedBook.Author = scanner.Text()

	fmt.Print("Release Year : ")
	scanner.Scan()
	updatedBook.ReleaseYear = scanner.Text()

	fmt.Print("Pages : ")
	scanner.Scan()
	pages, _ := strconv.Atoi(scanner.Text())
	updatedBook.Pages = pages

	fmt.Print("Are you sure want to update this book (y/n)? ")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		for i := range books {
			if books[i].Id == book.Id {
				if strings.TrimSpace(updatedBook.Title) != "" {
					books[i].Title = updatedBook.Title
				}
				if strings.TrimSpace(updatedBook.Author) != "" {
					books[i].Author = updatedBook.Author
				}
				if strings.TrimSpace(updatedBook.ReleaseYear) != "" {
					books[i].ReleaseYear = updatedBook.ReleaseYear
				}
				if updatedBook.Pages != 0 {
					books[i].Pages = updatedBook.Pages
				}
			}
		}

		fmt.Println("Book updated successfully.")
		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not updated")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		updateBook()
	}

	return nil
}

func deleteBook() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book ID to delete : ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Are you sure want to delete this book (y/n)? ")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		book, err := findBookById(bookId)
		if err != nil {
			return err
		} else {
			for i := range books {
				if books[i].Id == book.Id {
					books = append(books[:i], books[i+1:]...)
					break
				}
			}
		}
		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}
		fmt.Println("Book deleted successfully.")
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not deleted")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		deleteBook()
	}

	return nil
}

func findBookById(id int) (Book, error) {
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("id: %d not found", id)
}
