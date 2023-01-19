package config

import (
	"github.com/JesusKian/WriteUp/src/structure"
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

var (
	err      error   = nil
	PWD      string  = ""
	Database *sql.DB = &sql.DB{}

	temp    []string = []string{}
	counter int      = 0
	logFile *os.File = &os.File{}

	// Log Info
	flags = log.Lshortfile
	// Sturct For Information Log
	infoLog = log.New(os.Stdout, "[?] Information -> ", flags)
	// Sturct For Warning Log
	warnLog = log.New(os.Stdout, "[*] Warning -> ", flags)
	// Sturct For Error Log
	errorLog = log.New(os.Stdout, "[!] Error -> ", flags)
	// Struct For Default Log
	defaultLog = log.New(os.Stderr, "[#] Default Log -> ", flags)
)

func SetLog(_logType string, _msg string) {
	PWD, err = os.Getwd()
	if err != nil {
		errorLog.Println("config.SetLog() -> Couldn't Get Output Of 'os.Getwd()'")
		os.Exit(1)
	}

	logFile, err = os.OpenFile(PWD+"/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errorLog.Println(`
		Invalid Directory -> /log
		Invalid File -> /log/log.log
		`)
		os.Exit(1)
	}

	switch _logType {
	case "I":
		infoLog.SetOutput(logFile)
		infoLog.Println(_msg)

	case "W":
		warnLog.SetOutput(logFile)
		warnLog.Println(_msg)

	case "E":
		errorLog.SetOutput(logFile)
		errorLog.Println(_msg)

	case "D":
		defaultLog.SetOutput(logFile)
		defaultLog.Println(_msg)

	default:
		errorLog.SetOutput(logFile)
		errorLog.Printf("SetLog() -> Trying To Add Log Without Valid LogType '%s'", _logType)
	}

	logFile.Close()
}

func ConnectToSqlDatabase() {
	var (
		secretFile *os.File          = &os.File{}
		scanner    *bufio.Scanner    = &bufio.Scanner{}
		MySqlData  *structure.Secret = &structure.Secret{}
		secretData string            = ""
	)

	secretFile, err = os.Open("./json/Secret.json")
	if err != nil {
		SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open ./json/Secret.json")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
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
		SetLog("E", "config.ConnectToSqlDatabase() -> Unknow Error From bufio.Scanner()")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}

	Database, err = sql.Open("mysql", MySqlData.Username+":"+MySqlData.Password+"@tcp(0.0.0.0:3306)/WriteUp")
	if err != nil {
		SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open URLShortner Database")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
}

func DatabaseStatus() {
	err = Database.Ping()
	if err != nil {
		SetLog("E", "config.DatabaseStatus() -> MySQL Dosn't Response")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
	SetLog("I", "config.DatabaseStatus() -> MySQL Is Ready To Use")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	infoLog.SetFlags(log.Ldate | log.Ltime)
	warnLog.SetFlags(log.Ldate | log.Ltime)
	errorLog.SetFlags(log.Ldate | log.Ltime)
	SetLog("I", "config.init() -> Flags Setuped")
}
