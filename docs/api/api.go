// Package docs Code generated by swaggo/swag. DO NOT EDIT
package api

import (
	"github.com/swaggo/swag"
	"io/ioutil"
)

func init() {
	var readFile, _ = ioutil.ReadFile("docs/api/swagger.yaml")
	var SwaggerInfo = &swag.Spec{
		Version:          "",
		Host:             "",
		BasePath:         "",
		Schemes:          []string{},
		Title:            "",
		Description:      "",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(readFile),
		LeftDelim:        "{{",
		RightDelim:       "}}",
	}
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
