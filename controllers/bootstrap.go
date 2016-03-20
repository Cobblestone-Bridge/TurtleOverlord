package controllers

import (
	"github.com/astaxie/beego"
	"crypto/sha256"
	"encoding/hex"
)

// operations for Bootstrap
type BootstrapController struct {
	beego.Controller
}

func (c *BootstrapController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Bootstrap
// @Param	body		body 	models.Bootstrap	true		"body for Bootstrap content"
// @Success 201 {object} models.Bootstrap
// @Failure 403 body is empty
// @router / [post]
func (c *BootstrapController) Post() {
	//computerID := c.GetString("id")
	computerVersion := c.GetString("version")

	if (computerVersion == "CraftOS 1.7") {
		//("/computerDrivers/driver1_7.lua")

		byteValues := []byte("test")
		md5Sum := sha256.Sum256(byteValues)
		driverCheckSum := hex.EncodeToString(md5Sum[:])

		c.Data["json"] = map[string]interface{}{"driverChecksum": driverCheckSum, "setLabel": "Zeus"}
		c.ServeJSON()
	}
}

// @Title Get
// @Description get Bootstrap by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Bootstrap
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BootstrapController) GetOne() {

}

// @Title Get All
// @Description get Bootstrap
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Bootstrap
// @Failure 403
// @router / [get]
func (c *BootstrapController) GetAll() {

}

// @Title Update
// @Description update the Bootstrap
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Bootstrap	true		"body for Bootstrap content"
// @Success 200 {object} models.Bootstrap
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BootstrapController) Put() {

}

// @Title Delete
// @Description delete the Bootstrap
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BootstrapController) Delete() {

}
