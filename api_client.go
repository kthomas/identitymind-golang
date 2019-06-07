package identitymind

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/vincent-petithory/dataurl"
)

const defaultContentType = "application/json"

// IdentityMindAPIClient is a generic base class for calling the identitymind API
type IdentityMindAPIClient struct {
	Host     string
	Path     string
	Scheme   string
	Token    *string
	Username *string
	Password *string
}

// NewIdentityMindAPIClient initializes an IdentityMindAPIClient using the environment-configured API
// user and token to construct an HTTP basic authorization header for access to the IdentityMind API.
func NewIdentityMindAPIClient() (*IdentityMindAPIClient, error) {
	apiURL, err := url.Parse(identitymindAPIBaseURL)
	if err != nil {
		log.Warningf("Failed to parse identitymind API base url; %s", err.Error())
		return nil, err
	}

	return &IdentityMindAPIClient{
		Host:     apiURL.Host,
		Scheme:   apiURL.Scheme,
		Path:     "",
		Username: stringOrNil(identitymindAPIUser),
		Password: stringOrNil(identitymindAPIToken),
	}, nil
}

func (i *IdentityMindAPIClient) sendRequest(method, urlString, contentType string, params map[string]interface{}, response interface{}) (status int, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Second * 30,
	}

	mthd := strings.ToUpper(method)
	reqURL, err := url.Parse(urlString)
	if err != nil {
		log.Warningf("Failed to parse URL for identitymind API (%s %s) invocation; %s", method, urlString, err.Error())
		return -1, err
	}

	if mthd == "GET" && params != nil {
		q := reqURL.Query()
		for name := range params {
			if val, valOk := params[name].(string); valOk {
				q.Set(name, val)
			}
		}
		reqURL.RawQuery = q.Encode()
	}

	headers := map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Accept":          {"application/json"},
	}
	if i.Username != nil && i.Password != nil {
		headers["Authorization"] = []string{buildBasicAuthorizationHeader(*i.Username, *i.Password)}
	} else if i.Token != nil {
		headers["Authorization"] = []string{fmt.Sprintf("Bearer %s", *i.Token)}
	}

	var req *http.Request

	if mthd == "POST" || mthd == "PUT" {
		var payload []byte
		if contentType == "application/json" {
			payload, err = json.Marshal(params)
			if err != nil {
				log.Warningf("Failed to marshal JSON payload for identitymind API (%s %s) invocation; %s", method, urlString, err.Error())
				return -1, err
			}
		} else if contentType == "application/x-www-form-urlencoded" {
			urlEncodedForm := url.Values{}
			for key, val := range params {
				if valStr, valOk := val.(string); valOk {
					urlEncodedForm.Add(key, valStr)
				} else {
					log.Warningf("Failed to marshal application/x-www-form-urlencoded parameter: %s; value was non-string", key)
				}
			}
			payload = []byte(urlEncodedForm.Encode())
		} else if contentType == "multipart/form-data" {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			for key, val := range params {
				if valStr, valStrOk := val.(string); valStrOk {
					dURL, err := dataurl.DecodeString(valStr)
					if err == nil {
						log.Debugf("Parsed data url parameter: %s", key)
						part, err := writer.CreateFormFile(key, key)
						if err != nil {
							return 0, err
						}
						part.Write(dURL.Data)
					} else {
						_ = writer.WriteField(key, valStr)
					}
				} else {
					log.Warningf("Skipping non-string value when constructing multipart/form-data request: %s", key)
				}
			}
			err = writer.Close()
			if err != nil {
				return 0, err
			}

			payload = []byte(body.Bytes())
		}

		req, _ = http.NewRequest(method, urlString, bytes.NewReader(payload))
		headers["Content-Type"] = []string{contentType}
	} else {
		req = &http.Request{
			URL:    reqURL,
			Method: mthd,
		}
	}

	req.Header = headers

	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Warningf("Failed to invoke identitymind API (%s %s) method: %s; %s", method, urlString, err.Error())
		return 0, err
	}

	log.Debugf("Received %v response for identitymind API (%s %s) invocation", resp.StatusCode, method, urlString)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("Failed to unmarshal identitymind API (%s %s) response: %s; %s", method, urlString, buf.Bytes(), err.Error())
	}

	log.Debugf("Invocation of identitymind API (%s %s) succeeded (%v-byte response)", method, urlString, buf.Len())
	return resp.StatusCode, nil
}

// Get constructs and synchronously sends an API GET request
func (i *IdentityMindAPIClient) Get(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("GET", url, defaultContentType, params, response)
}

// Post constructs and synchronously sends an API POST request
func (i *IdentityMindAPIClient) Post(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("POST", url, defaultContentType, params, response)
}

// PostWWWFormURLEncoded constructs and synchronously sends an API POST request using
func (i *IdentityMindAPIClient) PostWWWFormURLEncoded(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("POST", url, "application/x-www-form-urlencoded", params, response)
}

// PostMultipartFormData constructs and synchronously sends an API POST request using multipart/form-data as the content-type
func (i *IdentityMindAPIClient) PostMultipartFormData(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("POST", url, "multipart/form-data", params, response)
}

// Put constructs and synchronously sends an API PUT request
func (i *IdentityMindAPIClient) Put(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("PUT", url, defaultContentType, params, response)
}

// Delete constructs and synchronously sends an API DELETE request
func (i *IdentityMindAPIClient) Delete(uri string) (status int, err error) {
	url := i.buildURL(uri)
	return i.sendRequest("DELETE", url, defaultContentType, nil, nil)
}

func (i *IdentityMindAPIClient) buildURL(uri string) string {
	path := i.Path
	if len(path) == 1 && path == "/" {
		path = ""
	} else if len(path) > 1 && strings.Index(path, "/") != 0 {
		path = fmt.Sprintf("/%s", path)
	}
	return fmt.Sprintf("%s://%s%s/%s", i.Scheme, i.Host, path, uri)
}

func buildBasicAuthorizationHeader(username, password string) string {
	auth := fmt.Sprintf("%s:%s", username, password)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))
}
