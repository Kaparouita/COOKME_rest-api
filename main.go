package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

  url := "https://google.serper.dev/search"
  method := "POST"

  payload := strings.NewReader(`{"q":"Tomato Xalkiadakis","gl":"gr"}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("X-API-KEY", "b91571be77615909bc89930b2008a450a9132d69")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}