errors
=======

[godoc.org/github.com/atomosio/errors]

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

