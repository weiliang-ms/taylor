package errors

import (
	"fmt"
	"runtime"
)

/*
	主要为封装错误
*/

// 定义错误类型
type errNotFound struct{ error }

type errConflict struct{ error }

type errUnavailable struct{ error }

type errNotModified struct{ error }

type errSystem struct{ error }

type errForbidden struct{ error }

type errNotImplemented struct{ error }

type errInvalidParameter struct{ error }

type errUnauthorized struct{ error }

type errUnknown struct{ error }

// causer接口，需实现Cause()方法
type causer interface {
	Cause() error
}

// 404接口
type ErrNotFound interface {
	NotFound()
}

// ErrInvalidParameter signals that the user input is invalid
type ErrInvalidParameter interface {
	InvalidParameter()
}

// ErrConflict signals that some internal state conflicts with the requested action and can't be performed.
// A change in state should be able to clear this error.
type ErrConflict interface {
	Conflict()
}

// ErrUnauthorized is used to signify that the user is not authorized to perform a specific action
type ErrUnauthorized interface {
	Unauthorized()
}

// ErrUnavailable signals that the requested action/subsystem is not available.
type ErrUnavailable interface {
	Unavailable()
}

// 403接口
type ErrForbidden interface {
	Forbidden()
}

// ErrSystem signals that some internal error occurred.
// An example of this would be a failed mount request.
type ErrSystem interface {
	System()
}

// ErrNotModified signals that an action can't be performed because it's already in the desired state
type ErrNotModified interface {
	NotModified()
}

// ErrNotImplemented signals that the requested action/feature is not implemented on the system as configured.
type ErrNotImplemented interface {
	NotImplemented()
}

// 未知错误类型，需实现Unknown()方法
type ErrUnknown interface {
	Unknown()
}

// ErrCancelled signals that the action was cancelled.
type ErrCancelled interface {
	Cancelled()
}

// ErrDeadline signals that the deadline was reached before the action completed.
type ErrDeadline interface {
	DeadlineExceeded()
}

// ErrDataLoss indicates that data was lost or there is data corruption.
type ErrDataLoss interface {
	DataLoss()
}

// stack represents a stack of program counters.
type stack []uintptr

// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
// 如果error未实现Cause方法，error会直接返回；如果错误为空，直接返回nil
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	msg string
	*stack
}

func (f *fundamental) Error() string { return f.msg }

type withMessage struct {
	cause error
	msg   string
}

// 不初始化报错
func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }

type withStack struct {
	error
	*stack
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return &withStack{
		err,
		callers(),
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		err,
		callers(),
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(),
	}
}

// NotFound is a helper to create an error of the class with the same name from any error type
func NotFound(err error) error {
	if err == nil || IsNotFound(err) {
		return err
	}
	return errNotFound{err}
}

// 判断404状态码
func IsNotFound(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)
	return ok
}

// Conflict is a helper to create an error of the class with the same name from any error type
func Conflict(err error) error {
	if err == nil || IsConflict(err) {
		return err
	}
	return errConflict{err}
}

// IsConflict returns if the passed in error is an ErrConflict
func IsConflict(err error) bool {
	_, ok := getImplementer(err).(ErrConflict)
	return ok
}

// InvalidParameter is a helper to create an error of the class with the same name from any error type
func InvalidParameter(err error) error {
	if err == nil || IsInvalidParameter(err) {
		return err
	}
	return errInvalidParameter{err}
}

// IsInvalidParameter returns if the passed in error is an ErrInvalidParameter
func IsInvalidParameter(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidParameter)
	return ok
}

// Unauthorized is a helper to create an error of the class with the same name from any error type
func Unauthorized(err error) error {
	if err == nil || IsUnauthorized(err) {
		return err
	}
	return errUnauthorized{err}
}

// IsUnauthorized returns if the passed in error is an ErrUnauthorized
func IsUnauthorized(err error) bool {
	_, ok := getImplementer(err).(ErrUnauthorized)
	return ok
}

// Unavailable is a helper to create an error of the class with the same name from any error type
func Unavailable(err error) error {
	if err == nil || IsUnavailable(err) {
		return err
	}
	return errUnavailable{err}
}

// IsUnavailable returns if the passed in error is an ErrUnavailable
func IsUnavailable(err error) bool {
	_, ok := getImplementer(err).(ErrUnavailable)
	return ok
}

// Forbidden is a helper to create an error of the class with the same name from any error type
func Forbidden(err error) error {
	if err == nil || IsForbidden(err) {
		return err
	}
	return errForbidden{err}
}

// IsForbidden returns if the passed in error is an ErrForbidden
func IsForbidden(err error) bool {
	_, ok := getImplementer(err).(ErrForbidden)
	return ok
}

// System is a helper to create an error of the class with the same name from any error type
func System(err error) error {
	if err == nil || IsSystem(err) {
		return err
	}
	return errSystem{err}
}

// IsSystem returns if the passed in error is an ErrSystem
func IsSystem(err error) bool {
	_, ok := getImplementer(err).(ErrSystem)
	return ok
}

// NotModified is a helper to create an error of the class with the same name from any error type
func NotModified(err error) error {
	if err == nil || IsNotModified(err) {
		return err
	}
	return errNotModified{err}
}

// IsNotModified returns if the passed in error is a NotModified error
func IsNotModified(err error) bool {
	_, ok := getImplementer(err).(ErrNotModified)
	return ok
}

// NotImplemented is a helper to create an error of the class with the same name from any error type
func NotImplemented(err error) error {
	if err == nil || IsNotImplemented(err) {
		return err
	}
	return errNotImplemented{err}
}

// IsNotImplemented returns if the passed in error is an ErrNotImplemented
func IsNotImplemented(err error) bool {
	_, ok := getImplementer(err).(ErrNotImplemented)
	return ok
}

// IsDataLoss returns if the passed in error is an ErrDataLoss
func IsDataLoss(err error) bool {
	_, ok := getImplementer(err).(ErrDataLoss)
	return ok
}

// IsDeadline returns if the passed in error is an ErrDeadline
func IsDeadline(err error) bool {
	_, ok := getImplementer(err).(ErrDeadline)
	return ok
}

// IsCancelled returns if the passed in error is an ErrCancelled
func IsCancelled(err error) bool {
	_, ok := getImplementer(err).(ErrCancelled)
	return ok
}

// Unknown is a helper to create an error of the class with the same name from any error type
// 判断err是否为空，非空且为`未知`错误，直接返回错误；除此之外直接返回错误err信息
func Unknown(err error) error {
	if err == nil || IsUnknown(err) {
		return err
	}
	return errUnknown{err}
}

// IsUnknown returns if the passed in error is an ErrUnknown
func IsUnknown(err error) bool {
	_, ok := getImplementer(err).(ErrUnknown)
	return ok
}

// 解析错误类型
func getImplementer(err error) error {

	// 判断错误类型
	switch e := err.(type) {
	// 如果为以下类型，直接返回
	case
		ErrNotFound,
		ErrInvalidParameter,
		ErrConflict,
		ErrUnauthorized,
		ErrUnavailable,
		ErrForbidden,
		ErrSystem,
		ErrNotModified,
		ErrNotImplemented,
		ErrCancelled,
		ErrDeadline,
		ErrDataLoss,
		ErrUnknown:
		return err
	//
	case causer:
		return getImplementer(e.Cause())
		// 其他情况直接返回
	default:
		return err
	}
}
