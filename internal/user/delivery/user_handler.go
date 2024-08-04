package delivery

import (
	"github.com/gin-gonic/gin"
	user "github.com/shakezidin/internal/user/entity"
	"github.com/shakezidin/internal/user/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type UserUserCase interface {
	RegisterUserHandler(c *gin.Context)
	LoginUserHandler(c *gin.Context)
}

func (u *UserHandler) RegisterUserHandler(c *gin.Context) {
	var user user.UserRegister
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error": "binding error"})
		return
	}

	err := u.userUseCase.RegisterUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"Error": "user already exist"})
		return
	}

	c.JSON(200, gin.H{"Status": "User registration Success"})
}

func (u *UserHandler) LoginUserHandler(c *gin.Context) {
	var userLogin user.UserLogin
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"Error": "binding error"})
		return
	}

	user, err := u.userUseCase.Login(&userLogin)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "Success", "user": user})
}

func UserDelivery(userUseCase usecase.UserUseCase) UserUserCase {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
