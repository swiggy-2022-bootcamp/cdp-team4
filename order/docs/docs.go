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
            "name": "Suhas R",
            "email": "suhas7thfeb2000@gmail.com"
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
                "description": "use tocheck whether order service is up and running or not",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health of order service",
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
        "/confirm/:userid": {
            "post": {
                "description": "This Handle adds order from checkout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Confirm Order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.OrderConfirmResponseDTO"
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
        "/order": {
            "post": {
                "description": "This Handle allows admin to create new Order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "Create order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.OrderRecordDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
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
        "/order/:id": {
            "get": {
                "description": "This Handle returns Order given order id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.OrderRecordDTO"
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
                "description": "This Handle deletes order given order id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
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
        "/order/invoice/:orderid": {
            "get": {
                "description": "This generated invoice given order id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order Invoice given order id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order id",
                        "name": "orderid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.InvoiceDTO"
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
        "/order/status": {
            "put": {
                "description": "This Handle Update order status given order id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update order status",
                "parameters": [
                    {
                        "description": "Update order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.RequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
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
        "/order/status/:status": {
            "get": {
                "description": "This Handle returns Order given status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order by status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "status",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/app.OrderRecordDTO"
                            }
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
        "/order/user/:userid": {
            "get": {
                "description": "This Handle returns Order given user id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order by user id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/app.OrderRecordDTO"
                            }
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
        "/orders": {
            "get": {
                "description": "This Handle returns all of the orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get All Order records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/app.OrderRecordDTO"
                            }
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
    },
    "definitions": {
        "app.InvoiceDTO": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.ProductRecordDTO"
                    }
                },
                "reward_points": {
                    "type": "integer"
                },
                "shipping_price": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "total_cost": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "app.OrderConfirmResponseDTO": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "reward_points": {
                    "type": "integer"
                },
                "shipping_price": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "total_cost": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "app.OrderRecordDTO": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.ProductRecordDTO"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total_cost": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "app.ProductRecordDTO": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "app.RequestDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:7000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Gin Swagger",
	Description:      "Order Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
