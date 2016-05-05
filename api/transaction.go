package api

import "net/http"

type Transaction interface {
	// If StartTransaction is called with a non-nil http.ResponseWriter, the
	// Transaction itself may be used in it's place.  Doing so will allow
	// future instrumentation of the response code and response headers.
	http.ResponseWriter

	// End finishes the current transaction, stopping all further
	// instrumentation.  Subsequent calls to End will have no effect. If End
	// is not called, the transaction is effectively ignored, and its data
	// will not be reported.
	End() error

	// SetName names the transaction.  Care should be taken to use a small
	// number of names:  If too many names are used, transactions will not
	// be grouped usefully.
	SetName(name string) error

	// NoticeError records an error an associates it with the Transaction. A
	// stack trace is created for the error at the point at which this
	// method is called.  If NoticeError is called multiple times in the
	// same transaction, the first five errors are recorded.  This behavior
	// is subject to potential change in the future.
	NoticeError(err error) error
}
