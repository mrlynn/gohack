// Package storage is a wrapper package designed to simplify utilization
// of the MongoDB Go driver. It creates an interface that facilitates
// access to Get, Save, and Info functions related to the database
// storage and retrieval of shortened urls.
package storage

import "github.com/mrlynn/shortener/models"

// Storage is an interface which provides details on the utilization of
// Get, Save and Info routines
type Storage interface {
	GetURL(code string) (string, error)
	SaveURL(url string) (string, error)
	GetInfo() ([]models.Shortener, error)
	URLExists(code string) (bool, error)
}

// repository is an internal variable which instantiates the Storage interface
var (
	repository Storage
)

// SetStorage is a function that accepts a single parameter of type Storage and
// assigns the repository type to the received variable param
func SetStorage(s Storage) {
	repository = s
}

// SaveURL accepts a single url parameter and returns a string containing the
// encoded version of the url that has been saved in the database.
func SaveURL(url string) (string, error) {
	return repository.SaveURL(url)
}

// URLExists checks if a string matches the url of a saved, shortened URL.
func URLExists(url string) (bool, error) {
	return repository.URLExists(url)
}

// GetURL retrieves a URL from the database based on the encoded version of that url.
func GetURL(code string) (string, error) {
	return repository.GetURL(code)
}

// GetInfo retrieves a slice of shortened urls from the database
func GetInfo() ([]models.Shortener, error) {
	return repository.GetInfo()
}
