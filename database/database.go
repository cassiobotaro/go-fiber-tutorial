package database

import (
	"github.com/jinzhu/gorm"
	// sqlite support
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DBConn is a database connection singleton
	DBConn *gorm.DB
)
