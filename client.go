// Package cloudsigma provides an http rest client for CloudSigma's cloud api.
package cloudsigma

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Client is the CloudSigma http client.
type Client struct{}

// Args holds the arguments for the http request.
type Args struct {
	Resource     string
	Headers      []Header
	Verb         string
	ObjectId     string
	GetReqParams map[string]string
	ActionName   string
	Body         interface{}
	RequiresAuth bool
	Username     string
	Password     string
	Location     string
	Format       string
	OutFile      string
}

// Header is a name value pair http header.
type Header struct {
	Name  string
	Value string
}

// CloudSigmaRequest is a local *http.Request, defined because new methods
// cannot be defined for types non-local to the current package.  *http.Request
// is embedded in a new type, rather than type-aliased because that results in
// an incomplete replica.  Using an anonymous field avoids field references.
type CloudSigmaRequest struct {
	*http.Request
}

const (
	BaseUrl         = "https://%s/api/%s/"
	ApiEndpoint     = "%s.cloudsigma.com"
	ApiVersion      = "2.0"
	UrlResourceList = "%s/"
	UrlResource     = "%s/%s"
	//UrlAction       = "action/?do=%s"
	UrlAction = "action/"
)

// NewArgs returns an Args object.
func NewArgs() *Args {
	a := Args{}
	a.GetReqParams = make(map[string]string)
	return &a
}

// NewClient returns a Client object.
func NewClient() *Client {
	c := Client{}
	return &c
}

// buildBaseUrl returns the url on which all requests are built.
func (c *Client) buildBaseUrl(location string) (string, error) {
	if location == "" {
		return "", errors.New("Service location not set.  Use command: set config location")
	}
	apiEndpoint := fmt.Sprintf(ApiEndpoint, location)
	return fmt.Sprintf(BaseUrl, apiEndpoint, ApiVersion), nil
}

// buildResourceUrl returns the url for a specified resource.
func (c *Client) buildResourceUrl(location string, resource string) (*url.URL, error) {
	u, err := c.buildBaseUrl(location)
	if err != nil {
		return &url.URL{}, err
	}
	u += resource
	if resource != "" {
		u += "/"
	}
	newUrl, err := url.Parse(u)
	if err != nil {
		return &url.URL{}, err
	}
	return newUrl, nil
}

func (a *Args) AddHeaderPair(name string, value string) {
	a.Headers = append(a.Headers, Header{name, value})
}

// AddHeader adds a header to the Headers field of an Args object.
func (a *Args) AddHeader(header Header) {
	a.Headers = append(a.Headers, header)
}

// AddHeaders add Args.Headers to a CloudSigmaRequest object.
func (r *CloudSigmaRequest) AddHeaders(headers []Header) {
	for _, j := range headers {
		r.Header.Add(j.Name, j.Value)
	}
}

// Call builds and sends a request, given Args, and returns the http result.
func (c *Client) Call(client *http.Client, args *Args) ([]byte, error) {
	req, err := c.buildRequest(args)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendRequest(client, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Download builds and sends a request, given Args, and saves the body to file.
func (c *Client) Download(client *http.Client, args *Args) ([]byte, error) {
	req, err := c.buildRequest(args)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendDownloadRequest(client, req, args.OutFile)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// buildRequest creates a http CloudSigmaRequest with the supplied Args.
func (c *Client) buildRequest(args *Args) (*CloudSigmaRequest, error) {
	u, err := c.buildResourceUrl(args.Location, args.Resource)
	if err != nil {
		return nil, err
	}

	// Add object id to path, if provided.
	if args.ObjectId != "" {
		u, err = url.Parse(u.String() + args.ObjectId + "/")
		if err != nil {
			return nil, err
		}
	}

	// Response ContentType: Json or xml; json is default.
	// Only for GETs, add format querystring paramater.
	if strings.ToUpper(args.Verb) == "GET" {
		format := "json"
		if args.Format == "xml" {
			format = "xml"
		}
		args.GetReqParams["format"] = format
	}

	// Add the action path and do querystring parameter if provided.
	params := url.Values{}
	if args.ActionName != "" {
		u, err = url.Parse(u.String() + UrlAction)
		if err != nil {
			return nil, err
		}
		params.Add("do", args.ActionName)
	}

	// Add querystring params.
	for k, v := range args.GetReqParams {
		params.Add(k, v)
	}
	u.RawQuery = params.Encode()

	// Build the json body if required.
	bodybuf := bytes.NewBuffer([]byte{})
	if args.Body != nil {
		jsonBody, err := json.Marshal(args.Body)
		if err != nil {
			fmt.Println("Error marshalling body to json.")
			return nil, err
		}
		bodybuf = bytes.NewBuffer([]byte(jsonBody))
	}

	// Build the request.
	newreq, err := http.NewRequest(strings.ToUpper(args.Verb), u.String(), bodybuf)
	req := CloudSigmaRequest{Request: newreq}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Add auth if required.
	if args.RequiresAuth {
		args.AddHeader(c.GetHttpBasicAuthHeader(args.Username, args.Password))
	}

	// Add headers to request.
	req.AddHeaders(args.Headers)
	return &req, nil
}

// sendRequest sends the given http CloudSigmaRequest and returns the result.
func (c *Client) sendRequest(client *http.Client, req *CloudSigmaRequest) ([]byte, error) {
	if client == nil {
		client = &http.Client{}
	}
	req.Header.Add("Accept-Encoding", "identity")
	fmt.Printf("Req: %+v\n", req.Request)
	resp, err := client.Do(req.Request)
	if err != nil {
		log.Println("Error in client.Do.")
		log.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	if !strings.HasPrefix(strconv.Itoa(resp.StatusCode), "2") {
		return []byte{}, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in ioutil.ReadAll.")
		return []byte{}, err
	}
	return body, nil
}

// sendRequest sends the given http CloudSigmaRequest and returns the result.
func (c *Client) sendDownloadRequest(client *http.Client, req *CloudSigmaRequest, filename string) ([]byte, error) {
	f, err := os.Create(filename)
	defer f.Close()

	if client == nil {
		client = &http.Client{}
	}
	req.Header.Add("Accept-Encoding", "identity")
	fmt.Printf("Req: %+v\n", req.Request)
	resp, err := client.Do(req.Request)
	if err != nil {
		log.Println("Error in client.Do.")
		log.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	if !strings.HasPrefix(strconv.Itoa(resp.StatusCode), "2") {
		return []byte{}, errors.New(resp.Status)
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println("Error in io.Copy.")
		return []byte{}, err
	}
	return []byte("OK"), nil
}

// GetHttpBasicAuthHeader returns a base 64 encoded auth header,
// given a username and password.
func (c *Client) GetHttpBasicAuthHeader(userEmail string, password string) Header {
	// Authorization: Basic base64_encode(useremail:password)
	header := "Authorization"
	valueFormat := "Basic %s"
	credsFormat := "%s:%s"
	credsData := []byte(fmt.Sprintf(credsFormat, userEmail, password))
	credsBase64 := base64.StdEncoding.EncodeToString(credsData)
	value := fmt.Sprintf(valueFormat, credsBase64)
	return Header{header, value}
}
