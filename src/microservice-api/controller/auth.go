package controller

import (
	"log"
	"restapi/database"
)

// AuthenticateKey checks if key exists in the database
func AuthenticateKey(key string) bool {
	db := database.GetInstance().GetConnection()
	defer db.Close()
	k := Key{}
	db.Where("`keys`.`key` = ?", key).First(&k)
	log.Println(k)
	if k.Key == key {
		return true
	}
	return false
}
