package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)


var Db *sql.DB
var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)


func Init() *sql.DB{
	err := godotenv.Load()

	if err != nil{
		fmt.Println("err")

	}
	

	

	user := os.Getenv("MYSQL_USER")

	pass:= os.Getenv("MYSQL_PASSWORD")

	database := os.Getenv("MYSQL_DATABASE")

	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pass, database)

	return open(path)
}

    func open(path string) *sql.DB{


		Db, err = sql.Open("mysql", path); 
		if err != nil {
		log.Fatalln("Db open error:", err)

		}

		err = Db.Ping()
		if err != nil{
			log.Fatalln(err)
		}
		log.Println("DB_ok")


/*
		cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
        uuid VARCHAR(255) NOT NULL UNIQUE,
        name VARCHAR(255),
        email VARCHAR(255),
        password VARCHAR(255),
        created_at DATETIME)`, tableNameUser)
    
		Db.Exec(cmdU)

	   cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		content VARCHAR(255),
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255),
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

	*/
	/*
	
	err = Db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("db connected!!")
*/

		return Db

	}


	func createdUUID()(uuidobj uuid.UUID){
		uuidobj, _ = uuid.NewUUID()
		return uuidobj
	}

	func Encrypt(plaintext string) (cryptext string){
		cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
		return cryptext
	}



	

	




	