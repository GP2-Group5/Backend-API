package factory

import (
	authDelivery "github.com/GP2-Group5/Backend/feature/auth/delivery"
	authRepo "github.com/GP2-Group5/Backend/feature/auth/repository"
	authService "github.com/GP2-Group5/Backend/feature/auth/service"

	userDelivery "github.com/GP2-Group5/Backend/feature/users/delivery"
	userRepo "github.com/GP2-Group5/Backend/feature/users/repository"
	userService "github.com/GP2-Group5/Backend/feature/users/service"

	menteeDelivery "github.com/GP2-Group5/Backend/feature/mentee/delivery"
	menteeRepo "github.com/GP2-Group5/Backend/feature/mentee/repository"
	menteeService "github.com/GP2-Group5/Backend/feature/mentee/service"

	classDelivery "github.com/GP2-Group5/Backend/feature/classes/delivery"
	classRepo "github.com/GP2-Group5/Backend/feature/classes/repository"
	classService "github.com/GP2-Group5/Backend/feature/classes/service"

	logDelivery "github.com/GP2-Group5/Backend/feature/log/delivery"
	logRepo "github.com/GP2-Group5/Backend/feature/log/repository"
	logService "github.com/GP2-Group5/Backend/feature/log/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	menteeRepoFactory := menteeRepo.New(db)
	menteeServiceFactory := menteeService.New(menteeRepoFactory)
	menteeDelivery.New(menteeServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	classRepoFactory := classRepo.New(db)
	classServiceFactory := classService.New(classRepoFactory)
	classDelivery.New(classServiceFactory, e)

	logRepoFactory := logRepo.New(db)
	logServiceFactory := logService.New(logRepoFactory)
	logDelivery.New(logServiceFactory, e)
}
