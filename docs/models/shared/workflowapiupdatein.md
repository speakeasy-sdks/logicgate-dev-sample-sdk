# WorkflowAPIUpdateIn

Workflow (Update)


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Name`                                                                                | **string*                                                                             | :heavy_minus_sign:                                                                    | The name of the workflow                                                              | Risk Assessments                                                                      |
| `RecordPrefix`                                                                        | **string*                                                                             | :heavy_minus_sign:                                                                    | The prefix to be used in the name of every record created from this workflow          | Assessment                                                                            |
| `Xpos`                                                                                | **int*                                                                                | :heavy_minus_sign:                                                                    | The x-coordinate of the workflow in the application builder (must not be less than 0) | 20                                                                                    |
| `Ypos`                                                                                | **int*                                                                                | :heavy_minus_sign:                                                                    | The y-coordinate of the workflow in the application builder (must not be less than 0) | 20                                                                                    |