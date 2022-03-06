package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/configs"
	log "github.com/Toflex/oris_log"
	"io"
	"net/http"
)

type Client struct {
	UrlName string
	Params  string
	Body interface{}
	UserName string // For basic auth
	Password string // For basic auth
	Header map[string]string // set request header
	Config *configs.Config
	Logger log.Logger
	ServiceName string // Name of external server
}

// ClientRequest returns response from an external server
func (c *Client) ClientRequest() (interface{}, error) {

	ctx := make(map[string]interface{})
	l:=c.Logger.NewContext(ctx)
	defer l.Info("Service Log")

	client := &http.Client{}

	body, err := json.Marshal(c.Body)
	if err!=nil{
		c.Logger.Error("Failed to marshal request body. %s", err)
		return nil, err
	}

	Endpoint := GetEndpoint(c.UrlName, c.Config)
	url:=fmt.Sprintf("%s%s%s", Endpoint.BaseURL,Endpoint.Path,c.Params)
	c.Logger.Info("URL: %s", url)

	req, err := http.NewRequest(Endpoint.Method, url, bytes.NewBuffer(body))
	l.AddContext("Service", c.ServiceName)
	l.AddContext("Request", string(body))
	l.AddContext("URL", req.RequestURI)
	l.AddContext("Method", req.Method)
	if err!=nil{
		c.Logger.Error("Error occurred %s: %s", url, err)
		return nil, err
	}

	// Set request headers
	req.Header.Add("Content-Type", "application/json")
	for k, v := range c.Header {
		req.Header.Add(k, v)
	}

	// Set basic authentication
	if c.UserName!="" {
		req.SetBasicAuth(c.UserName,c.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Logger.Error("Error occurred %s: %s", url, err)
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error("Error occurred %s: %s", url, err)
		return nil, err
	}

	var result interface{}
	json.Unmarshal(res, &result)
	c.Logger.Info("Response body %+v", result)

	l.AddContext("Response", result)
	l.AddContext("Code", resp.StatusCode)
	return result, nil
}
