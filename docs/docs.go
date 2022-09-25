// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alan Syu",
            "email": "qoo753951ooq@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/order": {
            "get": {
                "description": "訂單列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trade"
                ],
                "summary": "getList",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2001-01-01",
                        "description": "開始時間 (yyyy-MM-dd)",
                        "name": "starttime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2001-01-02",
                        "description": "結束時間 (yyyy-MM-dd)",
                        "name": "endtime",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "取前幾筆",
                        "name": "top",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vo.OrderVO"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "新增訂單",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trade"
                ],
                "summary": "addOne",
                "parameters": [
                    {
                        "type": "boolean",
                        "default": false,
                        "name": "buy",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "訂單日期(yyyy-MM-dd)",
                        "name": "date",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": 0,
                        "name": "limit_price",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "default": 0,
                        "name": "market_price",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "quantity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "name": "sell",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.OrderVO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/{id}": {
            "get": {
                "description": "單筆訂單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trade"
                ],
                "summary": "getOne",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.OrderVO"
                        }
                    }
                }
            },
            "put": {
                "description": "編輯訂單",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "trade"
                ],
                "summary": "editOne",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body for OrderPutVO content",
                        "name": "putVO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.OrderPutVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除訂單",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "trade"
                ],
                "summary": "deleteOne",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "vo.OrderPutVO": {
            "type": "object",
            "required": [
                "date",
                "quantity"
            ],
            "properties": {
                "buy": {
                    "type": "boolean",
                    "example": false
                },
                "date": {
                    "type": "string",
                    "example": "訂單日期(yyyy-MM-dd)"
                },
                "limit_price": {
                    "type": "number",
                    "example": 0
                },
                "market_price": {
                    "type": "number",
                    "example": 0
                },
                "quantity": {
                    "type": "integer",
                    "example": 0
                },
                "sell": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "vo.OrderVO": {
            "type": "object",
            "properties": {
                "buy": {
                    "type": "boolean"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "limit_price": {
                    "type": "number"
                },
                "market_price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "sell": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/home-work-z-api",
	Schemes:          []string{"http"},
	Title:            "Swagger Home Work Z API",
	Description:      "Home work Z Swagger 說明文件",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
