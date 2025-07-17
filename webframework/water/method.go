package water

type Method string

const (
	MethodGET     Method = "GET"
	MethodPOST    Method = "POST"
	MethodPUT     Method = "PUT"
	MethodPATCH   Method = "PATCH"
	MethodDELETE  Method = "DELETE"
	MethodHEAD    Method = "HEAD"
	MethodOPTIONS Method = "OPTIONS"
	MethodTRACE   Method = "TRACE"
	MethodCONNECT Method = "CONNECT"
)
