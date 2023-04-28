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
        "/api/equipment": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Get all equipments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Equipment"
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
                "description": "Creating new Equipment data",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Create Equipment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nama Alat",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Jumlah Alat",
                        "name": "count",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Status Alat: ready atau broken",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Foto Alat",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success create equipment",
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
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
        "/api/equipment/{id}": {
            "get": {
                "description": "Get a single Equipment data by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Get a single Equipment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Equipment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
                        }
                    },
                    "404": {
                        "description": "Equipment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to get equipment",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Equipment data by ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Update Equipment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Equipment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Nama Alat",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Jumlah Alat",
                        "name": "count",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Status Alat: ready atau broken",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Foto Alat",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success create equipment",
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
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
            },
            "delete": {
                "description": "Delete Equipment data by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Delete Equipment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Equipment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Delete Equipment Data",
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
                            "$ref": "#/definitions/controller.memberRequest"
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
                            "$ref": "#/definitions/controller.memberRequest"
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
                        "type": "integer",
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
                        "description": "Instagram trainer",
                        "name": "instagram",
                        "in": "formData"
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
                        "type": "string",
                        "description": "Deskripsi singkat trainer",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Pengalaman trainer",
                        "name": "experience",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Spesialisasi trainer",
                        "name": "specialization",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Pencapaian/Sertifikasi/Lisensi trainer",
                        "name": "achievement",
                        "in": "formData"
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
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.memberRequest": {
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
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "instagram": {
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
                "status": {
                    "type": "string"
                },
                "trainer_id": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
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
        "models.Equipment": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Member": {
            "type": "object",
            "properties": {
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
                "detail_member": {
                    "$ref": "#/definitions/models.MemberDetail"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                },
                "status": {
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
                }
            }
        },
        "models.MemberDetail": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instagram": {
                    "type": "string"
                },
                "memberID": {
                    "type": "integer"
                },
                "nik": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.Trainer": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                },
                "instagram": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "specialization": {
                    "type": "string"
                },
                "trainerDetail": {
                    "$ref": "#/definitions/models.TrainerDetail"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.TrainerDetail": {
            "type": "object",
            "properties": {
                "achievement": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "experience": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
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
                "nik": {
                    "type": "integer"
                },
                "trainerID": {
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
