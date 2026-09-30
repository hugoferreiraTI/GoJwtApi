package handler

import (
	"go_jwt/model"
	usecases "go_jwt/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHnadler struct {
	authUseCase *usecases.AuthuseCase
}

func NewAuthHandler(authUseCas *usecases.AuthuseCase) *AuthHnadler {
	return &AuthHnadler{
		authUseCase: authUseCas,
	}
}

func (h *AuthHnadler) Register(c *gin.Context) {
	var input model.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email incorreto"})
		return
	}

	userId, err := h.authUseCase.Register(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Cadastro concluido",
		"user_id": userId,
	})
}

func (h *AuthHnadler) Login(c *gin.Context) {
	var input model.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUseCase.Login(input)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Verifique o login e a senha",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *AuthHnadler) RefreshToken(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID, ok := userIDVal.(int)

	if !ok {
		if idFloat, okFloat := userIDVal.(float64); okFloat {
			userID = int(idFloat)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
			return
		}
	}

	token, err := h.authUseCase.RefreshToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token refresh failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_in": h.authUseCase.GetTokenExpiration().Seconds(),
		"token_type": "Bearer",
	})
}

func (h *AuthHnadler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":      "Successfully logged out",
		"instructions": "Please remo the token fron your cliente storage",
	})
}
