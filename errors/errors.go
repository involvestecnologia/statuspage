package errors

const (
	ErrAlreadyExistsMessage              = "Already exists"
	ErrInvalidRefMessage                 = "Invalid reference %s"
	ErrComponentNameIsEmptyMessage       = "Component name shouldn't be empty"
	ErrInvalidQueryMessage               = "Invalid query"
	ErrNotFoundMessage                   = "Not found"
	ErrComponentRefAlreadyExistsMessage  = "Component ref already exists"
	ErrComponentNameAlreadyExistsMessage = "Component name already exists"
	ErrClientRefAlreadyExistsMessage     = "Client ref already exists"
	ErrClientNameAlreadyExistsMessage    = "Client name already exists"
	ErrInvalidMonthMessage               = "Invalid month"
	ErrInvalidYearMessage                = "Invalid year"
	ErrTriggerUnavailableMessage         = "Unavailable Trigger"
	ErrInvalidIncidentJSONDateMessage    = "Field occurrence_date not found"
	ErrMongoFailuereMessage              = "Failed to perform operation on MongoDB"
	ErrIncidentStatusIgnoredMessage      = "Status didn't close last incident or escaled it's status"
)

type ErrAlreadyExists struct {
	Message string
}
type ErrInvalidRef struct {
	Message string
}
type ErrComponentNameIsEmpty struct {
	Message string
}
type ErrInvalidQuery struct {
	Message string
}
type ErrNotFound struct {
	Message string
}
type ErrComponentRefAlreadyExists struct {
	Message string
}
type ErrComponentNameAlreadyExists struct {
	Message string
}
type ErrClientRefAlreadyExists struct {
	Message string
}
type ErrClientNameAlreadyExists struct {
	Message string
}
type ErrInvalidMonth struct {
	Message string
}
type ErrInvalidYear struct {
	Message string
}
type ErrTriggerUnavailable struct {
	Message string
}
type ErrInvalidIncidentJSONDate struct {
	Message string
}
type ErrMongoFailuere struct {
	Message string
}
type ErrIncidentStatusIgnored struct {
	Message string
}

func (e *ErrAlreadyExists) Error() string {
	return e.Message
}
func (e *ErrInvalidRef) Error() string {
	return e.Message
}
func (e *ErrComponentNameIsEmpty) Error() string {
	return e.Message
}
func (e *ErrInvalidQuery) Error() string {
	return e.Message
}
func (e *ErrNotFound) Error() string {
	return e.Message
}
func (e *ErrComponentRefAlreadyExists) Error() string {
	return e.Message
}
func (e *ErrComponentNameAlreadyExists) Error() string {
	return e.Message
}
func (e *ErrClientRefAlreadyExists) Error() string {
	return e.Message
}
func (e *ErrClientNameAlreadyExists) Error() string {
	return e.Message
}
func (e *ErrInvalidMonth) Error() string {
	return e.Message
}
func (e *ErrInvalidYear) Error() string {
	return e.Message
}
func (e *ErrTriggerUnavailable) Error() string {
	return e.Message
}
func (e *ErrInvalidIncidentJSONDate) Error() string {
	return e.Message
}
func (e *ErrMongoFailuere) Error() string {
	return e.Message
}
func (e *ErrIncidentStatusIgnored) Error() string {
	return e.Message
}
