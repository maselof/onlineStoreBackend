package constants

const (
	ErrorGetParams = iota + 1
	ErrorUnmarshalRequest
	ErrorEmptyCast
)

var Errors = map[int]string{
	ErrorGetParams:        "bad GET params",
	ErrorUnmarshalRequest: "cannot unmarshal request",
	ErrorEmptyCast:        "Cast is empty",
}
