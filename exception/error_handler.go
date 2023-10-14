package exception

import (
	"example-go-restapi/model/api"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	errResponse := notFoundError(ctx, err)
	if errResponse != nil {
		mobileResponse := api.APIResponse{
			Code:  http.StatusNotFound,
			Data:  nil,
			Error: errResponse.Error(),
		}
		return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	}
	errResponse = badRequestError(ctx, err)
	if errResponse != nil {
		mobileResponse := api.APIResponse{
			Code:  http.StatusBadRequest,
			Data:  nil,
			Error: errResponse.Error(),
		}
		return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	}

	errResponse = unauthorizedError(ctx, err)
	if errResponse != nil {
		mobileResponse := api.APIResponse{
			Code:  http.StatusUnauthorized,
			Data:  nil,
			Error: errResponse.Error,
		}
		return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	}
	errResponse = validationErrors(ctx, err)
	if errResponse != nil {
		mobileResponse := api.APIResponse{
			Code:  http.StatusBadRequest,
			Data:  nil,
			Error: errResponse.Error,
		}
		return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	}

	mobileResponse := api.APIResponse{
		Code:  http.StatusInternalServerError,
		Data:  nil,
		Error: errResponse.Error,
	}
	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
}

func validationErrors(ctx *fiber.Ctx, err error) error {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		return exception
	} else {
		return nil
	}
}

func notFoundError(ctx *fiber.Ctx, err error) error {
	exception, ok := err.(NotFoundError)
	if ok {
		return exception
	} else {
		return nil
	}
}
func badRequestError(ctx *fiber.Ctx, err error) error {
	exception, ok := err.(BadRequestError)
	if ok {
		return exception
	} else {
		return nil
	}
}

func unauthorizedError(ctx *fiber.Ctx, err error) error {
	exception, ok := err.(UnauthorizedError)
	if ok {
		return exception
	} else {
		return nil
	}
}
