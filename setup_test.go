package m3_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/omniboost/go-m3-accounting"
)

var (
	client            *m3.Client
	customerShortName string
	propertyCode      string
	sourceSystem      string
)

func TestMain(m *testing.M) {
	var err error

	customerShortName = os.Getenv("CUSTOMER_SHORT_NAME")
	propertyCode = os.Getenv("PROPERTY_CODE")
	sourceSystem = os.Getenv("SOURCE_SYSTEM")

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	debug := os.Getenv("DEBUG")

	client = m3.NewClient(nil, clientID, clientSecret)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	client.SetDisallowUnknownFields(true)
	m.Run()
}
