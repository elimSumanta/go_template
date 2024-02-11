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
        "/private/main/ads/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mainflow"
                ],
                "summary": "Get Ads List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseResponseSchema"
                        }
                    }
                }
            }
        },
        "/private/main/crypto/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mainflow"
                ],
                "summary": "Get Crypto List",
                "operationId": "mainflow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseResponseSchema"
                        }
                    }
                }
            }
        },
        "/private/user/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User Profile",
                "operationId": "user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseResponseSchema"
                        }
                    }
                }
            }
        },
        "/public/healthcheck": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.BaseResponseSchema": {
            "type": "object",
            "properties": {
                "error": {},
                "meta": {},
                "result": {},
                "success": {
                    "type": "boolean"
                },
                "traceId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Bitwyre P2P Develop",
	Description:      "This is P2P Gateway Documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
