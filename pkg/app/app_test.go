package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHelloHandler(t *testing.T) {

    want := "Hello there!"

    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintln(w, want)
    }))

    defer ts.Close()

    client := ts.Client()

    res, err := client.Get(ts.URL)

    if err != nil {
        t.Errorf("expected nil got %v", err)
    }

    data, err := io.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {

        t.Errorf("expected nil got %v", err)
    }

    got := strings.TrimSpace(string(data))
  if string(data) == nil {

        t.Errorf("data is nil")
    }
}
