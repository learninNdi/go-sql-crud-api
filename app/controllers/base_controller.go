package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	DB        *sql.DB
	Router    *mux.Router
	AppConfig *AppConfig
	Context   context.Context
}

type AppConfig struct {
	AppName string
	AppURL  string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBDriver string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.Context = context.Background()
	server.InitializeAppConfig(appConfig)
	server.InitializeDB(dbConfig)
	server.InitializeRoutes()
}

func (server *Server) InitializeAppConfig(appConfig AppConfig) {
	server.AppConfig = &appConfig
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)

	server.DB, err = sql.Open(dbConfig.DBDriver, dsn)

	if err != nil {
		panic("Failed to connect to database server")
	}
}
