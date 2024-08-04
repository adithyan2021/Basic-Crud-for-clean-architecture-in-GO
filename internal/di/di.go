package di

import (
	"CRUD-CLEAN/internal/config"
	"github.com/shakezidin/internal/server"
)

func Init(conf config.Config) *server.ServerStruct {
	server := server.NewHTTPServer()
	db := db.ConnectPGDB(conf)
	repo := repository.NewUserRepository(db)
	usecase := usecase.UseCase(repo)
	delivery := delivery.UserDelivery(usecase)
	userRoutes := routes.NewUserInit(server, delivery)
	userRoutes.UserRoutes()
	return server
}

// func InitiallizeEvent(conf config.Config) (*bootserver.ServerHttp, error) {

// 	DB := db.ConnectPGDB(conf)

// 	userRepository := userauth.NewRepository(DB)
// 	userService := userauth.NewService(userRepository)
// 	userHandler := userauth.NewHandler(userService)

// 	serverHttp := bootserver.NewServerHttp(*userHandler)

// 	return serverHttp, nil

// }
