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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/auth": {
            "post": {
                "description": "Request to authentication token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "request to authentication token",
                "parameters": [
                    {
                        "description": "Authenticate existing user",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.EmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/auth/activate": {
            "post": {
                "description": "Request to verification token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Institution type",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Institution type",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Authenticate user and get jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Login to account",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/types.JWTResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Add a new user",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Registration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.AuthRequest": {
            "type": "object",
            "required": [
                "email-address",
                "token"
            ],
            "properties": {
                "email-address": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "types.EmailRequest": {
            "type": "object",
            "required": [
                "email_address"
            ],
            "properties": {
                "email_address": {
                    "type": "string"
                }
            }
        },
        "types.JWTResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/types.UserProfile"
                }
            }
        },
        "types.Registration": {
            "type": "object",
            "required": [
                "email_address",
                "first_name",
                "last_name"
            ],
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "email_address": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "types.UserProfile": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "email_address": {
                    "type": "string"
                },
                "email_verified": {
                    "type": "boolean"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_login": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "utils.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/directory/api/v1",
	Schemes:          []string{},
	Title:            "Business directory API",
	Description:      "This is the API for business directory api..",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
