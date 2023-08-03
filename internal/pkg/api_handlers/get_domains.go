package api_handlers

import (
	"log"
	"net/http"

	"github.com/ivankuchin/domain-name/internal/pkg/utils"
)

var domain string = ""

func GetDomain(w http.ResponseWriter, r *http.Request) {
	traceID := utils.GenerateTraceID()
	log.Printf("traceID: %s, HTTP request: %s", traceID, r.RequestURI)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(domain))
}

func SetDomain(w http.ResponseWriter, r *http.Request) {
	traceID := utils.GenerateTraceID()

	// get value from environment variable DOMAIN_NAME
	domain = r.Host

	log.Printf("traceID: %s, HTTP request: %s, set domain name to %s", traceID, r.RequestURI, domain)

	w.WriteHeader(http.StatusOK)
}
