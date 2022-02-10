package utility


// Response
type Response struct {
	Message string
	Code string
	Status bool
}

// ResponseObj
type ResponseObj struct {
	Response
	Data interface{}
}

// ValidationResponseObj
type ValidationResponseObj struct {
	Response
	ValidationMsg string
}

// Success returns success response with data
func (r Response) Success(code, message string, data interface{}) ResponseObj {
	res:=Response{}
	res.Status= true
	res.Message= message
	res.Code= code

	return ResponseObj{
		Response: res,
		Data: data,
	}
}

//PlainSuccess returns success response without data
func (r Response) PlainSuccess(code, message string) Response {
	return Response{
		Status: true,
		Message: message,
		Code: code,
	}
}

// Error returns error response
func (r Response) Error(code, message string) Response {
	return Response{
		Status: false,
		Message: message,
		Code: code,
	}
}

// ValidationError return validation error
func (r Response) ValidationError(code, message, error string) ValidationResponseObj {
	res:=Response{
		Status: false,
		Message: message,
		Code: code,
	}
	return ValidationResponseObj{
		Response:   res,
		ValidationMsg: error,
	}
}

func NewResponse() Response {
	return Response{}
}

