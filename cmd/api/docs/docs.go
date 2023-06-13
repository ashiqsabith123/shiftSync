// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/admin/signin": {
            "get": {
                "description": "api for admin to signin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin -Sign in"
                ],
                "summary": "Api for get admin signin page",
                "operationId": "Admin Sign In",
                "responses": {
                    "200": {
                        "description": "Welcome to sign in page",
                        "schema": {
                            "$ref": "#/definitions/request.LoginStruct"
                        }
                    }
                }
            },
            "post": {
                "description": "api for admin to signin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin -Sign in"
                ],
                "summary": "Api for post admin signin details",
                "operationId": "Admin Sign In",
                "responses": {
                    "200": {
                        "description": "Succesfuly logged in",
                        "schema": {
                            "$ref": "#/definitions/request.LoginStruct"
                        }
                    },
                    "400": {
                        "description": "Username and password is mandatory",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Unable to login",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/attendance": {
            "get": {
                "description": "api for get employees attendances",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "access attendance of employees",
                "operationId": "Attendance",
                "responses": {
                    "200": {
                        "description": "successfully fetched attendance",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "failed to get attendance",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "employee id not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/employee/leave/apply": {
            "post": {
                "description": "Api for employees to apply leave",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Apply leave"
                ],
                "operationId": "Apply leave",
                "parameters": [
                    {
                        "description": "Leave request details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Leave"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully applied for leave",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "error while applying leave",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "employee id not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/employee/leave/status": {
            "get": {
                "description": "Api for employees to get leaave status/history",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "operationId": "Leave status",
                "responses": {
                    "200": {
                        "description": "successfully fetched leave status/history",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "no leave history found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "employee id not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/employee/signin": {
            "get": {
                "description": "api for employees to signin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Sign in"
                ],
                "summary": "Api for get signin page",
                "operationId": "Sign In",
                "responses": {
                    "200": {
                        "description": "Welcome to sign in page",
                        "schema": {
                            "$ref": "#/definitions/request.LoginStruct"
                        }
                    }
                }
            },
            "post": {
                "description": "api for employees to signin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Sign in"
                ],
                "summary": "Api for post signin details",
                "operationId": "Sign In",
                "responses": {
                    "200": {
                        "description": "Succesfuly logged in",
                        "schema": {
                            "$ref": "#/definitions/request.LoginStruct"
                        }
                    },
                    "400": {
                        "description": "Username and password is mandatory",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Unable to login",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/employee/signup": {
            "get": {
                "description": "api for employees to signup",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Sign up"
                ],
                "summary": "Api for get signup page",
                "operationId": "Signup",
                "responses": {
                    "200": {
                        "description": "Welcome to signup page",
                        "schema": {
                            "$ref": "#/definitions/request.SignUp"
                        }
                    }
                }
            },
            "post": {
                "description": "api for employees to signup",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Sign up"
                ],
                "summary": "Api for post signup details",
                "operationId": "Signup",
                "parameters": [
                    {
                        "description": "Sign up details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Otp senEmployeed succesfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "User already exist",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "unable to signup",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/employee/signup/verify-otp": {
            "post": {
                "description": "api for employees to verify otp",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee -Sign up"
                ],
                "summary": "Api for post otp",
                "operationId": "Verify otp",
                "parameters": [
                    {
                        "description": "Otp",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OTPStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Account Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid otp",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Unable to find details",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Employee": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.Leave": {
            "type": "object",
            "properties": {
                "fromdate": {
                    "type": "string"
                },
                "leavetype": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "todate": {
                    "type": "string"
                }
            }
        },
        "request.LoginStruct": {
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
        "request.OTPStruct": {
            "type": "object",
            "properties": {
                "otp": {
                    "type": "string"
                }
            }
        },
        "request.SignUp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
