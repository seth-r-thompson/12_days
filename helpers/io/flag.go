package io

import (
	"errors"
	"flag"
	"reflect"
	"strings"
)

// flaggable is all valid Flag types.
type flaggable interface {
	string | []string | bool
}

// Flag represents a command line flag.
// Flags are passed in with a format --name=value.
// For Flags with []string values, separate values by commas.
type Flag[T flaggable] struct {
	Name        string
	Description string
	Value       T
}

// Parse accepts a list of Flag pointers, defines them as flags, parses them, and updates the Flag values.
func Parse(flags ...any) error {
	for _, f := range flags {
		if reflect.ValueOf(f).Kind() != reflect.Pointer {
			return errors.New("must pass flag as pointer")
		}
	}

	// setup flags
	values := make([]any, len(flags))
	for i, f := range flags {
		switch f.(type) {
		case *Flag[string]:
			values[i] = flag.String(f.(*Flag[string]).Name, f.(*Flag[string]).Value, f.(*Flag[string]).Description)
		case *Flag[[]string]:
			values[i] = flag.String(f.(*Flag[[]string]).Name, "", f.(*Flag[[]string]).Description)
		case *Flag[bool]:
			values[i] = flag.Bool(f.(*Flag[bool]).Name, f.(*Flag[bool]).Value, f.(*Flag[bool]).Description)
		default:
			return errors.New("unsupported flag type")
		}
	}

	// parse flags from command line
	flag.Parse()

	// dereference parsed flags
	for i, f := range flags {
		switch f.(type) {
		case *Flag[string]:
			f.(*Flag[string]).Value = *values[i].(*string)
		case *Flag[[]string]:
			f.(*Flag[[]string]).Value = strings.Split(*values[i].(*string), ",")
		case *Flag[bool]:
			f.(*Flag[bool]).Value = *values[i].(*bool)
		}
	}

	return nil
}
