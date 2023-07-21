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

import (
	"net/http"
)

func SendSms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}