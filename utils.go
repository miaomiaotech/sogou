package sogou

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func GobEncode(value any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func GobDecode(file io.ReadCloser, data any) error {
	return gob.NewDecoder(file).Decode(data)
}

func WriteCookiesTo(cookies []*http.Cookie, path string) error {
	bs, err := GobEncode(cookies)
	if err != nil {
		return fmt.Errorf("GobEncode error: %v", err)
	} else {
		if err := os.MkdirAll(filepath.Dir(cookiePath), 0755); err != nil {
			return fmt.Errorf("MkdirAll error: %v", err)
		}
		if err := os.WriteFile(cookiePath, bs, 0644); err != nil {
			return fmt.Errorf("write %s error: %v", cookiePath, err)
		}
	}
	return nil
}

func ReadCookiesFrom(path string) ([]*http.Cookie, error) {
	file, err := os.Open(cookiePath)
	if err != nil {
		return nil, fmt.Errorf("read %s error: %v", cookiePath, err)
	}
	defer file.Close()

	var cookies []*http.Cookie
	err = GobDecode(file, &cookies)
	if err != nil {
		return nil, fmt.Errorf("GobDecode error: %v", err)
	}
	return cookies, nil
}
