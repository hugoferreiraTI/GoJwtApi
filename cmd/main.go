package main

import (
	"go_jwt/config/db"
	jwtconfig "go_jwt/config/jwt_config"
	"go_jwt/handler"
	"go_jwt/repository"
	usecases "go_jwt/useCases"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	//iniciando o database
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	//camada de repository
	authRepository := repository.NewUserRepository(dbConnection)

	cfg, _ := jwtconfig.Load() //carregando o env e o expiry
	HandlerCase := usecases.NewAuthUseCase(authRepository, []byte(cfg.Secret), cfg.TokenExpiry)

	//camada handler
	authHandler := handler.NewAuthHandler(HandlerCase)

	//Configurando GIN
	server := gin.Default()

	//rotas publicas
	public := server.Group("/api/v1")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	protected := server.Group("/api/v1")
	protected.Use(jwtconfig.AutMiddleware([]byte(cfg.Secret)))
	{
		protected.POST("/refresh-token", authHandler.RefreshToken)
		protected.POST("/logout", authHandler.Logout)
	}

	// Inicia o servidor na porta 8000
	server.Run(":8000")
}
