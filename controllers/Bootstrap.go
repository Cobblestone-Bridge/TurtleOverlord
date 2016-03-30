package controllers

import (
  "net/http"
  "fmt"
  "encoding/json"
  "hash/crc32"
  "io/ioutil"
  "strings"
  //"TurtleOverlord/views"
)

/*
** Bootstrap controller. The turtle calls this controller upon start. This controller tells the turtle what
** Label this turtle should have and what hash his driver should be at.
*/
func Bootstrap(w http.ResponseWriter, r *http.Request) {
      if r.Method == "GET" {
        http.Error(w, "error", http.StatusMethodNotAllowed)
      } else {
          r.ParseForm()
          fmt.Println("id:", r.PostFormValue("id"))
          fmt.Println("version:", r.PostFormValue("version"))

          // Read the required file into a string,
          val, err := ioutil.ReadFile("../src/TurtleOverlord/static/computer/test.txt")
          fmt.Println(err)

          // Trim all leading and tailing spaces.
          s := strings.TrimSpace(string(val[:]))

          crc32q := crc32.MakeTable(0xEDB88320)
          h := (uint32)(crc32.Checksum([]byte(s), crc32q))
          hash := fmt.Sprintf("%X", h)

          out := map[string]string{"setLabel": "zeus", "driverChecksum": hash}
          jsonOut, _ := json.Marshal(out);
          fmt.Fprintf(w,string(jsonOut))
      }
}
