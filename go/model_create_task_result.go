// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 */

package openapi




type CreateTaskResult struct {

	Code int64 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	TaskId string `json:"taskId,omitempty"`
}

// AssertCreateTaskResultRequired checks if the required fields are not zero-ed
func AssertCreateTaskResultRequired(obj CreateTaskResult) error {
	return nil
}

// AssertCreateTaskResultConstraints checks if the values respects the defined constraints
func AssertCreateTaskResultConstraints(obj CreateTaskResult) error {
	return nil
}
