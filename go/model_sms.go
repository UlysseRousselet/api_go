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

type Sms struct {

	Id int64 `json:"id,omitempty"`
	// Client's phone number
	To int64 `json:"to,omitempty"`
	// Group registered to the Database to send sms
	Group string `json:"group,omitempty"`
	// Name of the sender that have been registered to the customer portail
	From string `json:"from"`
	// Message to send to the client
	Message string `json:"message"`
	// Print the datetime
	Date string `json:"date,omitempty"`
}