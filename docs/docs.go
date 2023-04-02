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
            "post": {
                "description": "Creating new member data",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
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
                    "application/json"
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
                        "description": "Data trainer",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TrainerBody"
                        }
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
                    "application/json"
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
                        "description": "Update Trainer data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TrainerBody"
                        }
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
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "member_address": {
                    "type": "string"
                },
                "member_contact": {
                    "type": "string"
                },
                "member_email": {
                    "type": "string"
                },
                "member_gender": {
                    "type": "string"
                },
                "member_name": {
                    "type": "string"
                },
                "member_nik": {
                    "type": "integer"
                },
                "member_package": {
                    "type": "string"
                },
                "member_wight": {
                    "type": "integer"
                },
                "trainer": {
                    "$ref": "#/definitions/models.Trainer"
                },
                "trainer_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.MemberBody": {
            "type": "object",
            "properties": {
                "member_address": {
                    "type": "string"
                },
                "member_contact": {
                    "type": "string"
                },
                "member_email": {
                    "type": "string"
                },
                "member_gender": {
                    "type": "string"
                },
                "member_name": {
                    "type": "string"
                },
                "member_nik": {
                    "type": "integer"
                },
                "member_package": {
                    "type": "string"
                },
                "member_wight": {
                    "type": "integer"
                },
                "trainer_id": {
                    "type": "integer"
                }
            }
        },
        "models.Trainer": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Member"
                    }
                },
                "trainer_address": {
                    "type": "string"
                },
                "trainer_birthday": {
                    "type": "string"
                },
                "trainer_contact": {
                    "type": "string"
                },
                "trainer_email": {
                    "type": "string"
                },
                "trainer_gender": {
                    "type": "string"
                },
                "trainer_name": {
                    "type": "string"
                },
                "trainer_nik": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.TrainerBody": {
            "type": "object",
            "properties": {
                "trainer_address": {
                    "type": "string"
                },
                "trainer_birthday": {
                    "type": "string"
                },
                "trainer_contact": {
                    "type": "string"
                },
                "trainer_email": {
                    "type": "string"
                },
                "trainer_gender": {
                    "type": "string"
                },
                "trainer_name": {
                    "type": "string"
                },
                "trainer_nik": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Jagona Gym API",
	Description:      "This is a sample swagger for Jagona Gym API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
