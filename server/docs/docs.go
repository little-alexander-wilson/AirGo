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
        "contact": {
            "url": "https://github.com/ppoonk/AirGo"
        },
        "license": {
            "name": "GPL v3.0",
            "url": "https://github.com/ppoonk/AirGo/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/public/getBase64Captcha": {
            "get": {
                "description": "发送base64验证码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public_api"
                ],
                "summary": "发送base64验证码",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Base64CaptchaInfo"
                        }
                    }
                }
            }
        },
        "/api/public/getEmailCode": {
            "post": {
                "description": "发送邮箱验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public_api"
                ],
                "summary": "发送邮箱验证码",
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Base64CaptchaInfo": {
            "type": "object",
            "properties": {
                "b64s": {
                    "description": "响应时存base64数据，请求时存前端看到的验证码。响应，请求共用该结构体",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.9 版本",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AirGo",
	Description:      "AirGo前后分离多用户代理面板",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
