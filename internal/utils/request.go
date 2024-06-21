package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ReadReqBody(req *http.Request, v any) error {
	r := bufio.NewReader(req.Body)
	bstr := ""
	for {
		l, err := r.ReadString('\n')
		bstr += l
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	if err := json.Unmarshal([]byte(bstr), v); err != nil {
		return err
	}
	return nil
}

func ReadReqHeader(req *http.Request, headerName string) *string {
	headerValue := strings.Trim(req.Header.Get(headerName), " ")
	if headerValue == "" {
		return nil
	}
	return &headerValue
}
