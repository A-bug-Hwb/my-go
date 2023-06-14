// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/WrMenu/getWrMenuById": {
            "get": {
                "description": "根据菜单id获取菜单详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WrMenu"
                ],
                "summary": "获取详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜单id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "0": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/WrMenu/getWrMenuList": {
            "post": {
                "description": "获取所有菜单列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WrMenu"
                ],
                "summary": "获取列表",
                "parameters": [
                    {
                        "description": "实体类信息",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pojo.WrMenuVo"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pojo.WrMenuVo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "menuName": {
                    "type": "string"
                },
                "orderNum": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "visible": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}