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
        "contact": {
            "name": "davidleitw",
            "email": "davidleitw@gmail.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user/userLogin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用戶登入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶電子信箱(帳號)",
                        "name": "account",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用戶密碼",
                        "name": "password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serialization.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/serialization.Response"
                        }
                    }
                }
            }
        },
        "/api/user/userLogout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用戶登出",
                "responses": {
                    "200": {
                        "description": "登出成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/serialization.Response"
                        }
                    }
                }
            }
        },
        "/api/user/userRegister": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用戶註冊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "填入用戶的電子信箱",
                        "name": "UserMail",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "填入用戶的使用者名稱 可隨時更改",
                        "name": "UserName",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "填入用戶的密碼 需要八碼以上",
                        "name": "UserPass",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密碼確認, 需要UserPass參數相同",
                        "name": "UserPassConfirm",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用戶的電話號碼 如果使用者沒填 前端會補上00000000",
                        "name": "USerRid",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serialization.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/serialization.Response"
                        }
                    }
                }
            }
        },
        "/api/user/usercheckLogin": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "確認登入狀態",
                "responses": {
                    "200": {
                        "description": "false",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serialization.Response": {
            "type": "object"
        }
    }
}`

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
	Version:     "0.0.1",
	Host:        "http://espresso.nctu.me:8080/",
	BasePath:    "/api/",
	Schemes:     []string{},
	Title:       "Espresso Example API File",
	Description: "Expreesso calendar的api文檔, 方便串接前後端",
}

type s struct{}

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

func init() {
	swag.Register(swag.Name, &s{})
}