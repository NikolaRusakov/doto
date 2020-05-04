/*
 * Do-to
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger


type Goal struct {

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Id string`json:"id,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	Section *Goal_sections `json:"section,omitempty"`

	Tasks *Task `json:"tasks,omitempty"`
}
