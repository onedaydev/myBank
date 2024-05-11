package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type AccountData struct {
	AccountID string
	OwnerName string
	Balance   float64
	Currency  string
}

type DBConfig struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBHost     string `json:"dbHost"`
	DBPort     string `json:"dbPort"`
	DBName     string `json:"dbName"`
}

func ConnectToDB() (*sql.DB, error) {
	// configPath := "/../config/config.json"
	configPath := "../../config/config.json" // Use for server_test
	configFile, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	bytes, _ := io.ReadAll(configFile)

	var config DBConfig

	json.Unmarshal(bytes, &config)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateAccount(db *sql.DB, info AccountData) error {
	query := `INSERT INTO accounts (account_id, owner_name, balance, currency) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, info.AccountID, info.OwnerName, info.Balance, info.Currency)
	if err != nil {
		return fmt.Errorf("CreateAccount error: %v", err)
	}
	return nil
}

func GetRow(db *sql.DB, requestID string) (AccountData, error) {
	query := `SELECT account_id, owner_name, balance, currency FROM accounts where account_id = ?`

	var data AccountData

	err := db.QueryRow(query, requestID).Scan(&data.AccountID, &data.OwnerName, &data.Balance, &data.Currency)
	if err != nil {
		return AccountData{}, fmt.Errorf("QueryRow error: %v", err)
	}

	return data, nil
}

// func UpdateBalance(db *sql.DB) {
// 	query := `UPDATE accounts SET balance = ? WHERE account_id = ?`

// 	var data AccountData

// 	err := db.QueryRow(query, requestID).Scan(&data.AccountID, &data.OwnerName, &data.Balance, &data.Currency)
// 	if err != nil {
// 		return AccountData{}, fmt.Errorf("QueryRow error: %v", err)
// 	}
// 	return data, nil
// }

func UpdateOwnerName(db *sql.DB, requestID string, NewName string) (AccountData, error) {
	query := `UPDATE accounts SET OwnerName = ? WHERE account_id = ?`

	var data AccountData
	err := db.QueryRow(query, NewName, requestID).Scan(&data.AccountID, &data.OwnerName, &data.Balance, &data.Currency)
	if err != nil {
		return AccountData{}, fmt.Errorf("QueryRow error: %v", err)
	}
	return data, nil
}

func DeleteRow(db *sql.DB, requestID string) error {
	query := `DELETE FROM accounts WHERE account_id = ?`
	_, err := db.Exec(query, requestID)
	if err != nil {
		return fmt.Errorf("QueryRow error: %v", err)
	}
	return nil
}
