# Golang CLI Reading List Manager sample

## Introduction
The Reading List Manager is a command-line interface (CLI) application developed in Go that allows users to manage their reading list and track their reading status. This application helps users keep track of books they want to read, mark books as read, and update the reading progress of ongoing books.

## Features
- Add a book to the reading list
- Mark a book as read
- Update the reading progress of a book
- List all books in the reading list
- Filter books by reading status (unread, read, in progress)
- Delete a book from the reading list

## Installation
To install the Reading List Manager, follow these steps:

1. Ensure you have Go installed on your system.
2. Clone the repository: `git clone https://github.com/your-username/reading-list-manager.git`
3. Navigate to the project directory: `cd reading-list-manager`
4. Build the application: `go build`
5. Run the application: `./reading-list-manager`

## Usage
The Reading List Manager supports the following commands:

- `add`: Add a book to the reading list.
  - Syntax: `./reading-list-manager add -t <title> -a <author> [-p <progress>]`
  - Example: `./reading-list-manager add -t "The Great Gatsby" -a "F. Scott Fitzgerald" -p 50`

- `read`: Mark a book as read.
  - Syntax: `./reading-list-manager read -t <title>`
  - Example: `./reading-list-manager read -t "The Great Gatsby"`

- `update`: Update the reading progress of a book.
  - Syntax: `./reading-list-manager update -t <title> -p <progress>`
  - Example: `./reading-list-manager update -t "The Great Gatsby" -p 75`

- `list`: List all books in the reading list.
  - Syntax: `./reading-list-manager list [-s <status>]`
  - Example: `./reading-list-manager list -s unread`

- `delete`: Delete a book from the reading list.
  - Syntax: `./reading-list-manager delete -t <title>`
  - Example: `./reading-list-manager delete -t "The Great Gatsby"`

## Roadmap
Here's the planned roadmap for future enhancements:

- Implement data persistence using a database or file storage.
- Allow users to categorize books into genres or tags.
- Provide options to search for books by title, author, or other criteria.
- Add a command to generate reading statistics and insights.
- Develop a user-friendly graphical interface in addition to the CLI.

## Conclusion
The Reading List Manager is a convenient CLI application for managing your reading list and tracking reading status. With its easy-to-use commands, you can stay organized and keep track of your reading progress. Happy reading!
