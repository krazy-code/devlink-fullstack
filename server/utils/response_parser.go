package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
	Errors any    `json:"errors"`
}

func ResponseParser(c *fiber.Ctx, r Response) error {
	return c.Status(r.Code).JSON(
		Response{
			Code:   r.Code,
			Status: GetStatusMessage(r.Code),
			Data:   r.Data,
			Errors: r.Errors,
		},
	)
}

func GetStatusMessage(c int) string {
	statusCodeString := map[int]string{
		// 2xx
		200: "OK",
		201: "Created",
		204: "No Content Found",
		// 4xx
		400: "Bad Request",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		// 5xx
		500: "Internal Server Error",
	}
	status := statusCodeString[c]
	return status

}
