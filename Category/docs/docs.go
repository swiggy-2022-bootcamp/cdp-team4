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
        "termsOfService": "http://swagger.io/terms/",
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
        "/": {
            "get": {
                "description": "this api is used to check whether category service is up and running or not",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health of category service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/categories/": {
            "get": {
                "description": "This Handle allows admin to fetch all the categories in the datastore",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Get all Categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "post": {
                "description": "This Handle allows admin to create a new category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Add Category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "delete": {
                "description": "This Handle allows admin to delete categories, given Ids",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete Categories details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/categories/:id": {
            "get": {
                "description": "This Handle allows admin to get a category, given Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get Category details by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "put": {
                "description": "This Handle allows admin to update a category, given Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category details by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "delete": {
                "description": "This Handle allows admin to delete a category, given Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete Category details by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "category Service",
	Description:      "Category Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
