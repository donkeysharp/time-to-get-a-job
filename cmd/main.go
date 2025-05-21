package main

import (
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/controllers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/services"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/providers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/repository"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
)

func main() {
	settings := &web.Settings{
		BindAddress: "127.0.0.1",
		Port:        8000,
	}

	jwt := &providers.JWTProvider{}
	accountRepo := repository.NewAccountRepository()
	accountService := services.NewAccountService(accountRepo)

	application := web.NewWebApplication(settings)
	application.RegisterController(controllers.NewAccountController("account-controller"))
	application.RegisterController(controllers.NewAuthController(
		"auth-controller",
		accountService,
		jwt,
	))
	application.RegisterController(controllers.NewJobPostController("jobpost-controller"))
	application.Start()
}
