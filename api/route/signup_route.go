package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-backend-clean-architecture/api/controller"
	"go-backend-clean-architecture/bootstrap"
	"go-backend-clean-architecture/domain"
	"go-backend-clean-architecture/mongo"
	"go-backend-clean-architecture/repository"
	"go-backend-clean-architecture/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
