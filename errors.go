//The MIT License (MIT)
//Copyright (c) 2014 Atomos
//

/*
Package errors implements the representation of enhanced error conditions with
internal and public messages customized to each instance of an error. These
enhanced error conditions are meant to be used in systems where error messages are
displayed to untrusted third parties.

The private error message allows for a detailed description that could potentially
reveal the internal structure of the system. The public message is meant to be
viewed by the general public so care must be taken to ensure that the internal
structure of the system is not revealed in the public message.

Each error can have an error code associated with it to allow for a simple way of
determining the class of an error condition in cases where each instance of
an error can potentially have a different error message.

A nil value represents no error.

This package can be a full replacement for the default errors package.
*/
package errors

var (
	DefaultCode = 0
)

// An Error represents an error condition
type Error interface {
	// An integer value representing the error class. This value can be used to
	// create error conditions where the message strings can be customized for each
	// specific error condition while allowing for the error class to be determined
	// without searching through the error strings.
	Code() int

	Private() string

	Public() string
}

type genericError struct {
	code    int
	private string
	public  string
}

func (err genericError) Code() int       { return err.code }
func (err genericError) Private() string { return err.private }
func (err genericError) Public() string  { return err.public }

// Returns the Private string. Used to satisfy the builtin Error interface.
func (err genericError) Error() string { return err.private }

// Returns a genericError which implements the Error interface with Code set to
// DefaultCode.
func NewErrorDefaultCode(private, public string) Error {
	return genericError{
		code:    DefaultCode,
		private: private,
		public:  public,
	}
}

// Returns a genericError which implements the Error interface with Code set to
// the given value.
func NewError(private, public string, code int) Error {
	return genericError{
		code:    code,
		private: private,
		public:  public,
	}
}

// New returns an Error with Code set to DefaultCode and Private and Public messages
// set to the given text.
//
// New is used to support code that uses the function signature set forth by the
// default errors package.
func New(text string) Error {
	return NewErrorDefaultCode(text, text)
}
