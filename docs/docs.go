// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-23 21:21:56.1068985 +0800 CST m=+0.054998601

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
            "name": "API Support"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/movie/batch/save": {
            "post": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "tags": [
                    "电影模块"
                ],
                "summary": "批量保存电影信息",
                "parameters": [
                    {
                        "description": "批量保存电影基本信息",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/movie/delete/{id}": {
            "get": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "tags": [
                    "电影模块"
                ],
                "summary": "根据id删除电影信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "要删除的电影id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/movie/query/{id}": {
            "get": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "tags": [
                    "电影模块"
                ],
                "summary": "获取某个电影的信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "电影id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/movie/save": {
            "post": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "tags": [
                    "电影模块"
                ],
                "summary": "保存电影信息",
                "parameters": [
                    {
                        "description": "保存的电影基本信息",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Movie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "describe": {
                    "type": "string"
                },
                "directors": {
                    "type": "string"
                },
                "fileLength": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "languages": {
                    "type": "string"
                },
                "movieName": {
                    "type": "string"
                },
                "nations": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "47.98.164.130:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "gin实战",
	Description: "gin开发实战接口列表",
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
