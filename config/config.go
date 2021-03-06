package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"log"
	"os"
	"strings"
)

// ServerPort refers to server port.
var ServerPort string

// DbServer refers to database server ip.
var DbServer string

// DbPort refers to database server port.
var DbPort string

// DbUsername refers to database name.
var DbUsername string

// DbPassword refers to database password.
var DbPassword string

// DatabaseConnectionString refers to database connection string.
var DatabaseConnectionString string

// DatabaseName refers to database name.
var DatabaseName string

// Database refers to database options.
var Database string

// Publickey refers to publickey of EventStoreToken.
var Publickey string

// EnableAuthentication refers if service to service authentication is enabled.
var EnableAuthentication bool

// Token refers to oauth token for service to service communication.
var Token string

// EnableOpenTracing set true if opentracing is needed.
var EnableOpenTracing bool

// ServiceName service name of this application for opentracing
var ServiceName string

// RunMode refers to run mode.
var RunMode string

// Url for klovercloud ci-core
var CiCoreBaseUrl string

// InitEnvironmentVariables initializes environment variables
func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = string(enums.DEVELOP)
	}

	if RunMode != string(enums.PRODUCTION) {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)

	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database = os.Getenv("DATABASE")
	CiCoreBaseUrl = os.Getenv("KLOVERCLOUD_CI_CORE_URL")
	if Database == enums.MONGO {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}

	Publickey = os.Getenv("PUBLIC_KEY")

	if os.Getenv("ENABLE_AUTHENTICATION") == "" {
		EnableAuthentication = false
	} else {
		if strings.ToLower(os.Getenv("ENABLE_AUTHENTICATION")) == "true" {
			EnableAuthentication = true
		} else {
			EnableAuthentication = false
		}
	}

	if os.Getenv("ENABLE_OPENTRACING") == "" {
		EnableOpenTracing = false
	} else {
		if strings.ToLower(os.Getenv("ENABLE_OPENTRACING")) == "true" {
			EnableOpenTracing = true
		} else {
			EnableOpenTracing = false
		}
	}
	ServiceName = os.Getenv("JAEGER_SERVICE_NAME")
	Token = os.Getenv("TOKEN")
}
