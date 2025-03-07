package warehouse

import (
	"testing"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
	"github.com/stretchr/testify/require"
)

func TestHandleSchemaChange(t *testing.T) {
	inputs := []struct {
		name             string
		existingDatatype string
		currentDataType  string
		value            any

		newColumnVal any
		convError    error
	}{
		{
			name:             "should send int values if existing datatype is int, new datatype is float",
			existingDatatype: "int",
			currentDataType:  "float",
			value:            1.501,
			newColumnVal:     1,
		},
		{
			name:             "should send float values if existing datatype is float, new datatype is int",
			existingDatatype: "float",
			currentDataType:  "int",
			value:            1,
			newColumnVal:     1.0,
		},
		{
			name:             "should send string values if existing datatype is string, new datatype is boolean",
			existingDatatype: "string",
			currentDataType:  "boolean",
			value:            false,
			newColumnVal:     "false",
		},
		{
			name:             "should send string values if existing datatype is string, new datatype is int",
			existingDatatype: "string",
			currentDataType:  "int",
			value:            1,
			newColumnVal:     "1",
		},
		{
			name:             "should send string values if existing datatype is string, new datatype is float",
			existingDatatype: "string",
			currentDataType:  "float",
			value:            1.501,
			newColumnVal:     "1.501",
		},
		{
			name:             "should send string values if existing datatype is string, new datatype is datetime",
			existingDatatype: "string",
			currentDataType:  "datetime",
			value:            "2022-05-05T00:00:00.000Z",
			newColumnVal:     "2022-05-05T00:00:00.000Z",
		},
		{
			name:             "should send string values if existing datatype is string, new datatype is string",
			existingDatatype: "string",
			currentDataType:  "json",
			value:            `{"json":true}`,
			newColumnVal:     `{"json":true}`,
		},
		{
			name:             "should send json string values if existing datatype is json, new datatype is boolean",
			existingDatatype: "json",
			currentDataType:  "boolean",
			value:            false,
			newColumnVal:     "false",
		},
		{
			name:             "should send json string values if existing datatype is jso, new datatype is int",
			existingDatatype: "json",
			currentDataType:  "int",
			value:            1,
			newColumnVal:     "1",
		},
		{
			name:             "should send json string values if existing datatype is json, new datatype is float",
			existingDatatype: "json",
			currentDataType:  "float",
			value:            1.501,
			newColumnVal:     "1.501",
		},
		{
			name:             "should send json string values if existing datatype is json, new datatype is json",
			existingDatatype: "json",
			currentDataType:  "datetime",
			value:            "2022-05-05T00:00:00.000Z",
			newColumnVal:     `"2022-05-05T00:00:00.000Z"`,
		},
		{
			name:             "should send json string values if existing datatype is json, new datatype is string",
			existingDatatype: "json",
			currentDataType:  "string",
			value:            "string value",
			newColumnVal:     `"string value"`,
		},
		{
			name:             "should send json string values if existing datatype is json, new datatype is array",
			existingDatatype: "json",
			currentDataType:  "array",
			value:            []any{false, 1, "string value"},
			newColumnVal:     []any{false, 1, "string value"},
		},
		{
			name:             "existing datatype is boolean, new datatype is int",
			existingDatatype: "boolean",
			currentDataType:  "int",
			value:            1,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is boolean, new datatype is float",
			existingDatatype: "boolean",
			currentDataType:  "float",
			value:            1.501,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is boolean, new datatype is string",
			existingDatatype: "boolean",
			currentDataType:  "string",
			value:            "string value",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is boolean, new datatype is datetime",
			existingDatatype: "boolean",
			currentDataType:  "datetime",
			value:            "2022-05-05T00:00:00.000Z",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is boolean, new datatype is json",
			existingDatatype: "boolean",
			currentDataType:  "json",
			value:            `{"json":true}`,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is int, new datatype is boolean",
			existingDatatype: "int",
			currentDataType:  "boolean",
			value:            false,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is int, new datatype is string",
			existingDatatype: "int",
			currentDataType:  "string",
			value:            "string value",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is int, new datatype is datetime",
			existingDatatype: "int",
			currentDataType:  "datetime",
			value:            "2022-05-05T00:00:00.000Z",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is int, new datatype is json",
			existingDatatype: "int",
			currentDataType:  "json",
			value:            `{"json":true}`,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is int, new datatype is float",
			existingDatatype: "int",
			currentDataType:  "float",
			value:            1,
			convError:        ErrIncompatibleSchemaConversion,
		},
		{
			name:             "existing datatype is float, new datatype is boolean",
			existingDatatype: "float",
			currentDataType:  "boolean",
			value:            false,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is float, new datatype is int",
			existingDatatype: "float",
			currentDataType:  "int",
			value:            1.0,
			convError:        ErrIncompatibleSchemaConversion,
		},
		{
			name:             "existing datatype is float, new datatype is string",
			existingDatatype: "float",
			currentDataType:  "string",
			value:            "string value",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is float, new datatype is datetime",
			existingDatatype: "float",
			currentDataType:  "datetime",
			value:            "2022-05-05T00:00:00.000Z",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is float, new datatype is json",
			existingDatatype: "float",
			currentDataType:  "json",
			value:            `{"json":true}`,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is datetime, new datatype is boolean",
			existingDatatype: "datetime",
			currentDataType:  "boolean",
			value:            false,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is datetime, new datatype is string",
			existingDatatype: "datetime",
			currentDataType:  "string",
			value:            "string value",
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is datetime, new datatype is int",
			existingDatatype: "datetime",
			currentDataType:  "int",
			value:            1,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is datetime, new datatype is float",
			existingDatatype: "datetime",
			currentDataType:  "float",
			value:            1.501,
			convError:        ErrSchemaConversionNotSupported,
		},
		{
			name:             "existing datatype is datetime, new datatype is json",
			existingDatatype: "datetime",
			currentDataType:  "json",
			value:            `{"json":true}`,
			convError:        ErrSchemaConversionNotSupported,
		},
	}
	for _, ip := range inputs {
		tc := ip

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			newColumnVal, convError := HandleSchemaChange(
				model.SchemaType(tc.existingDatatype),
				model.SchemaType(tc.currentDataType),
				tc.value,
			)
			require.Equal(t, newColumnVal, tc.newColumnVal)
			require.Equal(t, convError, tc.convError)
		})
	}
}

var _ = Describe("Schema", func() {
	DescribeTable("Get table schema diff", func(tableName string, currentSchema, uploadSchema warehouseutils.SchemaT, expected warehouseutils.TableSchemaDiffT) {
		Expect(getTableSchemaDiff(tableName, currentSchema, uploadSchema)).To(Equal(expected))
	},
		Entry(nil, "test-table", warehouseutils.SchemaT{}, warehouseutils.SchemaT{}, warehouseutils.TableSchemaDiffT{
			ColumnMap:     map[string]string{},
			UpdatedSchema: map[string]string{},
		}),

		Entry(nil, "test-table", warehouseutils.SchemaT{}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value",
			},
		}, warehouseutils.TableSchemaDiffT{
			Exists:           true,
			TableToBeCreated: true,
			ColumnMap: map[string]string{
				"test-column": "test-value",
			},
			UpdatedSchema: map[string]string{
				"test-column": "test-value",
			},
		}),

		Entry(nil, "test-table", warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-1",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-2",
			},
		}, warehouseutils.TableSchemaDiffT{
			Exists:           false,
			TableToBeCreated: false,
			ColumnMap:        map[string]string{},
			UpdatedSchema: map[string]string{
				"test-column": "test-value-1",
			},
		}),

		Entry(nil, "test-table", warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column-1": "test-value-1",
				"test-column-2": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-2",
			},
		}, warehouseutils.TableSchemaDiffT{
			Exists:           true,
			TableToBeCreated: false,
			ColumnMap: map[string]string{
				"test-column": "test-value-2",
			},
			UpdatedSchema: map[string]string{
				"test-column-1": "test-value-1",
				"test-column-2": "test-value-2",
				"test-column":   "test-value-2",
			},
		}),

		Entry(nil, "test-table", warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column":   "string",
				"test-column-2": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "text",
			},
		}, warehouseutils.TableSchemaDiffT{
			Exists:           true,
			TableToBeCreated: false,
			ColumnMap:        map[string]string{},
			UpdatedSchema: map[string]string{
				"test-column-2": "test-value-2",
				"test-column":   "text",
			},
			StringColumnsToBeAlteredToText: []string{"test-column"},
		}),
	)

	DescribeTable("Merge Upload and Local Schema", func(uploadSchema, schemaInWarehousePreUpload, expected warehouseutils.SchemaT) {
		Expect(mergeUploadAndLocalSchemas(uploadSchema, schemaInWarehousePreUpload)).To(Equal(expected))
	},

		Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{}, warehouseutils.SchemaT{}),

		Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value",
			},
		}, warehouseutils.SchemaT{}),

		Entry(nil, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-1",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": {
				"test-column": "test-value-2",
			},
		}),

		Entry(nil, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column-1": "test-value-1",
				"test-column-2": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": {
				"test-column":   "test-value-2",
				"test-column-1": "test-value-1",
				"test-column-2": "test-value-2",
			},
		}),

		Entry(nil, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column":   "string",
				"test-column-2": "test-value-2",
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "text",
			},
		}, warehouseutils.SchemaT{
			"test-table": {
				"test-column-2": "test-value-2",
				"test-column":   "text",
			},
		}),
	)

	Describe("Has schema changed", func() {
		g := GinkgoT()

		Context("when skipping deep equals", func() {
			BeforeEach(func() {
				g.Setenv("RSERVER_WAREHOUSE_SKIP_DEEP_EQUAL_SCHEMAS", "true")
				Init4()
			})

			DescribeTable("Check has schema changed", func(localSchema, schemaInWarehouse warehouseutils.SchemaT, expected bool) {
				Expect(hasSchemaChanged(localSchema, schemaInWarehouse)).To(Equal(expected))
			},

				Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{}, false),

				Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "test-value",
					},
				}, false),

				Entry(nil, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "test-value-1",
					},
				}, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "test-value-2",
					},
				}, true),

				Entry(nil, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column-1": "test-value-1",
						"test-column-2": "test-value-2",
					},
				}, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "test-value-2",
					},
				}, true),

				Entry(nil, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "string",
					},
				}, warehouseutils.SchemaT{
					"test-table-1": map[string]string{
						"test-column": "text",
					},
				}, true),
			)
		})

		Context("when not skipping deep equals", func() {
			BeforeEach(func() {
				g.Setenv("RSERVER_WAREHOUSE_SKIP_DEEP_EQUAL_SCHEMAS", "false")
				Init4()
			})

			DescribeTable("Check has schema changed", func(localSchema, schemaInWarehouse warehouseutils.SchemaT, expected bool) {
				Expect(hasSchemaChanged(localSchema, schemaInWarehouse)).To(Equal(expected))
			},

				Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{}, false),

				Entry(nil, warehouseutils.SchemaT{}, warehouseutils.SchemaT{
					"test-table": map[string]string{
						"test-column": "test-value",
					},
				}, true),
			)
		})
	})

	DescribeTable("Safe name", func(warehouseType, columnName, expected string) {
		handle := SchemaHandleT{
			warehouse: warehouseutils.Warehouse{
				Type: warehouseType,
			},
		}
		Expect(handle.safeName(columnName)).To(Equal(expected))
	},
		Entry(nil, "BQ", "test-column", "test-column"),
		Entry(nil, "SNOWFLAKE", "test-column", "TEST-COLUMN"),
	)

	DescribeTable("Merge rules schema", func(warehouseType string, expected map[string]string) {
		handle := SchemaHandleT{
			warehouse: warehouseutils.Warehouse{
				Type: warehouseType,
			},
		}
		Expect(handle.getMergeRulesSchema()).To(Equal(expected))
	},
		Entry(nil, "BQ", map[string]string{
			"merge_property_1_type":  "string",
			"merge_property_1_value": "string",
			"merge_property_2_type":  "string",
			"merge_property_2_value": "string",
		}),
		Entry(nil, "SNOWFLAKE", map[string]string{
			"MERGE_PROPERTY_1_TYPE":  "string",
			"MERGE_PROPERTY_1_VALUE": "string",
			"MERGE_PROPERTY_2_TYPE":  "string",
			"MERGE_PROPERTY_2_VALUE": "string",
		}),
	)

	DescribeTable("Identities Mappings schema", func(warehouseType string, expected map[string]string) {
		handle := SchemaHandleT{
			warehouse: warehouseutils.Warehouse{
				Type: warehouseType,
			},
		}
		Expect(handle.getIdentitiesMappingsSchema()).To(Equal(expected))
	},
		Entry(nil, "BQ", map[string]string{
			"merge_property_type":  "string",
			"merge_property_value": "string",
			"rudder_id":            "string",
			"updated_at":           "datetime",
		}),
		Entry(nil, "SNOWFLAKE", map[string]string{
			"MERGE_PROPERTY_TYPE":  "string",
			"MERGE_PROPERTY_VALUE": "string",
			"RUDDER_ID":            "string",
			"UPDATED_AT":           "datetime",
		}),
	)

	DescribeTable("Discards schema", func(warehouseType string, expected map[string]string) {
		handle := SchemaHandleT{
			warehouse: warehouseutils.Warehouse{
				Type: warehouseType,
			},
		}
		Expect(handle.getDiscardsSchema()).To(Equal(expected))
	},
		Entry(nil, "BQ", map[string]string{
			"table_name":   "string",
			"row_id":       "string",
			"column_name":  "string",
			"column_value": "string",
			"received_at":  "datetime",
			"uuid_ts":      "datetime",
			"loaded_at":    "datetime",
		}),
		Entry(nil, "SNOWFLAKE", map[string]string{
			"TABLE_NAME":   "string",
			"ROW_ID":       "string",
			"COLUMN_NAME":  "string",
			"COLUMN_VALUE": "string",
			"RECEIVED_AT":  "datetime",
			"UUID_TS":      "datetime",
		}),
	)

	DescribeTable("Merge schema", func(currentSchema warehouseutils.SchemaT, schemaList []warehouseutils.SchemaT, currentMergedSchema warehouseutils.SchemaT, warehouseType string, expected warehouseutils.SchemaT) {
		Expect(mergeSchema(currentSchema, schemaList, currentMergedSchema, warehouseType)).To(Equal(expected))
	},
		Entry(nil, warehouseutils.SchemaT{}, []warehouseutils.SchemaT{}, warehouseutils.SchemaT{}, "BQ", warehouseutils.SchemaT{}),
		Entry(nil, warehouseutils.SchemaT{}, []warehouseutils.SchemaT{
			{
				"test-table": map[string]string{
					"test-column": "test-value",
				},
			},
		}, warehouseutils.SchemaT{}, "BQ", warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "test-value",
			},
		}),
		Entry(nil, warehouseutils.SchemaT{}, []warehouseutils.SchemaT{
			{
				"users": map[string]string{
					"test-column": "test-value",
				},
				"identifies": map[string]string{
					"test-column": "test-value",
				},
			},
		}, warehouseutils.SchemaT{}, "BQ", warehouseutils.SchemaT{
			"users": map[string]string{
				"test-column": "test-value",
			},
			"identifies": map[string]string{
				"test-column": "test-value",
			},
		}),
		Entry(nil, warehouseutils.SchemaT{
			"users": map[string]string{
				"test-column": "test-value",
			},
			"identifies": map[string]string{
				"test-column": "test-value",
			},
		}, []warehouseutils.SchemaT{
			{
				"users": map[string]string{
					"test-column": "test-value",
				},
				"identifies": map[string]string{
					"test-column": "test-value",
				},
			},
		}, warehouseutils.SchemaT{}, "BQ", warehouseutils.SchemaT{
			"users": map[string]string{
				"test-column": "test-value",
			},
			"identifies": map[string]string{
				"test-column": "test-value",
			},
		}),
		Entry(nil, warehouseutils.SchemaT{}, []warehouseutils.SchemaT{
			{
				"test-table": map[string]string{
					"test-column":   "string",
					"test-column-2": "test-value-2",
				},
			},
		}, warehouseutils.SchemaT{
			"test-table": map[string]string{
				"test-column": "text",
			},
		}, "BQ", warehouseutils.SchemaT{
			"test-table": {
				"test-column":   "text",
				"test-column-2": "test-value-2",
			},
		}),
	)
})
