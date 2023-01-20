package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
)

type Schedule struct {
	Task string    `json:"task"`
	Time time.Time `json:"time"`
}

func startScheduler(schedule Schedule) {

	duration := schedule.Time.Sub(time.Now())

	time.Sleep(duration)

	fmt.Printf("Reminder: %s at %s\n", schedule.Task, schedule.Time)
}

func main() {
	app := fiber.New()

	app.Post("/schedule", func(c *fiber.Ctx) {
		var lunchSchedule Schedule
		lunchSchedule.Task = "Eat lunch"
		lunchSchedule.Time = time.Now().Add(time.Minute)

		go startScheduler(lunchSchedule)

		var waterSchedule Schedule
		waterSchedule.Task = "Drink water"
		waterSchedule.Time = time.Now().Add(2 * time.Minute)

		go startScheduler(waterSchedule)

		c.Send("Task scheduled for lunch in 1 minute and task scheduled for drinking water in 2 minutes")
	})

	app.Listen(3000)
}
