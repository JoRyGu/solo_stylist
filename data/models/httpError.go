package models

// HttpError models the standard response body when an error is encountered in the program.
// Any user errors will be present in the Fields property, otherwise Fields will be nil.
type HttpError struct {
	StatusCode    int               `json:"statusCode"`
	StatusMessage string            `json:"statusMessage"`
	Message       string            `json:"message"`
	Fields        []*HttpFieldError `json:"fields,omitempty"`
}

// HttpFieldError models a user error in a request body field.
type HttpFieldError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
