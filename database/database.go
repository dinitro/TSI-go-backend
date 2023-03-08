package database

import (
	"crypto/tls"
	"crypto/x509"
	_ "database/sql"
	"io/ioutil"
	"os"

	mysql_driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := os.Getenv("DB_USER")
	if len(user) == 0 {
		user = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	if len(address) == 0 {
		address = "localhost"
	}

	param_string := "?parseTime=true"
	cert_path := os.Getenv("DB_TLS_CERT")
	if len(cert_path) > 0 {
		rootCertPool := x509.NewCertPool()
		pem, _ := ioutil.ReadFile(cert_path)
		rootCertPool.AppendCertsFromPEM(pem)
		mysql_driver.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
		param_string += "&allowNativePasswords=true&tls=custom"
	}

	conn := (user + ":" + password + "@tcp(" + address + ":3306)/sakila" + param_string)

	db, err := gorm.Open(
		mysql.Open(conn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
