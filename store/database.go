package store

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	// pq driver is used by sqlx
	_ "github.com/lib/pq"
)

// DbConfig is the data structure containing necessary data to connect to db
type DbConfig struct {
	addr     string
	user     string
	port     string
	password string
	dbName   string
	sslMode  string
}

var (
	db       *sqlx.DB
	dbConfig *DbConfig
)

// GetConnectionQuery returns formatted querystring according to viper configurations
func GetConnectionQuery() string {
	query := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s", dbConfig.addr, dbConfig.user, dbConfig.port, dbConfig.password, dbConfig.dbName, dbConfig.sslMode)
	return query
}

// CreateDBInstance return a new DB instance
func CreateDBInstance() (*sqlx.DB, error) {
	var (
		driver = viper.GetString("database.driver")
	)

	dbConfig = initConfig()

	query := GetConnectionQuery()

	db, err := sqlx.Connect(driver, query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// InitDatabase initializes db instance and executes schema query
func InitDatabase() {
	db, err := CreateDBInstance()

	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := ioutil.ReadFile("store/schema.sql")

	schema := string(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	db.MustExec(schema)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func initConfig() *DbConfig {
	var (
		addr     = viper.GetString("database.addr")
		user     = viper.GetString("database.user")
		port     = viper.GetString("database.port")
		password = viper.GetString("database.password")
		dbName   = viper.GetString("database.dbName")
		sslMode  = viper.GetString("database.sslMode")
	)
	return &DbConfig{addr, user, port, password, dbName, sslMode}
}
