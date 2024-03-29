package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HttpPostJSON(requestUrl string, requestBody string, response interface{}) error {
	httpResp, err := http.Post(requestUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(requestBody),
	)
	if err != nil {
		fmt.Printf("qq ai server error: %v", err)
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	fmt.Println(httpResp.Body)
	return json.NewDecoder(httpResp.Body).Decode(response)
}
