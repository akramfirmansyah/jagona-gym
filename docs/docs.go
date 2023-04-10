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
            "email": "akram.firman@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/member": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Get all members",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creating new Member data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Create Member",
                "parameters": [
                    {
                        "description": "Data member",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MemberBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success create member",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/member/{id}": {
            "get": {
                "description": "Get a single Member data by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Get a single Member",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "404": {
                        "description": "Member not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to get trainer",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a specific Member data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Update Member",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Member data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MemberBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Member not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Member data by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Delete Member",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Delete Member Data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Data not Found!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/trainer": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Get all trainers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Trainer"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creating new Trainer data",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Create Trainer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name trainer",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NIK trainer",
                        "name": "nik",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tanggal Lahir trainer",
                        "name": "birthday",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email trainer",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kontak trainer",
                        "name": "contact",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Alamat trainer",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Gender trainer",
                        "name": "gender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image trainer",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success create trainer",
                        "schema": {
                            "$ref": "#/definitions/models.Trainer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/trainer/{id}": {
            "get": {
                "description": "Get a single Trainer data by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Get a single Trainer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Trainer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trainer"
                        }
                    },
                    "404": {
                        "description": "Trainer not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to get trainer",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a specific Trainer data",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Update Trainer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Trainer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name trainer",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NIK trainer",
                        "name": "nik",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tanggal Lahir trainer",
                        "name": "birthday",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email trainer",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kontak trainer",
                        "name": "contact",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Alamat trainer",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Gender trainer",
                        "name": "gender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image trainer",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trainer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Trainer not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Trainer data by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Delete Trainer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Trainer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Delete Trainer Data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Data not Found!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Member": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nik": {
                    "type": "integer"
                },
                "package": {
                    "type": "string"
                },
                "trainer": {
                    "$ref": "#/definitions/models.Trainer"
                },
                "trainer_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "wight": {
                    "type": "integer"
                }
            }
        },
        "models.MemberBody": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nik": {
                    "type": "integer"
                },
                "package": {
                    "type": "string"
                },
                "trainer_id": {
                    "type": "integer"
                },
                "wight": {
                    "type": "integer"
                }
            }
        },
        "models.Trainer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Member"
                    }
                },
                "name": {
                    "type": "string"
                },
                "nik": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Jagona Gym API",
	Description:      "This is a swagger for Jagona Gym API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}