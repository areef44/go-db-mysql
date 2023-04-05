package gomysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-db?parseTime=true")
	if err != nil {
		panic(err)
	}

	//database pooling
	db.SetMaxIdleConns(10)                  //minimal connection
	db.SetMaxOpenConns(100)                 //maximal connection
	db.SetConnMaxIdleTime(5 * time.Minute)  // jika dalam lima menit ga ada request maka close connection
	db.SetConnMaxLifetime(60 * time.Minute) //jika sudah 1 jam perbarui semua koneksinya/dibikin baru

	return db
}
