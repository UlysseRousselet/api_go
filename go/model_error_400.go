/*
 * SMS - OpenAPI 3.0
 *
 * A simple api for sending sms.
 *
 * API version: 1.0.1
 * Contact: ulysse.rousselet@callr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Error400 struct {

	Id int64 `json:"id,omitempty"`

	Number int64 `json:"number,omitempty"`

	Message string `json:"message,omitempty"`
}
