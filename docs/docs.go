// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/v1/logs": {
            "post": {
                "description": "Stores logs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Save log",
                "parameters": [
                    {
                        "description": "LogEvent Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.LogEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/pipelines/{commitId}": {
            "get": {
                "description": "Gets pipeline by process id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pipeline"
                ],
                "summary": "Get by process id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "processId",
                        "name": "commitId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/pipelines/{processId}": {
            "get": {
                "description": "Gets logs by pipeline processId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pipeline"
                ],
                "summary": "Get Logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pipeline ProcessId",
                        "name": "processId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Record count",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/process_life_cycle_events": {
            "get": {
                "description": "Pulls auto trigger enabled steps",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessLifeCycle"
                ],
                "summary": "Pull Steps",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Agen name",
                        "name": "agent",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Pull size",
                        "name": "count",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Step type [BUILD, DEPLOY]",
                        "name": "step_type",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/v1.ProcessLifeCycleEvent"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Update reclaim step",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessLifeCycle"
                ],
                "summary": "Update Steps",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Step name",
                        "name": "step",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Process id",
                        "name": "processId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Process life cycle step status",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Stores process lifecycle event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessLifeCycle"
                ],
                "summary": "Save process lifecycle event",
                "parameters": [
                    {
                        "description": "ProcessLifeCycleEventList Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.ProcessLifeCycleEventList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/processes": {
            "get": {
                "description": "Get Process List or count process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get Process List or count process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company Id",
                        "name": "companyId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repositoryId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "App Id",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Commit Id",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Operation[countTodaysProcessByCompanyId]",
                        "name": "operation",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/v1.Process"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Stores process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Save process",
                "parameters": [
                    {
                        "description": "Process Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Process"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/processes/{processId}/steps/{step}": {
            "get": {
                "description": "Get Footmark List",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get Footmark List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Id",
                        "name": "processId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "step name",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/processes_events": {
            "post": {
                "description": "Stores Pipeline process event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessEvent"
                ],
                "summary": "Save Pipeline process event",
                "parameters": [
                    {
                        "description": "PipelineProcessEvent Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PipelineProcessEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.MetaData": {
            "type": "object",
            "properties": {
                "links": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "common.ResponseDTO": {
            "type": "object",
            "properties": {
                "_metadata": {
                    "$ref": "#/definitions/common.MetaData"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "unstructured.Unstructured": {
            "type": "object",
            "properties": {
                "object": {
                    "description": "Object is a JSON compatible map with string, float, int, bool, []interface{}, or\nmap[string]interface{}\nchildren.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "v1.CompanyMetadata": {
            "type": "object",
            "properties": {
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "number_of_concurrent_process": {
                    "type": "integer"
                },
                "total_process_per_day": {
                    "type": "integer"
                }
            }
        },
        "v1.LogEvent": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "footmark": {
                    "type": "string"
                },
                "log": {
                    "type": "string"
                },
                "processId": {
                    "type": "string"
                },
                "step": {
                    "type": "string"
                }
            }
        },
        "v1.Pipeline": {
            "type": "object",
            "properties": {
                "_metadata": {
                    "$ref": "#/definitions/v1.PipelineMetadata"
                },
                "api_version": {
                    "type": "string"
                },
                "label": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "$ref": "#/definitions/v1.PipelineApplyOption"
                },
                "process_id": {
                    "type": "string"
                },
                "steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Step"
                    }
                }
            }
        },
        "v1.PipelineApplyOption": {
            "type": "object",
            "properties": {
                "purging": {
                    "type": "string"
                }
            }
        },
        "v1.PipelineMetadata": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "company_metadata": {
                    "$ref": "#/definitions/v1.CompanyMetadata"
                }
            }
        },
        "v1.PipelineProcessEvent": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "process_id": {
                    "type": "string"
                }
            }
        },
        "v1.Process": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "commit_id": {
                    "type": "string"
                },
                "company_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "process_id": {
                    "type": "string"
                },
                "repository_id": {
                    "type": "string"
                }
            }
        },
        "v1.ProcessLifeCycleEvent": {
            "type": "object",
            "properties": {
                "agent": {
                    "type": "string"
                },
                "claim": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "next": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pipeline": {
                    "$ref": "#/definitions/v1.Pipeline"
                },
                "processId": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "step": {
                    "type": "string"
                },
                "stepType": {
                    "type": "string"
                },
                "trigger": {
                    "type": "string"
                }
            }
        },
        "v1.ProcessLifeCycleEventList": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ProcessLifeCycleEvent"
                    }
                }
            }
        },
        "v1.Step": {
            "type": "object",
            "properties": {
                "arg_data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "descriptors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/unstructured.Unstructured"
                    }
                },
                "env_data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "next": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "params": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                },
                "trigger": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Klovercloud-ci-event-bank API",
	Description: "Klovercloud-ci-event-bank API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
