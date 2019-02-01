package errors

const (
	ErrAlreadyExists                     = "Already exists"
	ErrInvalidRef                        = "Invalid reference %s"
	ErrComponentNameIsEmptyMessage       = "Component name shouldn't be empty"
	ErrInvalidQuery                      = "Invalid query"
	ErrNotFound                          = "Not found"
	ErrComponentRefAlreadyExistsMessage  = "Component ref already exists"
	ErrComponentNameAlreadyExistsMessage = "Component name already exists"
	ErrClientRefAlreadyExists            = "Client ref already exists"
	ErrClientNameAlreadyExists           = "Client name already exists"
	ErrInvalidMonth                      = "Invalid month"
	ErrInvalidYear                       = "Invalid year"
	ErrTriggerUnavailable                = "Unavailable Trigger"
	ErrInvalidIncidentJSONDate           = "Field occurrence_date not found"
	ErrMongoFailuere                     = "Failed to perform operation on MongoDB"
)

func E(msg string) error {
	return &errorMsg{msg}
}

type errorMsg struct {
	msg string
}

func (e *errorMsg) Error() string {
	return e.msg
}

type ErrComponentRefAlreadyExists struct {
	Message string
}

func (e *ErrComponentRefAlreadyExists) Error() string {
	return e.Message
}

type ErrComponentNameAlreadyExists struct {
	Message string
}

func (e *ErrComponentNameAlreadyExists) Error() string {
	return e.Message
}

type ErrComponentNameIsEmpty struct {
	Message string
}

func (e *ErrComponentNameIsEmpty) Error() string {
	return e.Message
}
