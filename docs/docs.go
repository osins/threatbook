// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ip/query": {
            "get": {
                "description": "获取IP相关地理位置、ASN信息，综合判定威胁类型如:远控（C2）、傀儡机（Zombie）、失陷主机（Compromised）、扫描（Scanner）、钓鱼（Phishing）等，相关攻击团伙或安全事件标签，原始情报，相关样本信息等。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "可针对业务日志，以及从防火墙、WAF等安防设备中获取的外部IP，进行分析。",
                "operationId": "ip-query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "resource string",
                        "name": "resource",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "exclude string",
                        "name": "exclude",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "lang string",
                        "name": "lang",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tb.Response"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "tb.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "response_code": {
                    "type": "integer"
                },
                "verbose_msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// swaggerInfo type
type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

// s type
type s struct{}

// ReadDoc inner method
func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

// init method
func init() {
	swag.Register(swag.Name, &s{})
}
