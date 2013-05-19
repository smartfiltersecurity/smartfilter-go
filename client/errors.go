package client

type SmartFilterBadInputParameter struct{}

func (e *SmartFilterBadInputParameter) Error() string {
	return "Bad input parameter"
}

type SmartFilterBadAPIKey struct{}

func (e *SmartFilterBadAPIKey) Error() string {
	return "Bad API key"
}

type SmartFilterRequestTooLarge struct{}

func (e *SmartFilterRequestTooLarge) Error() string {
	return "Request too large"
}

type SmartFilterInternalError struct{}

func (e *SmartFilterInternalError) Error() string {
	return "Internal error"
}

type SmartFilterAccountQuotaExceeded struct{}

func (e *SmartFilterAccountQuotaExceeded) Error() string {
	return "Account quota exceeded"
}
