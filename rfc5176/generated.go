// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc5176

import (
	. "github.com/goconnectx/radius/rfc3576"
)

func init() {
	ErrorCause_Strings[ErrorCause_Value_InvalidAttributeValue] = "Invalid-Attribute-Value"
	ErrorCause_Strings[ErrorCause_Value_MultipleSessionSelectionUnsupported] = "Multiple-Session-Selection-Unsupported"
}

const (
	ErrorCause_Value_InvalidAttributeValue               ErrorCause = 407
	ErrorCause_Value_MultipleSessionSelectionUnsupported ErrorCause = 508
)
