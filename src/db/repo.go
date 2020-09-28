package db

import (
	"errors"
)

// Repo : interface for CRUD actions on the datbase
type Repo interface {
	Create(u2 string, val bool, key string) (Entry, error)
	Fetch(id string) (Entry, error)
	Patch(id string, newval bool, newkey string) (Entry, error)
	Delete(id string) error
}

var repo Repo

// GetRepo : to access repo from outside the package
func GetRepo() Repo {
	return repo
}

// SetRepo : to initialize repo from outside the packgae
func SetRepo(r Repo) {
	repo = r
}

// RepoImpl : struct implementing Repo interface
type RepoImpl struct{}

//Create : adds a new entry to the database
func (r *RepoImpl) Create(id string, val bool, key string) (Entry, error) {
	e := Entry{Uuid: id, Value: val, Key: key}
	result := Database.DB.Model(&Entry{}).Create(&e)

	if result.Error != nil {
		return Entry{}, result.Error
	}

	return r.Fetch(id)
}

// Fetch is used to retrieve an entry from the database
func (r *RepoImpl) Fetch(id string) (Entry, error) {
	e := Entry{}
	result := Database.DB.Where("uuid = ?", id).First(&e)
	if result.RowsAffected != 1 {
		return Entry{}, errors.New("record not found")
	}
	return e, result.Error
}

// Delete fucntion deletes the entry from the database with macthing id
func (r *RepoImpl) Delete(id string) error {

	result := Database.DB.Where("uuid = ? ", id).Delete(&Entry{})
	if result.RowsAffected != 1 {
		return errors.New("record not found")
	}
	return result.Error
}

// Patch Function to update the entry in the database with given new values by matching id
func (r *RepoImpl) Patch(id string, newval bool, newkey string) (Entry, error) {
	// This update statement does not work properly. It cannot update a value to false as it is '0' (empty), and gorm does not update empty values
	// result := database.Model(&Entry{}).Where("uuid = ?", id).Updates(Entry{Uuid: id, Value: newval, Key: newkey})

	// Hence we need to update it after creating a string map
	result := Database.DB.Model(&Entry{}).Where("uuid = ?", id).Updates(map[string]interface{}{"Uuid": id, "Value": newval, "Key": newkey})

	if result.Error != nil {
		return Entry{}, result.Error
	}
	if result.RowsAffected != 1 {
		return Entry{}, errors.New("record not found")
	}

	return r.Fetch(id)
}
