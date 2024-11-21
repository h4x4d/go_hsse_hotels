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
    "version": "0.2.3"
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
          },
          {
            "type": "string",
            "name": "tag",
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
            "description": "Hotel not found",
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
    "/hotel/room": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Get hotel rooms",
        "operationId": "get_rooms",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "hotel_id",
            "in": "query"
          },
          {
            "type": "string",
            "name": "tag",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Room"
              }
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
          "room"
        ],
        "summary": "Create room",
        "operationId": "create_room",
        "parameters": [
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Room"
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
    "/hotel/room/{room_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Get room info by id",
        "operationId": "get_room_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to return",
            "name": "room_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Room"
            }
          },
          "404": {
            "description": "room not found",
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
          "room"
        ],
        "summary": "Update room",
        "operationId": "update_room",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to change",
            "name": "room_id",
            "in": "path",
            "required": true
          },
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Room"
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
                  "type": "integer"
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
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Delete room info by id",
        "operationId": "delete_room_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to delete",
            "name": "room_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "404": {
            "description": "room not found",
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
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Delete hotel info by id",
        "operationId": "delete_hotel_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to delete",
            "name": "hotel_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "404": {
            "description": "Hotel not found",
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
        },
        "rooms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Room"
          }
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
    },
    "Room": {
      "type": "object",
      "required": [
        "hotel_id",
        "cost",
        "person_count"
      ],
      "properties": {
        "cost": {
          "description": "cost per one night",
          "type": "integer",
          "example": 10000
        },
        "hotel_id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "person_count": {
          "description": "Amount of person who can suit in number",
          "type": "integer",
          "example": 3
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "Tag": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
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
      "description": "Hotel room operations",
      "name": "room"
    },
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
    "version": "0.2.3"
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
          },
          {
            "type": "string",
            "name": "tag",
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
            "description": "Hotel not found",
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
    "/hotel/room": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Get hotel rooms",
        "operationId": "get_rooms",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "hotel_id",
            "in": "query"
          },
          {
            "type": "string",
            "name": "tag",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Room"
              }
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
          "room"
        ],
        "summary": "Create room",
        "operationId": "create_room",
        "parameters": [
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Room"
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
    "/hotel/room/{room_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Get room info by id",
        "operationId": "get_room_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to return",
            "name": "room_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Room"
            }
          },
          "404": {
            "description": "room not found",
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
          "room"
        ],
        "summary": "Update room",
        "operationId": "update_room",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to change",
            "name": "room_id",
            "in": "path",
            "required": true
          },
          {
            "name": "object",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Room"
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
                  "type": "integer"
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
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "room"
        ],
        "summary": "Delete room info by id",
        "operationId": "delete_room_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of room to delete",
            "name": "room_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "404": {
            "description": "room not found",
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
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "hotel"
        ],
        "summary": "Delete hotel info by id",
        "operationId": "delete_hotel_by_id",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of hotel to delete",
            "name": "hotel_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "404": {
            "description": "Hotel not found",
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
        },
        "rooms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Room"
          }
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
    },
    "Room": {
      "type": "object",
      "required": [
        "hotel_id",
        "cost",
        "person_count"
      ],
      "properties": {
        "cost": {
          "description": "cost per one night",
          "type": "integer",
          "example": 10000
        },
        "hotel_id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "person_count": {
          "description": "Amount of person who can suit in number",
          "type": "integer",
          "example": 3
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "Tag": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
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
      "description": "Hotel room operations",
      "name": "room"
    },
    {
      "description": "Hotel operations",
      "name": "hotel"
    }
  ]
}`))
}
