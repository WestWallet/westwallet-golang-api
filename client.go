package westwallet

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	// ENDPOINT defaults to https://api.westwallet.info
	// but can be overridden for test purposes
	ENDPOINT = "https://api.westwallet.info"
)

// APIClient is the interface for most of the API calls
// If Endpoint aren't defined the library
// will use the default https://api.westwallet.info
type APIClient struct {
	Key      string
	Secret   string
	Endpoint string
}

// Fetch works as a wrapper for all kind of http requests. It requires a http method
// and a relative path to the API endpoint. It will try to decode all results into
// a single interface type which you can provide.
func (a *APIClient) Fetch(method, path string, body interface{}, result interface{}) error {
	if a.Endpoint == "" {
		// use default endpoint
		a.Endpoint = ENDPOINT
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	var bodyBuffered io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBuffered = bytes.NewBuffer([]byte(data))
	}
	req, err := http.NewRequest(method, a.Endpoint+path, bodyBuffered)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	err = a.Authenticate(path, req, body)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	err = checkHTTPErrors(resp)
	if err != nil {
		return err
	}
	return nil
}

func checkHTTPErrors(resp *http.Response) error {
	if resp.StatusCode == 401 {

	}
	return nil
}

// Authenticate works with the Fetch call and adds certain Headers
// to the http request. This includes the actual API key and the
// timestamp of the request. Also a signature which is encoded
// with hmac and the API secret key.
func (a *APIClient) Authenticate(path string, req *http.Request, body interface{}) error {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	message := timestamp
	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		message += string(bodyBytes)
	}
	sha := sha256.New
	h := hmac.New(sha, []byte(a.Secret))
	h.Write([]byte(message))

	signature := fmt.Sprintf("%x", h.Sum(nil))

	req.Header.Set("X-API-KEY", a.Key)
	req.Header.Set("X-ACCESS-SIGN", signature)
	req.Header.Set("X-ACCESS-TIMESTAMP", timestamp)

	return nil
}
