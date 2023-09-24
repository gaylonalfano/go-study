package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// === 'error' is a simple interface that defines a single method: Error()
//
// type error interface {
//   Error() string
// }
//
// Anything that implements its interface is considered an error!
// -- For an interface to be considered 'nil', both the underlying
// TYPE and VALUE must both be nil. E.g., When returning StatusErr,
// the underlying value is nil, but the type is non-nil!
// -- Use Error Wrapping with fmt.Errorf("...%w", err). See Orders API.
// -- Q: When to use errors.Is() vs. errors.
//    A: errors.Is() - when looking for specific instance and/or values
//       errors.As() - when looking for specific type

// === Why not exceptions?
// -- Go's requirement to read all vars keeps you from ignoring errors,
// unless devs use explicit _, or function only returns error

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		// NOTE: Error messages should not be capitalized nor should they
		// end with punctuation or a new line. Convention!
		// NOTE: Should set other return values to their zero value, when
		// a non-nil error is returned!
		return 0, 0, errors.New("denominator is 0")
	}
	// NOTE: 'nil' is the zero value for any interface type
	return numerator / denominator, numerator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

// === Sentinel errors (i.e., signal you can't start or continue processing)
func sentinelErrors() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}

// === Errors are Values
type Status int

// This is an Enum in Go! Using to represent status codes
const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	err     error
}

// Define a method on StatusErr so it meets the error interface
func (se StatusErr) Error() string {
	return se.Message
}

// If you want to UnWrap your custom StatusErr, then needs Unwrap() method
func (se StatusErr) Unwrap() error {
	return se.err
}

//	func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
//		err := login(uid, pwd)
//		if err != nil {
//			return nil, StatusErr{
//				Status:  InvalidLogin,
//				Message: fmt.Sprintf("invalid credentials for user %s", uid),
//			}
//		}
//		data, err := getData(file)
//		if err != nil {
//			return nil, StatusErr{
//				Status:  NotFound,
//				Message: fmt.Sprint("file %s not found", file),
//			}
//		}
//		return data, nil
//	}
//
// Either return nil for functions that return error,
// or wrap the error (see below)
func GenerateError(flag bool) error {
	if flag {
		return StatusErr{Status: NotFound}
	}
	return nil
}

func GenerateErrorReturnErrorType(flag bool) error {
	var genError error
	if flag {
		genError = StatusErr{Status: NotFound}
	}
	return genError
}

// Wrapping Errors with fmt.Errorf() to preserve error and add details
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// Wrap the error
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func wrappingErrors() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}

// errors.Is(err, ErrToCompare) is like a '==' comparison
func errorsIsHelper() {
	err := fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}

// Non-comparable types: slice, map, func, chan, or struct that contains one of those types
type MyErr struct {
	Codes []int // Non-comparable type
}

// Meet the error interface
func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

// Implement Is() on my custom error MyErr
func (me MyErr) Is(target error) bool {
	// NOTE: target.(MyErr) is a type assertion!
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

// errors.As() checks whether any error it wraps matches a specific type
func errorsAsHelper() {
	err := GenerateError(true) // a func that returns an error
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Codes)
	}
}

// errors.As() checks whether any error it wraps matches a specific type
// Can even set the target as a interface instead of an error type
func errorsAsHelperWithInterface() {
	err := GenerateError(true) // a func that returns an error
	var coder interface {
		Code() int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.Code())
	}
}

func main() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	// Error handling
	if err != nil {
		fmt.Println(err)
		// Code 0 = success; Code non-zero = error
		os.Exit(1)
	}
	fmt.Println(remainder, mod)

	wrappingErrors() // in fileChecker: open not_here.txt: no such file or directory
	// open not_here.txt: no such file or directory
	errorsIsHelper()
	errorsAsHelper()

}
