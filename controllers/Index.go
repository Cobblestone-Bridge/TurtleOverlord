package controllers

import (
  "net/http"
  "TurtleOverlord/views"
  //"fmt"
  //"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
  err := views.Templates.ExecuteTemplate(w, "indexPage", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }
}
