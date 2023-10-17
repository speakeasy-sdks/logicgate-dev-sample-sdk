# FieldRecordAPIOut

The fields and values of the record


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Global`                                                                         | **bool*                                                                          | :heavy_minus_sign:                                                               | Whether the field is global                                                      | false                                                                            |
| `ID`                                                                             | **string*                                                                        | :heavy_minus_sign:                                                               | The unique ID of this Risk Cloud resource                                        | a1b2c3d4                                                                         |
| `Label`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | The label of the field as shown on the record                                    | Enter the Risk Severity based on your assessment                                 |
| `Name`                                                                           | **string*                                                                        | :heavy_minus_sign:                                                               | The name of the field                                                            | Risk Severity                                                                    |
| `Object`                                                                         | **string*                                                                        | :heavy_minus_sign:                                                               | Identifies the type of object this data represents                               | field                                                                            |
| `Type`                                                                           | [*FieldRecordAPIOutFieldType](../../models/shared/fieldrecordapioutfieldtype.md) | :heavy_minus_sign:                                                               | The type of the field                                                            | SELECT                                                                           |
| `ValueType`                                                                      | [*FieldRecordAPIOutValueType](../../models/shared/fieldrecordapioutvaluetype.md) | :heavy_minus_sign:                                                               | The type of the field value                                                      | OPTION                                                                           |
| `Values`                                                                         | [][ValuePropertyAPIOut](../../models/shared/valuepropertyapiout.md)              | :heavy_minus_sign:                                                               | The values of the record field                                                   |                                                                                  |