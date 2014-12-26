package jesus

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectLocalDB(s string) (*gorm.DB, error) {
	var dns string
	db_type := "postgres"
	db_host := "localhost"
	db_port := "5432"
	db_user := "postgres"
	db_pass := "postgres"
	db_name := s
	db_sslmode := "disable"
	dns = fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)

	db, err := gorm.Open(db_type, dns)
	db.SingularTable(true)
	db.LogMode(true)
	return &db, err
}

func RemoteDB(s string) (*gorm.DB, error) {
	var dns string
	db_type := "postgres"
	db_host := "localhost"
	db_port := "5432"
	db_user := "postgres"
	db_pass := "postgres"
	db_name := s
	db_sslmode := "disable"
	dns = fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)

	db, err := gorm.Open(db_type, dns)
	db.LogMode(true)
	db.SingularTable(true)
	return &db, err
}
