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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/workspaces": {
            "get": {
                "description": "get workspace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Get all worskpaces by id",
                "operationId": "get-workspaces",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Workspace"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create workspace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Create worskpace",
                "operationId": "create-workspace",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.workspaceInput"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
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
                            "$ref": "#/definitions/rest.postBody"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/workspaces/{id}": {
            "get": {
                "description": "get workspace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Get worskpace by id",
                "operationId": "get-workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
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
                            "$ref": "#/definitions/entity.Workspace"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete workspace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Delete worskpace by id",
                "operationId": "delete-workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
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
                            "$ref": "#/definitions/rest.postBody"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.signInBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.tokenOutput"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.signUpBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.postBody"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/webjson.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Workspace": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "rest.postBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "rest.signInBody": {
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
        "rest.signUpBody": {
            "type": "object",
            "properties": {
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
        },
        "rest.tokenOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "rest.workspaceInput": {
            "type": "object",
            "properties": {
                "logo": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "webjson.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0// @host localhost:8080",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "trello-clone API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
