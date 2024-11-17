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
        "/bids": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new bid with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bid"
                ],
                "summary": "Create a new bid",
                "parameters": [
                    {
                        "description": "Bid Details",
                        "name": "bid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ApiCreateBidReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
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
        "/bids/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an Bid by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bid"
                ],
                "summary": "Delete an Bid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bid deleted successfully",
                        "schema": {
                            "type": "string"
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
        "/bids/get": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Bid by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bid"
                ],
                "summary": "Get Bid by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BidRes"
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
        "/bids/list": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all Bids with optional filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bid"
                ],
                "summary": "Get all Bids",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TenderId",
                        "name": "tender_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ContractorId",
                        "name": "contractor_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "DeliveryTime",
                        "name": "delivery_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Price",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortType",
                        "name": "sort_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllBidRes"
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
        "/bids/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an Bid's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bid"
                ],
                "summary": "Update an Bid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "description": "Bid Update Details",
                        "name": "bid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ApiUpdateBidReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
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
        "/client/tenders": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "List Tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tender"
                ],
                "summary": "List Tender",
                "responses": {
                    "200": {
                        "description": "Success list tenders",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create Tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tender"
                ],
                "summary": "Create Tender",
                "parameters": [
                    {
                        "description": "Create Tender",
                        "name": "tender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateTenderReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Create Tender",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/client/tenders/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update Tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tender"
                ],
                "summary": "Update Tender",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tender ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "description": "Create Tender",
                        "name": "tender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender status updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tender"
                ],
                "summary": "Delete Tender",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tender ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/img-upload": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "File upload",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file-upload"
                ],
                "summary": "File upload",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT tokens",
                        "schema": {
                            "$ref": "#/definitions/models.TokenRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "DeleteProfile",
                "responses": {
                    "200": {
                        "description": "Deleted profile",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "UpdateProfile",
                "parameters": [
                    {
                        "description": "Update request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated profile",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Registration request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "JWT tokens",
                        "schema": {
                            "$ref": "#/definitions/models.TokenRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ApiCreateBidReq": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "string"
                },
                "contractor_id": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "tender_id": {
                    "type": "string"
                }
            }
        },
        "models.ApiUpdateBidReq": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "models.BidRes": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "string"
                },
                "contractor_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "tender_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateTenderReq": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "number"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.GetAllBidRes": {
            "type": "object",
            "properties": {
                "bids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BidRes"
                    }
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "models.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.RegisterReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.TokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "models.UpdateProfile": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UpdateStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "models.UserRes": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pima",
	Description:      "API for Pima",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
