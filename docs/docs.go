// Code generated by swaggo/swag. DO NOT EDIT.

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
            "email": "support@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthy": {
            "get": {
                "description": "Returns the health status of the service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Check health status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Healthy"
                        }
                    }
                }
            }
        },
        "/topsecret": {
            "post": {
                "description": "Decodes a message and calculates the location of the sender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret"
                ],
                "summary": "Decode message and determine location",
                "parameters": [
                    {
                        "description": "Top Secret DTO",
                        "name": "topSecretDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopSecretDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TopsecretResponse"
                        }
                    }
                }
            }
        },
        "/topsecret_split": {
            "get": {
                "description": "Retrieves the most recent calculated position and decoded message from split data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret"
                ],
                "summary": "Retrieve split data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TopsecretResponse"
                        }
                    }
                }
            }
        },
        "/topsecret_split/{satellite_name}": {
            "post": {
                "description": "Saves the message and distance data for a specific satellite",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret"
                ],
                "summary": "Save satellite data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Satellite Name",
                        "name": "satellite_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Satellite Split DTO",
                        "name": "satelliteSplit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SatelliteSplit"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.TopSecretSplitResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app_internal_adapter_inbound_http_model.Position": {
            "type": "object",
            "properties": {
                "x": {
                    "type": "number"
                },
                "y": {
                    "type": "number"
                }
            }
        },
        "dto.Satellite": {
            "type": "object",
            "required": [
                "distance",
                "message",
                "name"
            ],
            "properties": {
                "distance": {
                    "type": "number"
                },
                "message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.SatelliteSplit": {
            "type": "object",
            "required": [
                "distance",
                "message"
            ],
            "properties": {
                "distance": {
                    "type": "number"
                },
                "message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.TopSecretDTO": {
            "type": "object",
            "required": [
                "satellites"
            ],
            "properties": {
                "satellites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Satellite"
                    }
                }
            }
        },
        "model.Healthy": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "model.TopSecretSplitResponse": {
            "type": "object",
            "properties": {
                "savedReceivedMessage": {
                    "type": "boolean"
                }
            }
        },
        "model.TopsecretResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "position": {
                    "$ref": "#/definitions/app_internal_adapter_inbound_http_model.Position"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Rebel Communication Decoder",
	Description:      "This is the Rebel Communication Decoder API documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
