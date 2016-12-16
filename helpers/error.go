// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

// basic error type
type Error struct {
	Context *Context
	Message string
}

func NewError(context *Context, message string) *Error {
	return &Error{Context: context, Message: message}
}

func (err *Error) Error() string {
	if err.Context != nil {
		return err.Context.Description() + ": " + err.Message
	} else {
		return err.Message
	}
}

// container for groups of errors
type ErrorGroup struct {
	Errors []error
}

func NewErrorGroup(errors []error) *ErrorGroup {
	return &ErrorGroup{Errors: errors}
}

func (group *ErrorGroup) Error() string {
	result := ""
	for i, err := range group.Errors {
		if i > 0 {
			result += "\n"
		}
		result += err.Error()
	}
	return result
}
