package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/laurel-public-schools/lps-api/env"
	"github.com/laurel-public-schools/lps-api/middleware"
)

type EmailRequest struct {
	Key     string
	To      string
	From    string
	Subject string
	HTML    string
}

type EmailRoute struct{}

// SendEmail godoc
//
//	@Summary		Send an email
//	@Description	Send an email via post request
//	@ID				send-email
//	@Tags			email
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/email [post]
func (rs EmailRequest) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthCtx)

	r.Post("/", rs.SendEmail)
	return r
}

func (rs EmailRequest) SendEmail(w http.ResponseWriter, r *http.Request) {
	var e EmailRequest
	var (
		config = env.GetConfig()
	)

	auth := smtp.PlainAuth("", config.User, config.Pass, "smtp.gmail.com")

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	from := mail.Address{Name: e.From, Address: config.User}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	subject := "Subject: " + e.Subject + "!\n"
	header := "From: " + from.Name + "<" + from.Address + ">\n"
	msg := []byte(header + subject + mime + "\n" + e.HTML)
	recipients := strings.Split(e.To, ", ")
	for i, recipient := range recipients {
		recipients[i] = strings.TrimSpace(recipient)
	}

	errs := smtp.SendMail("smtp.gmail.com:587", auth, config.User, recipients, msg)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		fmt.Println(errs)
		return
	}
}
