/*
 * Do-to
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Task struct {

	Id int64 `json:"id,omitempty"`

	Section *Goal_sections `json:"section,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	Estimation string `json:"estimation,omitempty"`

	Subtask *Task `json:"subtask,omitempty"`
}