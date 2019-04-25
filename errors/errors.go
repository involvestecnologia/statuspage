package errors

const (
	//ErrAlreadyExistsMessage is the resource already exists default message
	ErrAlreadyExistsMessage = "Already exists"

	//ErrInvalidRefMessage is the invalid reference default template message
	ErrInvalidRefMessage = "Invalid reference %s"

	//ErrComponentNameIsEmptyMessage is the empty component name default message
	ErrComponentNameIsEmptyMessage = "Component name shouldn't be empty"

	//ErrInvalidQueryMessage is the invalid query attempt default message
	ErrInvalidQueryMessage = "Invalid query"

	//ErrNotFoundMessage is the resource not found default message
	ErrNotFoundMessage = "Not found"

	//ErrComponentRefAlreadyExistsMessage is the Component reference already in use default message
	ErrComponentRefAlreadyExistsMessage = "Component ref already exists"

	//ErrComponentNameAlreadyExistsMessage is the Component name already in use default message
	ErrComponentNameAlreadyExistsMessage = "Component name already exists"

	//ErrClientRefAlreadyExistsMessage is the Client reference already in use default message
	ErrClientRefAlreadyExistsMessage = "Client ref already exists"

	//ErrClientNameAlreadyExistsMessage is the Client name already in use default message
	ErrClientNameAlreadyExistsMessage = "Client name already exists"

	//ErrInvalidDateMessage is the invalid date value default message
	ErrInvalidDateMessage = "Invalid date"

	//ErrTriggerUnavailableMessage is the trigger unavailable default message
	ErrTriggerUnavailableMessage = "Unavailable Trigger"

	//ErrInvalidIncidentJSONDateMessage is the invalid occourence_date value or format default message
	ErrInvalidIncidentJSONDateMessage = "Field occurrence_date not found"

	//ErrMongoFailureMessage is the mongo panic failure default message
	ErrMongoFailureMessage = "Failed to perform operation on MongoDB"

	//ErrIncidentStatusIgnoredMessage is the ignored Incident creation default message
	ErrIncidentStatusIgnoredMessage = "Status didn't close last incident or escaled it's status"
)

//ErrAlreadyExists is a error type throwed when the resource already exists
type ErrAlreadyExists struct {
	Message string
}

//ErrInvalidRef is a error type throwed when a reference is invalid
type ErrInvalidRef struct {
	Message string
}

//ErrComponentNameIsEmpty is a error type throwed when a name attribute of a Component is empty
type ErrComponentNameIsEmpty struct {
	Message string
}

//ErrInvalidQuery is a error type throwed when a query attempt is invalid
type ErrInvalidQuery struct {
	Message string
}

//ErrNotFound is a error type throwed when a resource was not found
type ErrNotFound struct {
	Message string
}

//ErrComponentRefAlreadyExists is a error type throwed when a Component reference is already in use
type ErrComponentRefAlreadyExists struct {
	Message string
}

//ErrComponentNameAlreadyExists is a error type throwed when a Component name is already in use
type ErrComponentNameAlreadyExists struct {
	Message string
}

//ErrClientRefAlreadyExists is a error type throwed when a Client reference is already in use
type ErrClientRefAlreadyExists struct {
	Message string
}

//ErrClientNameAlreadyExists is a error type throwed when a Client name is already in use
type ErrClientNameAlreadyExists struct {
	Message string
}

//ErrInvalidDate is a error type throwed when a invalid date value is provided
type ErrInvalidDate struct {
	Message string
}

//ErrTriggerUnavailable is a error type throwed when a trigger type is no available
type ErrTriggerUnavailable struct {
	Message string
}

//ErrInvalidIncidentJSONDate is a error type throwed when the date value of a Incident is in a invalid format
type ErrInvalidIncidentJSONDate struct {
	Message string
}

//ErrMongoFailure is a error type throwed when a panic occoured when connecting to MongoDB
type ErrMongoFailure struct {
	Message string
}

//ErrIncidentStatusIgnored is a error type throwed when a Incident status does not resolve last neither increase the criticity
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
func (e *ErrInvalidDate) Error() string {
	return e.Message
}
func (e *ErrTriggerUnavailable) Error() string {
	return e.Message
}
func (e *ErrInvalidIncidentJSONDate) Error() string {
	return e.Message
}
func (e *ErrMongoFailure) Error() string {
	return e.Message
}

func (e *ErrIncidentStatusIgnored) Error() string {
	return e.Message
}
