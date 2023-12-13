package route

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/userManager/mongo"
	"IoT-backend/server/userManager/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func dummy(any interface{}) {}

func Setup(env *configManager.Env, timeout time.Duration, db mongo.Database,
	gin *gin.Engine) {
	publicRouter := gin.Group("")
	dummy(publicRouter)

}

func NewSignupRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
