package db

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	insertQuery = "INSERT INTO `entries` (`uuid`,`value`,`key`) VALUES (?,?,?)"
	fetchQuery  = "SELECT * FROM `entries`  WHERE (uuid = ?) LIMIT 1"
	updateQuery = "UPDATE `entries` SET `key` = ?, `uuid` = ?, `value` = ? WHERE (uuid = ?)"
	deleteQuery = "DELETE FROM `entries` WHERE (uuid = ? )"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	d, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	Database = Db{DB: d}
	SetRepo(&repoImpl)

	rows := sqlmock.
		NewRows([]string{"id", "value", "key"}).
		AddRow("some uuid", true, "catto")

	mock.ExpectBegin()
	mock.ExpectExec(insertQuery).WithArgs("some uuid", true, "catto").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	mock.ExpectQuery(fetchQuery).WithArgs("some uuid").WillReturnRows(rows)

	if _, err = repo.Create("some uuid", true, "catto"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPatch(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	d, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	Database = Db{DB: d}
	SetRepo(&repoImpl)

	rows := sqlmock.
		NewRows([]string{"id", "value", "key"}).
		AddRow("some uuid", true, "fatto")

	mock.ExpectBegin()
	mock.ExpectExec(updateQuery).WithArgs("catto", "some uuid", true, "some uuid").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	mock.ExpectQuery(fetchQuery).WithArgs("some uuid").WillReturnRows(rows)

	if _, err = repo.Patch("some uuid", true, "catto"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	d, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	Database = Db{DB: d}
	SetRepo(&repoImpl)

	mock.ExpectBegin()
	mock.ExpectExec(deleteQuery).WithArgs("some uuid").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	if err = repo.Delete("some uuid"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
