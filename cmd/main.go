package main

import (
	"os"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/controllers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/database"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/services"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/providers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/repository"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/utils"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
	"github.com/labstack/gommon/log"
)

func main() {
	dbUsername := "root"
	dbPassword := "root"
	dbName := "app"
	dbPort := 5432
	dbHostname := "localhost"

	settings := &web.Settings{
		BindAddress:     "127.0.0.1",
		Port:            8000,
		FrontEndBaseUrl: "http://localhost:3000",
		JWTSecret:       "foobarken",
	}

	utils.LoadValidator()

	db, err := database.NewDatabaseConnection(dbUsername, dbPassword, dbName, dbHostname, dbPort)
	if err != nil {
		log.Errorf("Could not connect to database: %v", err.Error())
		os.Exit(1)
	}

	jwt := &providers.JWTProvider{
		Secret: settings.JWTSecret,
	}

	email := providers.NewEmailProvider()
	accountRepo := repository.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepo, email, settings)

	application := web.NewWebApplication(settings)
	application.RegisterController(controllers.NewAccountController("account-controller", accountService, settings))
	application.RegisterController(controllers.NewAuthController(
		"auth-controller",
		accountService,
		jwt,
	))
	application.RegisterController(controllers.NewJobPostController("jobpost-controller"))
	application.Start()
}
