// @title 黑名单管理系统 API
// @version 1.0
// @description 黑名单管理系统的API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
package docs

import "github.com/swaggo/swag"

func init() {
	swag.Register(swag.Name, &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  docTemplate,
	})
}

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "description": "黑名单管理系统的API文档",
        "title": "黑名单管理系统 API",
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
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/merchants/login": {
            "post": {
                "description": "商户登录",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merchants"
                ],
                "summary": "商户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "api_key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "API Secret",
                        "name": "api_secret",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    }
}`
