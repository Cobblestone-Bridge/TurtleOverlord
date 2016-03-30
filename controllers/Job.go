package controllers

import (
  "net/http"
  "fmt"
  "encoding/json"
)

func Job(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    http.Error(w, "error", http.StatusMethodNotAllowed)
  } else {
      r.ParseForm()
      fmt.Println("label:", r.PostFormValue("label"))

      out := map[string]string{"jsonRPC": "2.0", "method": "forward","id": "SFE4FRG3"}
      jsonOut, _ := json.Marshal(out);
      fmt.Fprintf(w,string(jsonOut))
  }
}
