package views

import (
  "html/template"
)

var Templates = template.Must(template.ParseFiles("../src/TurtleOverlord/views/head.html", "../src/TurtleOverlord/views/index.html"))
