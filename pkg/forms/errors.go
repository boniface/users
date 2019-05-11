package forms

// Define a new errors type, which we will use to hold the validation error
// messages for forms. The name of the form field will be used as the key in
// this map.
type errors map[string][]string

// Implement an Add() method to add error messages for a given field to the map.
func (fieldErrors errors) Add(field, message string) {
	fieldErrors[field] = append(fieldErrors[field], message)
}

// Implement a Get() method to retrieve the first error message for a given
// field from the map.
func (fieldErrors errors) Get(field string) string {
	fieldErrorSlice := fieldErrors[field]
	if len(fieldErrorSlice) == 0 {
		return ""
	}
	return fieldErrorSlice[0]
}
