// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "YPSI SAS",
            "url": "https://www.ypsi.fr/",
            "email": "info@ypsi.fr"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book": {
            "post": {
                "description": "Create book in DB",
                "tags": [
                    "book"
                ],
                "summary": "Create book",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BookCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Books return",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Error on request"
                    },
                    "500": {
                        "description": "Error to create"
                    }
                }
            }
        },
        "/book/{id}": {
            "delete": {
                "description": "Delete one book in DB by ID",
                "tags": [
                    "book"
                ],
                "summary": "Delete one book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error on request"
                    },
                    "404": {
                        "description": "Error book not find"
                    },
                    "500": {
                        "description": "Error to delete"
                    }
                }
            }
        },
        "/books": {
            "get": {
                "description": "Get all books in DB",
                "tags": [
                    "book"
                ],
                "summary": "Get all books",
                "responses": {
                    "200": {
                        "description": "Books return",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksList"
                        }
                    },
                    "500": {
                        "description": "Error to find"
                    }
                }
            }
        },
        "/books/author/{author}": {
            "get": {
                "description": "Get all books in DB by author name",
                "tags": [
                    "book"
                ],
                "summary": "Get all books by author name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author's name",
                        "name": "author",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Books return",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksList"
                        }
                    },
                    "400": {
                        "description": "Error on request"
                    },
                    "500": {
                        "description": "Error to find book by author"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.BookCreate": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.BooksList": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Book"
                    }
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
        "models.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
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
                "title": {
                    "type": "string"
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server for demonstration.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
