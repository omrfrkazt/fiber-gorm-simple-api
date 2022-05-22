package user

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/omrfrkazt/fiber-gorm-simple-api/db"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	err = repo.Migration()
	assert.Nil(t, err)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/user/:id", handler.Get)
	id, err := repo.Create(Model{Name: "test", Email: "test@test.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest(fiber.MethodGet, fmt.Sprintf("/user/%d", id), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}
