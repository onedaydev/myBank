package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
)

type DBConfig struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBHost     string `json:"dbHost"`
	DBPort     string `json:"dbPort"`
	DBName     string `json:"dbName"`
}

func ConnectToDB() (*sql.DB, error) {
	configFile, err := os.Open("../confing/config.json")
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

func CreateAccount(db *sql.DB, info *pb.AccountInfo) error {
	query := `INSERT INTO accounts (accountId, owner_name, balance, currency) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, info.AccountId, info.OwnerName, info.Balance, info.Currency)
	if err != nil {
		return fmt.Errorf("CreateAccount error: %v", err)
	}
	return nil
}
