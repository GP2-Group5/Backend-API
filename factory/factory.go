package factory

import (
	authDelivery "github.com/GP2-Group5/Backend/feature/auth/delivery"
	authRepo "github.com/GP2-Group5/Backend/feature/auth/repository"
	authService "github.com/GP2-Group5/Backend/feature/auth/service"

	userDelivery "github.com/GP2-Group5/Backend/feature/users/delivery"
	userRepo "github.com/GP2-Group5/Backend/feature/users/repository"
	userService "github.com/GP2-Group5/Backend/feature/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

}
