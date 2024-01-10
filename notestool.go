package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Note struct {
	ID   string
	Text string
}

type Notebook struct {
	Name  string
	Notes []Note
}

const fileExtension = ".txt"

func loadNotebook(name string) Notebook {
	fileName := name + fileExtension
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	notes := []Note{}
	id := 1
	for scanner.Scan() {
		notes = append(notes, Note{ID: fmt.Sprintf("%03d", id), Text: scanner.Text()})
		id++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Notebook{Name: name, Notes: notes}
}

func saveNotebook(notebook Notebook) {
	fileName := notebook.Name + fileExtension
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note := range notebook.Notes {
		_, err := writer.WriteString(note.Text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func showNotes(notebook Notebook) {
	fmt.Printf("\nNotes in %s:\n", notebook.Name)
	for _, note := range notebook.Notes {
		fmt.Printf("%s %s\n", note.ID, note.Text)
	}
}

func addNote(notebook *Notebook, text string) {
	note := Note{ID: fmt.Sprintf("%03d", len(notebook.Notes)+1), Text: text}
	notebook.Notes = append(notebook.Notes, note)
	saveNotebook(*notebook)
}

func deleteNote(notebook *Notebook, i int) {
	if i == 0 {
		notebook.Notes = notebook.Notes[1:]
		saveNotebook(*notebook)
	} else if i > 0 && i < len(notebook.Notes) {
		notebook.Notes = append(notebook.Notes[:i], notebook.Notes[i+1:]...)
		saveNotebook(*notebook)
	} else if i != 0 {
		fmt.Println("Invalid index")
	}
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [TAG]")
		return
	}

	notebook := loadNotebook(os.Args[1])

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			showNotes(notebook)
		case "2":
			fmt.Println("\nEnter the note text:")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			addNote(&notebook, text)
		case "3":
			fmt.Println("\nEnter the note index:")
			input, _ := reader.ReadString('\n')
			index, _ := strconv.Atoi(strings.TrimSpace(input))
			deleteNote(&notebook, index-1)
		case "4":
			return
		default:
			fmt.Println("Invalid operation.")
		}
	}
}
