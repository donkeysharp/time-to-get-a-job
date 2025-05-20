package main

import (
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/controllers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
)

func main() {
	settings := &web.Settings{
		BindAddress: "127.0.0.1",
		Port:        8000,
	}

	application := web.NewWebApplication(settings)
	application.RegisterController(controllers.NewAccountController("account-controller"))
	application.RegisterController(controllers.NewAuthController("auth-controller"))
	application.RegisterController(controllers.NewJobPostController("jobpost-controller"))
	application.Start()
}
