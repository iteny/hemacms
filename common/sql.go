package common

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Sql() *sqlx.DB {
	return DB
}
func (c *BaseCtrl) Sql() *sqlx.DB {
	return DB
}
func init() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "./sql/hemacms.db")
	Log().CheckErr("Database Error", err)
	// db.MustExec(schema)

	// tx := DB.MustBegin()
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// tx.Commit()
	// sx := DB.MustBegin()
	// sx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// sx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// sx.Commit()

}
