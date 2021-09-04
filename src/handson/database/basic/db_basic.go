package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type phoneBook struct {
	ID   int
	Name string
	Tel  string
}

func main() {
	db, err := sql.Open("sqlite3", "mydata.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	fmt.Println("Database opened.")
	if err := createTable(db); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Table created.")
	if err := showTable(db); err != nil {
		log.Fatalln(err)
	}

	if err := addUserPhone(db); err != nil {
		log.Fatalln(err)
	}
}

func createTable(db *sql.DB) error {
	cmd := `
CREATE TABLE IF NOT EXISTS PHONE_BOOK(
ID   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
NAME TEXT NOT NULL,
TEL  TEXT)`
	_, err := db.Exec(cmd)
	if err != nil {
		return fmt.Errorf("createTable: %w", err)
	}
	return nil
}

func showTable(db *sql.DB) error {
	cmd := `SELECT * FROM PHONE_BOOK`
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	var pbs []phoneBook
	for rows.Next() {
		var pb phoneBook
		if err := rows.Scan(&pb.ID, &pb.Name, &pb.Tel); err != nil {
			return fmt.Errorf("showTable - Scan: %w", err)
		}
		pbs = append(pbs, pb)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("showTable - rows.Err: %w", err)
	}
	if len(pbs) == 0 {
		log.Println("No record.")
	} else {
		for _, pb := range pbs {
			log.Println(pb)
		}
	}
	return nil
}

func addUserPhone(db *sql.DB) error {
	printUsage()
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		showRecords := false
		input := in.Text()
		switch {
		case input == "d":
			cmd := "DELETE FROM PHONE_BOOK"
			result, err := db.Exec(cmd)
			if err != nil {
				return fmt.Errorf("addUserPhone - d:Exec: %w", err)
			}
			rowNum, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("addUserPhone - d:RowsAffected: %w", err)
			}
			log.Printf("%d rows deleted.\n", rowNum)
			showRecords = true
		case input == "q":
			return nil
		case strings.HasPrefix(input, "u "):
			input = strings.ReplaceAll(input, "u ", "")
			params := strings.Split(input, ",")
			if len(params) != 3 {
				log.Println("wrong format. try again.")
				continue
			}
			cmd := "UPDATE PHONE_BOOK SET NAME = ?, TEL = ? WHERE ID = ?"
			result, err := db.Exec(cmd, params[1], params[2], params[0])
			if err != nil {
				return fmt.Errorf("addUserPhone - update:Exec: %w", err)
			}
			rowNum, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("addUserPhone - update:RowsAffected: %w", err)
			}
			log.Printf("%d rows update.\n", rowNum)
			showRecords = true
		default:
			params := strings.Split(input, ",")
			if len(params) != 2 {
				log.Println("wrong format. try again.")
				continue
			}
			cmd := "INSERT INTO PHONE_BOOK(NAME, TEL) VALUES (?, ?)"
			result, err := db.Exec(cmd, params[0], params[1])
			if err != nil {
				return fmt.Errorf("addUserPhone - insert:Exec: %w", err)
			}
			rowNum, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("addUserPhone - insert:RowsAffected: %w", err)
			}
			log.Printf("%d rows inserted.\n", rowNum)
			showRecords = true
		}
		if showRecords {
			if err := showTable(db); err != nil {
				return fmt.Errorf("addUserPhone - showRecords: %w", err)
			}
		}
		printUsage()
	}
	return nil
}

func printUsage() {
	log.Println("Type [name,phone] or [u id,name,phone] or [d] then delete all or [q] then exit.")
	log.Print("> ")
}
