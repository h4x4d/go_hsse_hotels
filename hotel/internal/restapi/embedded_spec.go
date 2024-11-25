// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "MTS HSSE Go project | Hotel svc",
    "title": "hotels.hotel",
    "version": "0.3.0"
  },
  "paths": {
    "/hotel": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Get suitable hotels",
        "operationId": "get_hotels",
        "parameters": [
          {
            "type": "string",
            "name": "city",
            "in": "query"
          },
          {
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "enum": [
              1,
              2,
              3,
              4,
              5
            ],
            "type": "integer",
            "name": "hotel_class",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Hotel"
              }
            }
          },
          "404": {
            "description": "Suitable hotels not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Create hotel",
        "operationId": "create_hotel",
        "parameters": [
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "int64"
                }
              }
            }
          },
          "400": {
            "description": "Incorrect data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "No access",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/hotel/{hotel_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Get hotel info by id",
        "operationId": "get_hotel_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to return",
            "name": "hotel_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          },
          "404": {
            "description": "Hotel not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Update hotel",
        "operationId": "update_hotel",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to change",
            "name": "hotel_id",
            "in": "path",
            "required": true
          },
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          },
          "400": {
            "description": "Incorrect data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "No access",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "error_status_code"
      ],
      "properties": {
        "error_message": {
          "type": "string"
        },
        "error_status_code": {
          "type": "integer"
        }
      }
    },
    "Hotel": {
      "type": "object",
      "required": [
        "name",
        "city",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string",
          "example": "Red Square №1"
        },
        "city": {
          "type": "string",
          "example": "Moscow"
        },
        "cost": {
          "type": "integer",
          "format": "int64"
        },
        "hotel_class": {
          "description": "number of stars of hotel",
          "type": "integer",
          "enum": [
            0,
            1,
            2,
            3,
            4,
            5
          ]
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "Radisson"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Hotel operations",
      "name": "hotel"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "MTS HSSE Go project | Hotel svc",
    "title": "hotels.hotel",
    "version": "0.3.0"
  },
  "paths": {
    "/hotel": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Get suitable hotels",
        "operationId": "get_hotels",
        "parameters": [
          {
            "type": "string",
            "name": "city",
            "in": "query"
          },
          {
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "enum": [
              1,
              2,
              3,
              4,
              5
            ],
            "type": "integer",
            "name": "hotel_class",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Hotel"
              }
            }
          },
          "404": {
            "description": "Suitable hotels not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Create hotel",
        "operationId": "create_hotel",
        "parameters": [
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "int64"
                }
              }
            }
          },
          "400": {
            "description": "Incorrect data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "No access",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/hotel/{hotel_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Get hotel info by id",
        "operationId": "get_hotel_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to return",
            "name": "hotel_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          },
          "404": {
            "description": "Hotel not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Update hotel",
        "operationId": "update_hotel",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to change",
            "name": "hotel_id",
            "in": "path",
            "required": true
          },
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hotel"
            }
          },
          "400": {
            "description": "Incorrect data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "No access",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "error_status_code"
      ],
      "properties": {
        "error_message": {
          "type": "string"
        },
        "error_status_code": {
          "type": "integer"
        }
      }
    },
    "Hotel": {
      "type": "object",
      "required": [
        "name",
        "city",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string",
          "example": "Red Square №1"
        },
        "city": {
          "type": "string",
          "example": "Moscow"
        },
        "cost": {
          "type": "integer",
          "format": "int64"
        },
        "hotel_class": {
          "description": "number of stars of hotel",
          "type": "integer",
          "enum": [
            0,
            1,
            2,
            3,
            4,
            5
          ]
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "Radisson"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Hotel operations",
      "name": "hotel"
    }
  ]
}`))
}