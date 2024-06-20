package routes

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/laurel-public-schools/lps-api/internal/env"
	"github.com/laurel-public-schools/lps-api/internal/utils"
)

type ICRoute struct{}

func (rs ICRoute) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", rs.getClasses)
	return r
}

type LinkObject struct {
	Href      string `json:"href"`
	SourcedId string `json:"sourcedId"`
	Type      string `json:"type"`
}

type ClassInfo struct {
	SourcedId        string       `json:"sourcedId"`
	Status           string       `json:"status"`
	DateLastModified string       `json:"dateLastModified"`
	Title            string       `json:"title"`
	ClassType        string       `json:"classType"`
	ClassCode        string       `json:"classCode"`
	Location         *string      `json:"location,omitempty"`
	Course           LinkObject   `json:"course"`
	School           LinkObject   `json:"school"`
	Terms            []LinkObject `json:"terms"`
	Periods          []string     `json:"periods"`
}

type ClassResponse struct {
	Classes []ClassInfo `json:"classes"`
}

// @Summary	Get Classes
func (rs ICRoute) getClasses(w http.ResponseWriter, r *http.Request) {
	envConfig := env.GetEnv()
	if envConfig == nil {
		http.Error(w, "Invalid environment configuration", http.StatusInternalServerError)
		return
	}

	token, err := utils.IcAuth()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	baseUrl := utils.IcClassQuery(envConfig.LHS_SOURCE_ID, "laurel")

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-XSRF-TOKEN", envConfig.XSRF_TOKEN)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
