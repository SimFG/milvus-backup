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
            "name": "wanganyang",
            "email": "wayasxxx@gmail.com"
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
        "/create": {
            "post": {
                "description": "Create a backup with the given name and collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Create backup interface",
                "parameters": [
                    {
                        "description": "CreateBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.CreateBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.BackupInfoResponse"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "Delete a backup with the given name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Delete backup interface",
                "parameters": [
                    {
                        "description": "DeleteBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.DeleteBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.DeleteBackupResponse"
                        }
                    }
                }
            }
        },
        "/get_backup": {
            "get": {
                "description": "Get the backup with the given name or id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Get backup interface",
                "parameters": [
                    {
                        "description": "GetBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.GetBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.BackupInfoResponse"
                        }
                    }
                }
            }
        },
        "/get_restore": {
            "get": {
                "description": "Get restore task state with the given id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restore"
                ],
                "summary": "Get restore interface",
                "parameters": [
                    {
                        "description": "GetRestoreStateRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.GetRestoreStateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupResponse"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "List all backups in current storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "List Backups interface",
                "parameters": [
                    {
                        "description": "ListBackupsRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.ListBackupsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.ListBackupsResponse"
                        }
                    }
                }
            }
        },
        "/restore": {
            "post": {
                "description": "Submit a request to restore the data from backup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restore"
                ],
                "summary": "Restore interface",
                "parameters": [
                    {
                        "description": "RestoreBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backuppb.BackupInfo": {
            "type": "object",
            "properties": {
                "backup_timestamp": {
                    "description": "backup timestamp",
                    "type": "integer"
                },
                "collection_backups": {
                    "description": "array of collection backup",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.CollectionBackupInfo"
                    }
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "progress": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.BackupTaskStateCode"
                }
            }
        },
        "backuppb.BackupInfoResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/backuppb.ResponseCode"
                },
                "data": {
                    "$ref": "#/definitions/backuppb.BackupInfo"
                },
                "msg": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.BackupTaskStateCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "BackupTaskStateCode_BACKUP_INITIAL",
                "BackupTaskStateCode_BACKUP_EXECUTING",
                "BackupTaskStateCode_BACKUP_SUCCESS",
                "BackupTaskStateCode_BACKUP_FAIL",
                "BackupTaskStateCode_BACKUP_TIMEOUT"
            ]
        },
        "backuppb.Binlog": {
            "type": "object",
            "properties": {
                "entries_num": {
                    "type": "integer"
                },
                "log_path": {
                    "type": "string"
                },
                "log_size": {
                    "type": "integer"
                },
                "timestamp_from": {
                    "type": "integer"
                },
                "timestamp_to": {
                    "type": "integer"
                }
            }
        },
        "backuppb.CollectionBackupInfo": {
            "type": "object",
            "properties": {
                "backup_timestamp": {
                    "type": "integer"
                },
                "collection_id": {
                    "type": "integer"
                },
                "collection_name": {
                    "type": "string"
                },
                "consistency_level": {
                    "$ref": "#/definitions/backuppb.ConsistencyLevel"
                },
                "db_name": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "partition_backups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.PartitionBackupInfo"
                    }
                },
                "progress": {
                    "type": "integer"
                },
                "schema": {
                    "$ref": "#/definitions/backuppb.CollectionSchema"
                },
                "shards_num": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.BackupTaskStateCode"
                }
            }
        },
        "backuppb.CollectionSchema": {
            "type": "object",
            "properties": {
                "autoID": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldSchema"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "backuppb.ConsistencyLevel": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "ConsistencyLevel_Strong",
                "ConsistencyLevel_Session",
                "ConsistencyLevel_Bounded",
                "ConsistencyLevel_Eventually",
                "ConsistencyLevel_Customized"
            ]
        },
        "backuppb.CreateBackupRequest": {
            "type": "object",
            "properties": {
                "async": {
                    "type": "boolean"
                },
                "backup_name": {
                    "description": "backup name, if not set, will generate one",
                    "type": "string"
                },
                "collection_names": {
                    "description": "collection names to backup, empty to backup all",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.DataType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                10,
                11,
                20,
                21,
                100,
                101
            ],
            "x-enum-varnames": [
                "DataType_None",
                "DataType_Bool",
                "DataType_Int8",
                "DataType_Int16",
                "DataType_Int32",
                "DataType_Int64",
                "DataType_Float",
                "DataType_Double",
                "DataType_String",
                "DataType_VarChar",
                "DataType_BinaryVector",
                "DataType_FloatVector"
            ]
        },
        "backuppb.DeleteBackupRequest": {
            "type": "object",
            "properties": {
                "backup_name": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.DeleteBackupResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/backuppb.ResponseCode"
                },
                "msg": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.FieldBinlog": {
            "type": "object",
            "properties": {
                "binlogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.Binlog"
                    }
                },
                "fieldID": {
                    "type": "integer"
                }
            }
        },
        "backuppb.FieldSchema": {
            "type": "object",
            "properties": {
                "autoID": {
                    "type": "boolean"
                },
                "data_type": {
                    "$ref": "#/definitions/backuppb.DataType"
                },
                "description": {
                    "type": "string"
                },
                "fieldID": {
                    "type": "integer"
                },
                "index_params": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.KeyValuePair"
                    }
                },
                "is_primary_key": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/backuppb.FieldState"
                },
                "type_params": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.KeyValuePair"
                    }
                }
            }
        },
        "backuppb.FieldState": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "FieldState_FieldCreated",
                "FieldState_FieldCreating",
                "FieldState_FieldDropping",
                "FieldState_FieldDropped"
            ]
        },
        "backuppb.GetBackupRequest": {
            "type": "object",
            "properties": {
                "backup_id": {
                    "type": "string"
                },
                "backup_name": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.GetRestoreStateRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.KeyValuePair": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "backuppb.ListBackupsRequest": {
            "type": "object",
            "properties": {
                "collection_name": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.ListBackupsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/backuppb.ResponseCode"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.BackupInfo"
                    }
                },
                "msg": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.PartitionBackupInfo": {
            "type": "object",
            "properties": {
                "collection_id": {
                    "type": "integer"
                },
                "partition_id": {
                    "type": "integer"
                },
                "partition_name": {
                    "type": "string"
                },
                "segment_backups": {
                    "description": "array of segment backup",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.SegmentBackupInfo"
                    }
                }
            }
        },
        "backuppb.ResponseCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                400,
                404
            ],
            "x-enum-varnames": [
                "ResponseCode_Success",
                "ResponseCode_Not_Support",
                "ResponseCode_No_Permission",
                "ResponseCode_Bad_Request",
                "ResponseCode_Parameter_Error",
                "ResponseCode_Request_Object_Not_Found"
            ]
        },
        "backuppb.RestoreBackupRequest": {
            "type": "object",
            "properties": {
                "async": {
                    "type": "boolean"
                },
                "backup_name": {
                    "type": "string"
                },
                "collection_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "collection_renames": {
                    "description": "2, give a map to rename the collections, if not given, use the original name",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "collection_suffix": {
                    "description": "Support two ways to rename the collections while recover\n1, set a suffix",
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.RestoreBackupResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/backuppb.ResponseCode"
                },
                "data": {
                    "$ref": "#/definitions/backuppb.RestoreBackupTask"
                },
                "msg": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "backuppb.RestoreBackupTask": {
            "type": "object",
            "properties": {
                "collection_restore_tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.RestoreCollectionTask"
                    }
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "progress": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                }
            }
        },
        "backuppb.RestoreCollectionTask": {
            "type": "object",
            "properties": {
                "coll_backup": {
                    "$ref": "#/definitions/backuppb.CollectionBackupInfo"
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "partition_restore_tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.RestorePartitionTask"
                    }
                },
                "progress": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                },
                "target_collection_name": {
                    "type": "string"
                }
            }
        },
        "backuppb.RestorePartitionTask": {
            "type": "object",
            "properties": {
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "part_backup": {
                    "$ref": "#/definitions/backuppb.PartitionBackupInfo"
                },
                "progress": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                }
            }
        },
        "backuppb.RestoreTaskStateCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "RestoreTaskStateCode_INITIAL",
                "RestoreTaskStateCode_EXECUTING",
                "RestoreTaskStateCode_SUCCESS",
                "RestoreTaskStateCode_FAIL",
                "RestoreTaskStateCode_TIMEOUT"
            ]
        },
        "backuppb.SegmentBackupInfo": {
            "type": "object",
            "properties": {
                "binlogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                },
                "collection_id": {
                    "type": "integer"
                },
                "deltalogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                },
                "num_of_rows": {
                    "type": "integer"
                },
                "partition_id": {
                    "type": "integer"
                },
                "segment_id": {
                    "type": "integer"
                },
                "statslogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Milvus Backup Service",
	Description:      "A data backup & restore tool for Milvus",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}