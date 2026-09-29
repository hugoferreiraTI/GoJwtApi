package gojwt

import (
	"fmt"
	"go_jwt/config/db"
	jwtconfig "go_jwt/config/jwt_config"
	"go_jwt/model"
	"go_jwt/repository"
	usecases "go_jwt/useCases"
	"testing"

	_ "github.com/lib/pq"
)

func TestJwt(t *testing.T) {
	response, _ := jwtconfig.Load()

	fmt.Println(response)
}

func TestEmail(t *testing.T) {
	email := model.UserRegister{
		Email: "testandoemail@gmail.com",
	}

	err := email.Validate()
	if err != nil {
		fmt.Println(err)
	}
}

func TestDbConnection(t *testing.T) {
	sql, err := db.ConnectDB()

	if err != nil {
		fmt.Println("Conexão falhou")
	}

	defer sql.Close()

	var versionPostreg string

	err = sql.QueryRow("SELECT version();").Scan(&versionPostreg)

	if err != nil {
		fmt.Println("Conexão falhou")
	}

	fmt.Println("Conexão bem sucedida")

}

func TestCheckEmail(t *testing.T) {
	sql, err := db.ConnectDB()
	initializeUse := repository.NewUserRepository(sql)

	if err != nil {
		t.Fatalf("Conexão com o banco de dados falhou: %v", err)
	}
	fmt.Print("Conexão bem sucedida, tentando iniciar comando sql")

	user := model.UserLogin{
		Email: "testando@gmail.com",
	}
	defer sql.Close()

	test, _ := initializeUse.EmailExists(user.Email)

	fmt.Print(test)
}

func TestCreatUser(t *testing.T) {
	user := model.UserRegister{
		Email:    "hugoferreiraferro@gmail.com",
		Password: "123456",
	}
	//abrindo conexão com banco de dados
	sql, err := db.ConnectDB()

	if err != nil {
		t.Fatalf("Conexão com o banco de dados falhou: %v", err)
	}

	fmt.Print("Conexão bem sucedida, tentando iniciar comando sql")
	//fecho conexão com banco de dados
	defer sql.Close()

	userRepo := repository.NewUserRepository(sql)

	initializeUserRegister := usecases.NewAuthUseCase(userRepo)

	value, errRegister := initializeUserRegister.Register(user.Email, user.Password)

	if errRegister != nil {
		fmt.Print("Teve um erro na criação do usuário")
	}

	fmt.Print(value)
	fmt.Print(errRegister)
}
