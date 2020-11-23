package controller

import (
	"crypto/rand"
	"fmt"
	"log"
	"restapi/database"
)

// Key model
type Key struct {
	ID     string `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Status bool   `json:"status,omitempty"`
}

func createAPIKey() string {
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	apiKey := fmt.Sprintf("%x", key)
	return apiKey
}

// CreateKey creates key in database
func (k *Key) CreateKey() (string, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()
	key := Key{
		Status: true,
		Key:    createAPIKey(),
	}
	err := db.Save(&key)

	if err != nil {
		return key.Key, err.Error
	}
	return "error", nil
}

// GetAllKeys returns all the keys
func (k *Key) GetAllKeys() ([]Key, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	keys := []Key{}

	err := db.Find(&keys)

	if err != nil {
		return keys, err.Error
	}
	return keys, nil
}

// DeleteKey deletes course in database
func (k *Key) DeleteKey(id string) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	kT := Key{}
	err := db.Where("id = ?", id).First(&kT)
	if err != nil {
		return err.Error
	}

	err = db.Delete(kT)

	if err != nil {
		return err.Error
	}
	return nil
}

// GetKeyByID gets the key
func (c *Key) GetKeyByID(id string) (Key, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	key := Key{}

	err := db.First(&key, id)

	if err != nil {
		return key, err.Error
	}
	if key.Key != "" {
		return key, nil
	}

	return key, nil
}

// ToggleAccess gets the key
func (c *Key) ToggleAccess(id string) (Key, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	key := Key{}

	err := db.First(&key, id)

	log.Println(key)
	newStatus := false
	if key.Status == true {
		newStatus = false
	} else {
		newStatus = true
	}

	err = db.Model(&key).Update("status", newStatus)

	if err != nil {
		return Key{}, nil
	}

	return key, nil
}
