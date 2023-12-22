package route

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/userManager/api/controller"
	"IoT-backend/server/userManager/api/middleware"
	"IoT-backend/server/userManager/domain"
	"IoT-backend/server/userManager/mongo"
	"IoT-backend/server/userManager/repository"
	"IoT-backend/server/userManager/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *configManager.Env, timeout time.Duration, db mongo.Database,
	gin *gin.Engine) {
	// Public APIs
	publicRouter := gin.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewProfileRouter(env, timeout, db, protectedRouter)
	NewUpdateSensorRouter(env, timeout, db, protectedRouter)
	NewRequestSensorRouter(env, timeout, db, protectedRouter)
}

func NewSignupRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}

func NewLoginRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

func NewRefreshTokenRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}

func NewProfileRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}

/* Request for add or remove sensors */
func NewUpdateSensorRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	usc := &controller.UpdateSensorController{
		UpdateSensorUsecase: usecase.NewUpdateSensorUsecase(ur, timeout),
	}
	group.POST("/updateSensor", usc.UpdateSensor)
}

/* Request for acquire current sensor parameter */
func NewRequestSensorRouter(env *configManager.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rsc := &controller.RequestSensorController{
		RequestSensorUsecase: usecase.NewRequestSensorUsecase(ur, timeout),
	}
	group.GET("/requestSensor", rsc.RequestSensor)
}
