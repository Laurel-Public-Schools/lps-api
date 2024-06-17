package handler

import (
	"encoding/json"
	// "errors"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/laurel-public-schools/lps-api/env"
	"github.com/laurel-public-schools/lps-api/utils"
)

type EmailRequest struct {
	Key     string
	To      string
	From    string
	Subject string
	HTML    string
}

// GetEmail godoc
// @Summary Send an email
// @Description Send an email via post request
// @ID send-email
// @Tags email
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /email [post]
func (h *Handler) SendEmail(c echo.Context) error {
	var e EmailRequest
	var (
		config = env.GetConfig()
	)

	auth := smtp.PlainAuth("", config.User, config.Pass, "smtp.gmail.com")

	bod := c.Request().Body

	err := json.NewDecoder(bod).Decode(&e)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	// if e.Key != config.APIKey {
	// 	return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("invalid api key")))
	// }

	from := mail.Address{Name: e.From, Address: config.User}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	subject := "Subject: " + e.Subject + "\n"

	header := "From: " + from.Name + "<" + from.Address + ">\n"

	msg := []byte(header + subject + mime + "\n" + e.HTML)

	recipients := strings.Split(e.To, ", ")

	for i, recipient := range recipients {
		recipients[i] = strings.TrimSpace(recipient)
	}

	errs := smtp.SendMail("smtp.gmail.com:587", auth, config.User, recipients, msg)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(errs))
	}

	return c.JSON(http.StatusOK, "Email sent")
}
