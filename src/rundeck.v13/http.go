package rundeck

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"gopkg.in/jmcvetta/napping.v2"
)

func (rc *RundeckClient) Get(i *[]byte, path string, options interface{}) error {
	return rc.makeRequest(i, nil, "GET", path, options)
}

func (rc *RundeckClient) Delete(path string, options interface{}) error {
	var b []byte
	return rc.makeRequest(&b, nil, "DELETE", path, options)
}

func (rc *RundeckClient) Post(i *[]byte, path string, data []byte, options interface{}) error {
	return rc.makeRequest(i, data, "POST", path, options)
}

func (rc *RundeckClient) Put(i *[]byte, path string, data []byte, options interface{}) error {
	return rc.makeRequest(i, data, "PUT", path, options)
}

func (client *RundeckClient) makeRequest(i *[]byte, payload []byte, method string, path string, params interface{}) error {
	qs := url.Values{}
	if params != nil {
		for q, p := range params.(map[string]string) {
			qs.Add(q, p)
		}
	}
	base_req_path := client.Config.BaseURL + "/api/13/" + path
	u, err := url.Parse(base_req_path)
	if err != nil {
		return err
	}
	if params != nil && len(params.(map[string]string)) != 0 {
		u.RawQuery = qs.Encode()
	}
	headers := http.Header{}
	headers.Add("user-agent", "rundeck-go.v13")
	jar, _ := cookiejar.New(nil)
	client.Client.Client.Jar = jar
	if client.Config.AuthMethod == "basic" {
		authQs := url.Values{}
		authQs.Add("j_username", client.Config.Username)
		authQs.Add("j_password", client.Config.Password)
		authPayload := bytes.NewBuffer(nil)
		base_auth_url := client.Config.BaseURL + "/j_security_check"
		cookieReq := napping.Request{
			Url:        base_auth_url,
			Params:     &authQs,
			Method:     "POST",
			RawPayload: true,
			Payload:    authPayload,
		}
		r, err := client.Client.Send(&cookieReq)
		if err != nil {
			return err
		}
		if r.Status() != 200 {
			return errors.New(r.RawText())
		}
	} else {
		headers.Add("X-Rundeck-Auth-Token", client.Config.Token)
	}
	headers.Add("Accept", "application/xml")
	headers.Add("user-agent", "rundeck-go.v13")
	req := napping.Request{
		Url:                 base_req_path,
		Header:              &headers,
		Params:              &qs,
		Method:              method,
		RawPayload:          true,
		Payload:             bytes.NewBuffer(payload),
		CaptureResponseBody: true,
	}
	r, err := client.Client.Send(&req)
	if err != nil {
		return err
	} else {
		if r.Status() == 404 {
			errormsg := fmt.Sprintf("No such item (%s)", r.Status())
			return errors.New(errormsg)
		}
		if r.Status() == 204 {
			return nil
		}
		if (r.Status() < 200) || (r.Status() > 299) {
			var data RundeckError
			xml.Unmarshal([]byte(r.RawText()), &data)
			errormsg := fmt.Sprintf("non-2xx response (code: %d): %s", r.Status(), data.Message)
			return errors.New(errormsg)
		} else {
			b := r.ResponseBody.Bytes()
			*i = append(*i, b...)
			return nil
		}
	}
}
