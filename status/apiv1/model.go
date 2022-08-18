package apiv1

// Status struct for Status
type Status struct {
	Healthy   bool                              `json:"healthy,omitempty"`
	Message   string                            `json:"message,omitempty"`
	Error     *Throwable                        `json:"error,omitempty"`
	Details   map[string]map[string]interface{} `json:"details,omitempty"`
	Time      int64                             `json:"time,omitempty"`
	Duration  int64                             `json:"duration,omitempty"`
	Timestamp string                            `json:"timestamp,omitempty"`
}

// Throwable struct for Throwable
type Throwable struct {
	Cause            *Throwable           `json:"cause,omitempty"`
	StackTrace       []*StackTraceElement `json:"stackTrace,omitempty"`
	Message          string               `json:"message,omitempty"`
	LocalizedMessage string               `json:"localizedMessage,omitempty"`
	Suppressed       []*Throwable         `json:"suppressed,omitempty"`
}

// StackTraceElement struct for StackTraceElement
type StackTraceElement struct {
	MethodName   string `json:"methodName,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	LineNumber   int32  `json:"lineNumber,omitempty"`
	ClassName    string `json:"className,omitempty"`
	NativeMethod bool   `json:"nativeMethod,omitempty"`
}
