package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/jedib0t/go-pretty/table"
)

var cache BookCache

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--------------------My Reading List------------------")
	for {
		cmd := GetCmdFromConsole(reader, "> ")
		cmd = strings.ToLower(cmd)
		switch cmd {
		case "nc":
			title := GetCmdFromConsole(reader, "Title: ")
			author := GetCmdFromConsole(reader, "Author: ")
			CreateBook(title, author)
		case "us":
			idStr := GetCmdFromConsole(reader, "ID: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			statusStr := GetCmdFromConsole(reader, "Status(1: Reading, 2: To Read 3: Finished): ")
			status, err := strconv.Atoi(statusStr)
			if err != nil {
				fmt.Println("Invalid status")
				continue
			}
			UpdateStatus(uint32(id), ReadingStatus(status))
		case "ll":
			PrintList(All)
		case "lr":
			PrintList(Reading)
		case "lt":
			PrintList(ToRead)
		case "lf":
			PrintList(Finished)
		case "x":
			os.WriteFile("./books.json", cache.ToJSON(), 0644)
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func PrintList(bookStatus ReadingStatus) {
	rowHeader := table.Row{"#", "Title", "Author", "Status", "Rating", "Date Added"}
	tableStruct := table.NewWriter()
	tableStruct.AppendHeader(rowHeader)
	for _, b := range cache.Books {
		if bookStatus == All || b.Status == bookStatus {
			tableStruct.AppendRow(table.Row{b.ID, b.Title, b.Author, b.Status.String(), getStars(b.Rating), b.CreatedAt.Format("2006-01-02 15:04:05")})
		}
	}
	fmt.Println(tableStruct.Render())
}

func init() {
	cache = BookCache{
		Books: make(map[uint32]Book),
	}

	fileContent, err := os.ReadFile("./books.json")
	if err != nil {
		panic(err)
	}

	var allBooks []Book

	err = json.Unmarshal(fileContent, &allBooks)
	if err != nil {
		panic(err)
	}
	for _, v := range allBooks {
		cache.AddBook(v)
	}
	atomic.AddUint32(&ID, uint32(len(allBooks)))
}

func GetCmdFromConsole(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.Trim(text, "\n")
}
