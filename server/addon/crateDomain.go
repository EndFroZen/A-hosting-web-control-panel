package addon

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func CreateDomain(c *fiber.Ctx) error {
	domain := c.FormValue("domain")
	langue := c.FormValue("langue")
	// express := c.FormValue("express")
	cmd := exec.Command("cmd", "/C", " if exist cd ../create/user1", domain, "echo exist alrady")
	ouput, err := cmd.CombinedOutput()
	s := string(ouput)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	if s == "exist alrady" {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	cmd = exec.Command("cmd", "/C", "cd ../create/user1 && mkdir ", domain, "&& echo suc")

	

	return c.Redirect("http://localhost:3000/dashboard")
}
