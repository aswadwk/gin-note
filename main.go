package main

import (
	"aswadwk/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.DBConnect()
)

func main() {
	// ...
}