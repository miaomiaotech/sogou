package sogou

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func sendRequest(ctx context.Context, method, urlStr string, body io.Reader, f func(*http.Request) error) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, nil, fmt.Errorf("http.NewRequest error: %v", err)
	}
	req = req.WithContext(ctx)
	if f != nil {
		if err := f(req); err != nil {
			return nil, nil, fmt.Errorf("f error: %v", err)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("http request error: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("ioutil.ReadAll error: %v", err)
	}
	return resp, respBody, nil
}

func addCookies(req *http.Request, cookies []*http.Cookie) *http.Request {
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	return req
}
