package goselenium

import (
	"errors"
	"testing"
)

func Test_CookieAllCookies_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.AllCookies()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CookieAllCookies_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.AllCookies()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CookieAllCookies_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.AllCookies()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CookieAllCookies_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": [
				{
					"name": "Test Cookie",
					"value": "Test Value",
					"path": "/",
					"domain": "www.google.com",
					"secure": true,
					"httpOnly": true,
					"expiry": "2016-12-25T00:00:00Z"
				}
			]
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.AllCookies()
	if err != nil || resp.State != "success" || resp.Cookies[0].Name != "Test Cookie" ||
		resp.Cookies[0].Value != "Test Value" || resp.Cookies[0].Path != "/" ||
		resp.Cookies[0].Domain != "www.google.com" || !resp.Cookies[0].SecureOnly ||
		!resp.Cookies[0].HTTPOnly {
		t.Errorf(correctResponseErrorText)
	}
}
