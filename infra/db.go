package infra

import (
	"database/sql"
	"time"

	constants "test/constants/general"
	"test/domain/general"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// IDatabase is interface for database
type Database interface {
	ConnectDB(dbAcc *general.DBDetailAccount)
	Close()

	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	// DriverName() string

	Begin() (*sql.Tx, error)
	In(query string, params ...interface{}) (string, []interface{}, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// QueryRowSqlx(query string, args ...interface{}) *sqlx.Row
	// QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type DatabaseList struct {
	Backend DatabaseType
}

type DatabaseType struct {
	Read  Database
	Write Database
}

// DBHandler - Database struct.
type DBHandler struct {
	DB  *sqlx.DB
	Err error
}

func NewDB() DBHandler {
	return DBHandler{}
}

// ConnectDB - function for connect DB.
func (d *DBHandler) ConnectDB(dbAcc *general.DBDetailAccount) {
	dbs, err := sqlx.Open("postgres", "user="+dbAcc.Username+" password="+dbAcc.Password+" sslmode="+dbAcc.SSLMode+" dbname="+dbAcc.DBName+" host="+dbAcc.URL+" port="+dbAcc.Port+" connect_timeout="+dbAcc.Timeout)
	if err != nil {
		log.Error(constants.ConnectDBFail + " | " + err.Error())
		d.Err = err
	}

	d.DB = dbs

	err = d.DB.Ping()
	if err != nil {
		// log.Error(constants.ConnectDBFail, err.Error())
		d.Err = err
	}

	// d.log.Info(constants.ConnectDBSuccess)
	d.DB.SetConnMaxLifetime(time.Duration(dbAcc.MaxLifeTime))
}

// Close - function for connection lost.
func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		// d.log.Println(constants.ClosingDBFailed + " | " + err.Error())
	} else {
		// d.log.Println(constants.ClosingDBSuccess)
	}
}

func (d *DBHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.Exec(query, args...)
	return result, err
}

func (d *DBHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := d.DB.Query(query, args...)
	return result, err
}

func (d *DBHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Select(dest, query, args...)
	return err
}

func (d *DBHandler) Get(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Get(dest, query, args...)
	return err
}

func (d *DBHandler) Rebind(query string) string {
	return d.DB.Rebind(query)
}

func (d *DBHandler) In(query string, params ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, params...)
	return query, args, err
}

func (d *DBHandler) Begin() (*sql.Tx, error) {
	return d.DB.Begin()
}

func (d *DBHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRow(query, args...)
}
