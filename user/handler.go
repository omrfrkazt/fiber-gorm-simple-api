package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(*fiber.Ctx) error
	Create(*fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
}

func (h handler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(Response{
			Error:      "Bad Request",
			StatusCode: 400,
			Success:    false,
		})
	}
	model, err := h.service.Get(uint(id))
	if err != nil {
		return c.Status(500).JSON(Response{
			Error:      "Internal Server Error",
			StatusCode: 500,
			Success:    false,
		})
	}
	return c.Status(200).JSON(Response{
		Data:       model,
		StatusCode: 200,
		Success:    true,
	})
}

func (h handler) Create(c *fiber.Ctx) error {
	model := Model{}
	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(Response{
			Error:      "Bad Request",
			StatusCode: 400,
			Success:    false,
		})
	}
	_, err = h.service.Create(model)
	if err != nil {
		return c.Status(500).JSON(Response{
			Error:      "Internal Server Error",
			StatusCode: 500,
			Success:    false,
		})
	}
	return c.Status(201).JSON(Response{
		StatusCode: 201,
		Success:    true,
	})
}
