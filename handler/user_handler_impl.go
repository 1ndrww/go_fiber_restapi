package handler

import (
	"example-go-restapi/exception"
	"example-go-restapi/model/api"
	"example-go-restapi/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHandlerImpl struct {
	UserService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) UserHandlerInterface {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

func (h UserHandlerImpl) FindAll(c *fiber.Ctx) error {
	users := h.UserService.FindAll()
	mobileResponse := api.APIResponse{
		Code:  200,
		Data:  users,
		Error: nil,
	}
	return c.Status(mobileResponse.Code).JSON(mobileResponse)
}

func (h UserHandlerImpl) FindByID(c *fiber.Ctx) error {
	userId := c.AllParams()["id"]
	user := h.UserService.FindById(userId)
	mobileResponse := api.APIResponse{
		Code:  200,
		Data:  user,
		Error: nil,
	}
	return c.Status(mobileResponse.Code).JSON(mobileResponse)

}

func (h UserHandlerImpl) Save(c *fiber.Ctx) error {
	userRequest := new(api.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	user := h.UserService.Save(*userRequest)
	mobileResponse := api.APIResponse{
		Code:  200,
		Data:  user,
		Error: nil,
	}
	return c.Status(mobileResponse.Code).JSON(mobileResponse)
}

func (h UserHandlerImpl) Edit(c *fiber.Ctx) error {
	userId := c.AllParams()["id"]
	userRequest := new(api.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		log.Warn(err)
		panic(exception.NewBadRequestError(err.Error()))
	}
	log.Info(*userRequest)

	user := h.UserService.Edit(*userRequest, userId)
	mobileResponse := api.APIResponse{
		Code:  200,
		Data:  user,
		Error: nil,
	}
	return c.Status(mobileResponse.Code).JSON(mobileResponse)
}

func (h UserHandlerImpl) Delete(c *fiber.Ctx) error {
	userId := c.AllParams()["id"]
	h.UserService.Delete(userId)
	mobileResponse := api.APIResponse{
		Code:  200,
		Data:  nil,
		Error: nil,
	}
	return c.Status(mobileResponse.Code).JSON(mobileResponse)
}
