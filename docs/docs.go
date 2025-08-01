// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/ping": {
            "get": {
                "description": "Do ping for test connection with the API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PING"
                ],
                "summary": "Do ping",
                "responses": {
                    "200": {
                        "description": "PONG",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "Create new user in database and return the created user withour the password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User successfully created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "description": "Delete one user and return nil",
                "tags": [
                    "User"
                ],
                "summary": "Delete one user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully deleted"
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        },
        "/user/getAll": {
            "get": {
                "description": "Retreive all user in database and return this without password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Retrieve all user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Users successfully retrieves",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        },
        "/user/getOne": {
            "get": {
                "description": "Retrieve one user with id in token and return this without password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Retrieve one user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully retrieve",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Sign in one user with this credentials and return user's token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Sign in one user",
                "parameters": [
                    {
                        "description": "User's credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User's token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "patch": {
                "description": "Update on user and return nil",
                "tags": [
                    "User"
                ],
                "summary": "Update one user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User data update",
                        "name": "user_update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully updated"
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/core.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "occurredAt": {
                    "type": "string"
                },
                "originalErr": {}
            }
        },
        "model.Login": {
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
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "test"
                },
                "last_name": {
                    "type": "string",
                    "example": "test"
                },
                "password": {
                    "type": "string",
                    "example": "aaaAAA111"
                },
                "username": {
                    "type": "string",
                    "example": "test_test"
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
