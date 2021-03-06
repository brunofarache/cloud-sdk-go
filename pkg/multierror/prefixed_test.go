// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package multierror

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

type apierr struct {
	Err error
}

func (e *apierr) Multierror() *Prefixed {
	payload := reflect.ValueOf(e.Err).Elem().FieldByName("Payload")
	if !payload.IsValid() {
		return nil
	}

	if r, ok := payload.Interface().(*models.BasicFailedReply); ok {
		merr := NewPrefixed("api error")
		for _, e := range r.Errors {
			merr = merr.Append(newBasicFailedReply(e))
		}

		return merr
	}
	return nil
}

func (e *apierr) Error() string { return "" }

func newBasicFailedReply(elem *models.BasicFailedReplyElement) error {
	var code, message = "unknown", "unknown"
	var fields string

	if elem.Code != nil {
		code = *elem.Code
	}

	if elem.Message != nil {
		message = *elem.Message
	}

	if elem.Fields != nil {
		fields = strings.Join(elem.Fields, ", ")
	}

	if fields != "" {
		return fmt.Errorf("%s: %s (%s)", code, message, fields)
	}

	return fmt.Errorf("%s: %s", code, message)
}

func TestNewPrefixed(t *testing.T) {
	type args struct {
		prefix string
		errs   []error
	}
	tests := []struct {
		name string
		args args
		want *Prefixed
	}{
		{
			name: "New without explicit prefix",
			want: &Prefixed{
				Errors: make([]error, 0),
			},
		},
		{
			name: "New with prefix and errors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				errors.New("yet another error"),
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("yet another error"),
				},
			},
		},
		{
			name: "New with prefix and errors and unpacking some other prefixed errors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				&Prefixed{Prefix: "a prefix", Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("a prefix: an error"),
					errors.New("a prefix: another error"),
				},
			},
		},
		{
			name: "New with prefix and errors and unpacking some other prefixed errors and some multierrors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				&Prefixed{Prefix: "a prefix", Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
				}},
				&multierror.Error{Errors: []error{
					errors.New("multierror error"),
					errors.New("multierror error 2"),
				}},
				&Prefixed{Prefix: "some prefix here", Errors: []error{
					errors.New("unprefixed error 1"),
					errors.New("unprefixed error 2"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("a prefix: an error"),
					errors.New("a prefix: another error"),
					errors.New("multierror error"),
					errors.New("multierror error 2"),
					errors.New("unprefixed error 1"),
					errors.New("unprefixed error 2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrefixed(tt.args.prefix, tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPrefixed_Append(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	type args struct {
		errs []error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Prefixed
	}{
		{
			name: "Adds some errors",
			fields: fields{Prefix: "some prefix", Errors: []error{
				errors.New("an error"),
			}},
			args: args{errs: []error{
				&Prefixed{Prefix: "prefix 1", Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
				&Prefixed{Prefix: "prefix 2", Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
				errors.New("a normal error"),
				// No prefix
				&Prefixed{Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix",
				Errors: []error{
					errors.New("an error"),
					errors.New("prefix 1: a prefixed error"),
					errors.New("prefix 1: another error"),
					errors.New("prefix 2: a prefixed error"),
					errors.New("prefix 2: another error"),
					errors.New("a normal error"),
					errors.New("github.com/elastic/cloud-sdk-go/pkg/multierror.unpackErrors: a prefixed error"),
					errors.New("github.com/elastic/cloud-sdk-go/pkg/multierror.unpackErrors: another error"),
				},
			},
		},
		{
			name: "Handles an apierror.Error",
			fields: fields{Prefix: "some prefix", Errors: []error{
				errors.New("an error"),
			}},
			args: args{errs: []error{
				&apierr{Err: &authentication.LoginUnauthorized{
					Payload: &models.BasicFailedReply{Errors: []*models.BasicFailedReplyElement{
						{
							Code:    ec.String("a code"),
							Message: ec.String("message"),
						},
					}},
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix",
				Errors: []error{
					errors.New("an error"),
					errors.New("api error: a code: message"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if got := p.Append(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prefixed.Append() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPrefixed_ErrorOrNil(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "empty error returns nil",
			err:  nil,
		},
		{
			name: "a Prefixed error with multiple errors returns the error",
			fields: fields{Errors: []error{
				errors.New("some error"),
				errors.New("some other error"),
			}},
			err: &Prefixed{Errors: []error{
				errors.New("some error"),
				errors.New("some other error"),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if err := p.ErrorOrNil(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Prefixed.ErrorOrNil() error = %+v, wantErr %+v", err, tt.err)
			}
		})
	}
}

func TestPrefixed_Error(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty errors return empty string",
			want: "",
		},
		{
			name: "Empty FormatFunc uses the default one",
			fields: fields{
				Prefix: "prefix",
				Errors: []error{
					errors.New("some error"),
				},
			},
			want: "prefix: " + multierror.ListFormatFunc([]error{
				errors.New("some error"),
			}),
		},
		{
			name: "Empty FormatFunc uses the default one with more than 1 error",
			fields: fields{
				Prefix: "prefix",
				Errors: []error{
					errors.New("some error"),
					errors.New("another error"),
				},
			},
			want: "prefix: " + multierror.ListFormatFunc([]error{
				errors.New("another error"),
				errors.New("some error"),
			}),
		},
		{
			name: "With a custom FormatFunc ",
			fields: fields{
				Prefix: "prefix",
				FormatFunc: func(es []error) string {
					return "some bogus return"
				},
				Errors: []error{
					errors.New("some error"),
				},
			},
			want: "some bogus return",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if got := p.Error(); got != tt.want {
				t.Errorf("Prefixed.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefixed_Concurrent(t *testing.T) {
	merr := NewPrefixed("someprefix", errors.New("an error"))
	ctr := 10

	var wg sync.WaitGroup
	wg.Add(ctr)
	for i := 0; i < 10; i++ {
		go func(counter int) {
			defer wg.Done()
			_ = merr.Append(fmt.Errorf("error %d", counter))
		}(i)
	}

	wg.Wait()

	assert.EqualError(t, merr,
		"someprefix: 11 errors occurred:\n\t* an error\n\t* error 0\n\t* error 1\n\t* error 2\n\t* error 3\n\t* error 4\n\t* error 5\n\t* error 6\n\t* error 7\n\t* error 8\n\t* error 9\n\n",
	)
}
