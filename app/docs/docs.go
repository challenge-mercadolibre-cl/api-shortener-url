// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.0",
  "info": {
    "title": "Shortener Services",
    "version": "1.0"
  },
  "servers": [
    {
      "url": "http://34.123.67.196",
      "variables": {}
    }
  ],
  "paths": {
    "/url/shortener": {
      "post": {
        "tags": ["Shortener"],
        "summary": "Create Shortener",
        "operationId": "CreateShortener",
        "parameters": [],
        "requestBody": {
          "description": "",
          "content": {
            "application/json": {
              "schema": {
                "allOf": [
                  {
                    "$ref": "#/components/schemas/CreateShortenerRequest"
                  },
                  {
                    "example": {
                      "user_id": "3",
                      "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
                    }
                  }
                ]
              },
              "example": {
                "user_id": "3",
                "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
              }
            }
          },
          "required": true
        },
        "responses": {
          "201": {
            "description": "Created",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/201CreatedCreateShortener"
                    },
                    {
                      "example": {
                        "status": 201,
                        "data": "dfb22053"
                      }
                    }
                  ]
                },
                "example": {
                  "status": 201,
                  "data": "dfb22053"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/400BadRequestCreateShortener"
                    },
                    {
                      "example": {
                        "error": "user id cant be empty"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "user id cant be empty"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/500InternalServerErrorCreateShortener"
                    },
                    {
                      "example": {
                        "error": "Internal Server Error"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "Internal Server Error"
                }
              }
            }
          }
        },
        "deprecated": false
      }
    },
    "/url/shortener/{id}": {
      "put": {
        "tags": ["Shortener"],
        "summary": "Update Shortener",
        "operationId": "UpdateShortener",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "",
            "required": true,
            "style": "simple",
            "schema": {
              "type": "string",
              "example": "null"
            }
          }
        ],
        "requestBody": {
          "description": "",
          "content": {
            "application/json": {
              "schema": {
                "allOf": [
                  {
                    "$ref": "#/components/schemas/UpdateShortenerRequest"
                  },
                  {
                    "example": {
                      "user_id": "3",
                      "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
                    }
                  }
                ]
              },
              "example": {
                "user_id": "3",
                "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
              }
            }
          },
          "required": true
        },
        "responses": {
          "201": {
            "description": "Created",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/201CreatedUpdateShortener"
                    },
                    {
                      "example": {
                        "status": 201,
                        "data": "dfb22053"
                      }
                    }
                  ]
                },
                "example": {
                  "status": 201,
                  "data": "dfb22053"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/400BadRequestUpdateShortener"
                    },
                    {
                      "example": {
                        "error": "user id cant be empty"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "user id cant be empty"
                }
              }
            }
          },
          "404": {
            "description": "Not Found",
            "headers": {
              "Vary": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Origin"
                  }
                }
              },
              "Date": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Wed, 19 Oct 2022 15:38:56 GMT"
                  }
                }
              },
              "Content-Length": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "33"
                  }
                }
              }
            },
            "content": {
              "application/json; charset=UTF-8": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/404NotfoundUpdateShortener"
                    },
                    {
                      "example": {
                        "error": "url id not found : 1"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "url id not found : 1"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/500InternalServerErrorUpdateShortener"
                    },
                    {
                      "example": {
                        "error": "Internal Server Error"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "Internal Server Error"
                }
              }
            }
          }
        },
        "deprecated": false
      },
      "get": {
        "tags": ["Shortener"],
        "summary": "Find One",
        "operationId": "FindOne",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "",
            "required": true,
            "style": "simple",
            "schema": {
              "type": "string",
              "example": "null"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Vary": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Origin"
                  }
                }
              },
              "Date": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Wed, 19 Oct 2022 15:37:42 GMT"
                  }
                }
              },
              "Content-Length": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "97"
                  }
                }
              }
            },
            "content": {
              "application/json; charset=UTF-8": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/200SuccessFindOneShortener"
                    },
                    {
                      "example": {
                        "status": 200,
                        "data": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
                      }
                    }
                  ]
                },
                "example": {
                  "status": 200,
                  "data": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
                }
              }
            }
          },
          "404": {
            "description": "Not Found",
            "headers": {
              "Vary": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Origin"
                  }
                }
              },
              "Date": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Wed, 19 Oct 2022 15:38:12 GMT"
                  }
                }
              },
              "Content-Length": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "16"
                  }
                }
              }
            },
            "content": {
              "text/plain; charset=UTF-8": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/404NotFoundFindOneShortener"
                    },
                    {
                      "example": {
                        "error": "url id not found : 1"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "url id not found : 1"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {},
            "content": {
              "application/json": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/500InternalServerErrorFindOneShortener"
                    },
                    {
                      "example": {
                        "error": "Internal Server Error"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "Internal Server Error"
                }
              }
            }
          }
        },
        "deprecated": false
      }
    },
    "/health": {
      "get": {
        "tags": ["Health"],
        "summary": "Health",
        "operationId": "Health",
        "parameters": [],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Vary": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Origin"
                  }
                }
              },
              "Date": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Wed, 19 Oct 2022 15:41:10 GMT"
                  }
                }
              },
              "Content-Length": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "89"
                  }
                }
              }
            },
            "content": {
              "application/json; charset=UTF-8": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/200OKHealth"
                    },
                    {
                      "example": {
                        "status": "UP",
                        "service_name": "",
                        "version": "",
                        "uptime": "7m8.78652375s",
                        "environment": ""
                      }
                    }
                  ]
                },
                "example": {
                  "status": "UP",
                  "service_name": "",
                  "version": "",
                  "uptime": "7m8.78652375s",
                  "environment": ""
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {
              "Vary": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Origin"
                  }
                }
              },
              "Date": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "Wed, 19 Oct 2022 15:39:48 GMT"
                  }
                }
              },
              "Content-Length": {
                "content": {
                  "text/plain": {
                    "schema": {
                      "type": "string"
                    },
                    "example": "91"
                  }
                }
              }
            },
            "content": {
              "application/json; charset=UTF-8": {
                "schema": {
                  "allOf": [
                    {
                      "$ref": "#/components/schemas/500InternalServerErrorHealth"
                    },
                    {
                      "example": {
                        "error": "Internal Server Error"
                      }
                    }
                  ]
                },
                "example": {
                  "error": "Internal Server Error"
                }
              }
            }
          }
        },
        "deprecated": false
      }
    }
  },
  "components": {
    "schemas": {
      "CreateShortenerRequest": {
        "title": "CreateShortenerRequest",
        "required": ["user_id", "url"],
        "type": "object",
        "properties": {
          "user_id": {
            "type": "string"
          },
          "url": {
            "type": "string"
          }
        },
        "example": {
          "user_id": "3",
          "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
        }
      },
      "201CreatedCreateShortener": {
        "title": "201CreatedCreateShortener",
        "required": ["status", "data"],
        "type": "object",
        "properties": {
          "status": {
            "type": "integer",
            "format": "int32"
          },
          "data": {
            "type": "string"
          }
        },
        "example": {
          "status": 201,
          "data": "dfb22053"
        }
      },
      "400BadRequestCreateShortener": {
        "title": "400BadRequestCreateShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "user id cant be empty"
        }
      },
      "500InternalServerErrorCreateShortener": {
        "title": "500InternalServerErrorCreateShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "Internal Server Error"
        }
      },
      "UpdateShortenerRequest": {
        "title": "UpdateShortenerRequest",
        "required": ["user_id", "url"],
        "type": "object",
        "properties": {
          "user_id": {
            "type": "string"
          },
          "url": {
            "type": "string"
          }
        },
        "example": {
          "user_id": "3",
          "url": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
        }
      },
      "201CreatedUpdateShortener": {
        "title": "201CreatedUpdateShortener",
        "required": ["status", "data"],
        "type": "object",
        "properties": {
          "status": {
            "type": "integer",
            "format": "int32"
          },
          "data": {
            "type": "string"
          }
        },
        "example": {
          "status": 201,
          "data": "dfb22053"
        }
      },
      "400BadRequestUpdateShortener": {
        "title": "400BadRequestUpdateShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "user id cant be empty"
        }
      },
      "404NotfoundUpdateShortener": {
        "title": "404NotfoundUpdateShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "url id not found : 1"
        }
      },
      "500InternalServerErrorUpdateShortener": {
        "title": "500InternalServerErrorUpdateShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "Internal Server Error"
        }
      },
      "200SuccessFindOneShortener": {
        "title": "200SuccessFindOneShortener",
        "required": ["status", "data"],
        "type": "object",
        "properties": {
          "status": {
            "type": "integer",
            "format": "int32"
          },
          "data": {
            "type": "string"
          }
        },
        "example": {
          "status": 200,
          "data": "https://www.youtube.com/watch?v=p38WgakuYDo&list=RDMM&index=12"
        }
      },
      "404NotFoundFindOneShortener": {
        "title": "404NotFoundFindOneShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "url id not found : 1"
        }
      },
      "500InternalServerErrorFindOneShortener": {
        "title": "500InternalServerErrorFindOneShortener",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "Internal Server Error"
        }
      },
      "500InternalServerErrorHealth": {
        "title": "500InternalServerErrorHealth",
        "required": ["error"],
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "Internal Server Error"
        }
      },
      "200OKHealth": {
        "title": "200OKHealth",
        "required": [
          "status",
          "service_name",
          "version",
          "uptime",
          "environment"
        ],
        "type": "object",
        "properties": {
          "status": {
            "type": "string"
          },
          "service_name": {
            "type": "string"
          },
          "version": {
            "type": "string"
          },
          "uptime": {
            "type": "string"
          },
          "environment": {
            "type": "string"
          }
        },
        "example": {
          "status": "UP",
          "service_name": "",
          "version": "",
          "uptime": "7m8.78652375s",
          "environment": ""
        }
      }
    }
  },
  "tags": [
    {
      "name": "Shortener",
      "description": "Shortener Services"
    },
    {
      "name": "Health",
      "description": "Health Services"
    }
  ]
}
`

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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
