package web

import (
	"net/http"

	"github.com/Cobblestone-Bridge/TurtleOverlord/helpers"
	"github.com/zenazn/goji/web"
  "github.com/golang/glog"

  "encoding/json"
)

func (controller *Controller) BootstrapPost(c web.C, r *http.Request) (string, int){
  id, version := r.FormValue("id"), r.FormValue("version")

  glog.Info(id + version)

  config := controller.GetConfiguration(c)
  hash := helpers.HashFile(config.PublicPath + "/computer/test.txt")

  out := map[string]string{"setLabel": "zeus", "driverChecksum": hash}
  jsonOut, _ := json.Marshal(out);

  return string(jsonOut[:]), http.StatusOK
}
