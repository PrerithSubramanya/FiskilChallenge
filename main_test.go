package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler(t *testing.T) {
	for i := int64(1); i <= 10; i++ {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)
		res := rec.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
		body, err := io.ReadAll(res.Body)
		err = res.Body.Close()
		if err != nil {
			return
		}
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		num, err := strconv.ParseInt(string(body), 10, 64)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}
		if num != i {
			t.Errorf("expected %v; got %v", i, num)
		}
	}
}
