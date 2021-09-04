package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type cashBalance struct {
	ID     int
	Name   string
	Amount int
}

type moneyTransaction struct {
	From       string
	To         string
	SendAmount int
}

func main() {
	db, err := sql.Open("sqlite3", "mydatatx.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	fmt.Println("Database opened.")
	if err := createTable(db); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Table created.")
	if err := initTable(db); err != nil {
		log.Fatalln(err)
	}
	if err := showTable(db); err != nil {
		log.Fatalln(err)
	}
	transaction := &moneyTransaction{
		From:       "user1",
		To:         "user2",
		SendAmount: 10,
	}
	log.Printf("sendMoney %d amount from %s to %s.",
		transaction.SendAmount, transaction.From, transaction.To)
	if err := sendMoneyTx(db, transaction); err != nil {
		log.Fatalln(err)
	}
	if err := showTable(db); err != nil {
		log.Fatalln(err)
	}
}

func createTable(db *sql.DB) error {
	cmd := `
CREATE TABLE IF NOT EXISTS CASH_BALANCE(
ID     INTEGER NOT NULL PRIMARY KEY,
NAME   TEXT NOT NULL,
AMOUNT INTEGER NOT NULL)`
	_, err := db.Exec(cmd)
	if err != nil {
		return fmt.Errorf("createTable: %w", err)
	}
	return nil
}

func initTable(db *sql.DB) error {
	cmd := `DELETE FROM CASH_BALANCE`
	_, err := db.Exec(cmd)
	if err != nil {
		return fmt.Errorf("initTable - delete: %w", err)
	}
	cashBalances := []*cashBalance{
		&cashBalance{
			ID:     1,
			Name:   "user1",
			Amount: 100,
		},
		&cashBalance{
			ID:     2,
			Name:   "user2",
			Amount: 20,
		},
	}
	cmd = "INSERT INTO CASH_BALANCE(ID, NAME, AMOUNT) VALUES (?, ?, ?)"
	for _, cb := range cashBalances {
		if _, err := db.Exec(cmd, cb.ID, cb.Name, cb.Amount); err != nil {
			return fmt.Errorf("initTable - insert: %w", err)
		}
	}
	return nil
}

func showTable(db *sql.DB) error {
	cmd := `SELECT * FROM CASH_BALANCE`
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	var cbs []cashBalance
	for rows.Next() {
		var cb cashBalance
		if err := rows.Scan(&cb.ID, &cb.Name, &cb.Amount); err != nil {
			return fmt.Errorf("showTable - Scan: %w", err)
		}
		cbs = append(cbs, cb)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("showTable - rows.Err: %w", err)
	}
	if len(cbs) == 0 {
		log.Println("No record.")
	} else {
		for _, pb := range cbs {
			log.Println(pb)
		}
	}
	return nil
}

func sendMoneyTx(db *sql.DB, moneyTx *moneyTransaction) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("sendMoneyTx - Begin %w", err)
	}
	cmd := "UPDATE CASH_BALANCE SET AMOUNT = AMOUNT + ? WHERE NAME = ?"
	if _, err := tx.Exec(cmd, -moneyTx.SendAmount, moneyTx.From); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("sendMoneyTx - update From Rollback %w", err)
		}
		return fmt.Errorf("sendMoneyTx - update From %w", err)
	}
	if _, err := tx.Exec(cmd, moneyTx.SendAmount, moneyTx.To); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("sendMoneyTx - update To Rollback %w", err)
		}
		return fmt.Errorf("sendMoneyTx - update To %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("sendMoneyTx - Commit %w", err)
	}
	return nil
}
