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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/private/logged": {
            "get": {
                "description": "Retrieves user based on access token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Retrieves user based on access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/private/logout": {
            "post": {
                "description": "Remove access token from cookie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Remove access token from cookie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/private/users": {
            "get": {
                "description": "Retrieves all users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Retrieves all users.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update selected user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Update selected user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Create a new user.",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/private/users/{id}": {
            "get": {
                "description": "Retrieves user based on given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Retrieves user based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user based on given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private Routes"
                ],
                "summary": "Delete user based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/public/healthCheck": {
            "get": {
                "description": "Healthcheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public Routes"
                ],
                "summary": "Healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public/login": {
            "post": {
                "description": "Public route to authenticate a user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public Routes"
                ],
                "summary": "Public route to authenticate a user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/public/register": {
            "post": {
                "description": "Public route to create a new user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public Routes"
                ],
                "summary": "Public route to create a new user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UserRegister": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Go test Swagger API",
	Description: "Swagger API for Golang Test.",
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
