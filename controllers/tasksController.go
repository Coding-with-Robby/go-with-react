package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robbyklein/gr/initializers"
	"github.com/robbyklein/gr/models"
)

func FetchTasks(c *fiber.Ctx) error {
	// Get all tasks
	var tasks []models.Task
	initializers.DB.Order("created_at desc").Find(&tasks)

	return c.JSON(fiber.Map{
		"tasks": tasks,
	})
}

func CreateTask(c *fiber.Ctx) error {
	// Parse body
	var body struct {
		Title string
		Body  string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Create record
	task := models.Task{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		return result.Error
	}

	// return it
	return c.JSON(fiber.Map{
		"task": task,
	})
}

func FetchTask(c *fiber.Ctx) error {
	// Fetch task
	var task models.Task
	taskId := c.Params("id")

	result := initializers.DB.First(&task, taskId)

	if result.Error != nil {
		return result.Error
	}

	// return it
	return c.JSON(fiber.Map{
		"task": task,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	result := initializers.DB.Delete(&models.Task{}, taskId)

	if result.Error != nil {
		return result.Error
	}

	// return it
	return c.JSON(fiber.Map{
		"success": "Task deleted",
	})
}
