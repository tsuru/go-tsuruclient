// Copyright 2023 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/tsuru/go-tsuruclient/pkg/client"
)

func TestIsTsuruEventLocked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err      error
		response *http.Response
		expected bool
	}{
		{},

		{
			err:      errors.New("event locked: fulano is running \"app-deploy\""),
			expected: true,
		},

		{
			err: errors.New("something went wrong"),
		},

		{
			err: errors.New("something went wrong"),
			response: &http.Response{
				StatusCode: http.StatusFound,
				Body:       io.NopCloser(strings.NewReader("unexpected status code")),
			},
		},

		{
			err: errors.New("something went wrong"),
			response: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("event locked: jhon.doe@example.org is running \"app-deploy\"")),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.expected, IsTsuruEventLocked(tt.response, tt.err))
		})
	}
}

func TestRetryOnTsuruEventLocked(t *testing.T) {
	t.Parallel()

	var retries int

	tests := []struct {
		ctx      context.Context
		fn       LikelyRetriableFunc
		expected string
	}{
		{
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			fn: func() (*http.Response, error) {
				return nil, nil
			},
			expected: "context canceled",
		},

		{
			ctx:      context.Background(),
			expected: "no likely retriable func provided",
		},

		{
			ctx: context.Background(),
			fn: func() (*http.Response, error) {
				return nil, nil
			},
		},

		{
			ctx: func() context.Context {
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				return ctx
			}(),
			fn: func() (*http.Response, error) {
				return nil, errors.New("event locked: forever")
			},
			expected: "context deadline exceeded",
		},

		{
			ctx: context.Background(),
			fn: func() (*http.Response, error) {
				retries++
				if retries == 5 {
					return nil, nil
				}

				return &http.Response{StatusCode: http.StatusConflict, Body: io.NopCloser(strings.NewReader("event locked"))}, errors.New("something went wrong")
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			retries = 0

			err := RetryOnTsuruEventLocked(tt.ctx, tt.fn)
			if tt.expected == "" {
				assert.NoError(t, err)
				return
			}

			assert.EqualError(t, err, tt.expected)
		})
	}
}
