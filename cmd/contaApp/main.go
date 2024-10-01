package main

import (
	"fmt"
	"log"

	"github.com/Lucas-Sampaio/ContaBancaria/configs"
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	webserver "github.com/Lucas-Sampaio/ContaBancaria/internal/Api"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Api/controllers"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// connectionString := fmt.Sprintf("sqlserver://sa:SqlServer2022!@localhost:1433?database=goexpert")
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", configs.DBDriver, configs.DBUser,
		configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	fmt.Print(connectionString)
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	db.AutoMigrate(&domain.Conta{})

	unitOfWork := database.NewUnitOfWork(db)
	contaController := controllers.NewContaController(unitOfWork)

	server := webserver.NewWebServer(configs.WebServerPort, *contaController)
	err = server.Start()
	if err != nil {
		log.Fatal("Erro a instanciar servidor", err)
	}
}