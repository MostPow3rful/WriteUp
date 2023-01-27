package sql

import (
	"github.com/JesusKian/WriteUp/src/config"
	"github.com/JesusKian/WriteUp/src/discord"
	"github.com/JesusKian/WriteUp/src/structure"
	"github.com/JesusKian/WriteUp/src/telegram"
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"database/sql"
	"encoding/json"
	"os"
)

var (
	rows     *sql.Rows = &sql.Rows{}
	Database *sql.DB   = &sql.DB{}
	err      error     = nil
)

func ConnectToSqlDatabase() {
	var (
		secretFile *os.File          = &os.File{}
		scanner    *bufio.Scanner    = &bufio.Scanner{}
		MySqlData  *structure.Secret = &structure.Secret{}
		secretData string            = ""
	)

	secretFile, err = os.Open("./json/Secret.json")
	if err != nil {
		config.SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open ./json/Secret.json")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}
	defer secretFile.Close()

	scanner = bufio.NewScanner(secretFile)
	for scanner.Scan() {
		if scanner.Text() == " " {
			continue
		}
		secretData += scanner.Text()
	}
	json.Unmarshal([]byte(secretData), MySqlData)

	err = scanner.Err()
	if err != nil {
		config.SetLog("E", "config.ConnectToSqlDatabase() -> Unknow Error From bufio.Scanner()")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}

	Database, err = sql.Open("mysql", MySqlData.Username+":"+MySqlData.Password+"@tcp(0.0.0.0:3306)/WriteUp")
	if err != nil {
		config.SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open WriteUp Database")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}
}

func DatabaseStatus() {
	err = Database.Ping()
	if err != nil {
		config.SetLog("E", "config.DatabaseStatus() -> MySQL Dosn't Response")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}
	config.SetLog("I", "config.DatabaseStatus() -> MySQL Is Ready To Use")
}

func SendWriteUp(_title, _link, _pubDate string) {
	discord.Send(_title, _link, _pubDate)
	telegram.Send(_title, _link, _pubDate)
}

func AddWriteUp(_title, _link, _pubDate string) {
	_, err = Database.Query(`
	INSERT INTO data
	(Title, Link, PublishDate)
	VALUES
	(?, ?, ?)`, _title, _link, _pubDate)

	if err != nil {
		config.SetLog("E", "sql.AddWriteUp() -> Couldn't Add WriteUp To Database")
		config.SetLog("D", err.Error())
	}

	SendWriteUp(_title, _link, _pubDate)
}

func CheckWriteUp(_title, _link, _pubDate string) {
	var temp string = ""

	rows, err = Database.Query(`SELECT Link FROM data where Link=?`, _link)
	if err != nil {
		config.SetLog("E", "sql.CheckWriteUp() -> Couldn't Execute Query")
		config.SetLog("D", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&temp)
		if err != nil {
			config.SetLog("E", "sql.CheckWriteUp() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	if temp == "" {
		AddWriteUp(_title, _link, _pubDate)
	}
}

func init() {
	ConnectToSqlDatabase()
	DatabaseStatus()
}
