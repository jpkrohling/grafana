// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// GraphiteDataQuery defines model for GraphiteDataQuery.
type GraphiteDataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource      *interface{} `json:"datasource,omitempty"`
	FromAnnotations *bool        `json:"fromAnnotations,omitempty"`

	// true if query is disabled (ie should not be returned to the dashboard)
	Hide *bool `json:"hide,omitempty"`

	// Unique, guid like, string used in explore mode
	Key *string `json:"key,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A - Z
	RefId      string    `json:"refId"`
	Tags       *[]string `json:"tags,omitempty"`
	Target     *string   `json:"target,omitempty"`
	TextEditor *bool     `json:"textEditor,omitempty"`
}
