package utilities

type Result struct {
	Success bool
	Message string
}

type DataResult struct {
	Result
	Data interface{}
}
type ResultOfSuccessData struct {
	DataResult
	Success bool `default:"true"`
}

func SuccessDataResult(msg string, d interface{}) *ResultOfSuccessData {
	return &ResultOfSuccessData{
		DataResult: DataResult{
			Result: Result{
				Message: msg,
			},
			Data: d,
		},
		Success: true,
	}
}

type ResultOfErrorData struct {
	DataResult
	Success bool
}

func ErrorDataResult(msg string, d interface{}) *ResultOfErrorData {
	return &ResultOfErrorData{
		DataResult: DataResult{
			Result: Result{
				Message: msg,
			},
			Data: d,
		},
		Success: false,
	}
}

type ResultSuccess struct {
	Success bool
	Message string
}

func SuccessResult(msg string) *ResultSuccess {
	return &ResultSuccess{
		Success: true,
		Message: msg,
	}
}

type ResultError struct {
	Success bool
	Message string
}

func ErrorResult(msg string) *ResultError {
	return &ResultError{
		Success: false,
		Message: msg,
	}
}
