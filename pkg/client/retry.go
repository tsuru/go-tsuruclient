// Copyright 2023 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

type IsRetriableFunc func(r *http.Response, err error) bool

func IsTsuruEventLocked(r *http.Response, err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "event locked") {
		return true
	}

	if r == nil {
		return false
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		return false
	}

	if r.Body == nil {
		return false
	}

	// TODO: we should ensure there's a specific header from Tsuru, e.g "supported-tsuru".

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}

	r.Body.Close()

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return strings.Contains(string(body), "event locked")
}

var DefaultBackoff = wait.Backoff{
	Steps:    100,
	Duration: 10 * time.Millisecond,
	Factor:   5.0,
	Jitter:   0.1,
}

type LikelyRetriableFunc func() (*http.Response, error)

func RetryWithBackoff(ctx context.Context, backoff wait.Backoff, retriable IsRetriableFunc, fn LikelyRetriableFunc) error {
	if fn == nil {
		return errors.New("no likely retriable func provided")
	}

	var lastError error
	err := wait.ExponentialBackoffWithContext(ctx, backoff, func(ctx context.Context) (bool, error) {
		response, err := fn()
		switch {
		case err == nil:
			return true, nil

		case retriable(response, err):
			lastError = err
			return false, nil

		default:
			return false, err
		}
	})

	if errors.Is(err, wait.ErrWaitTimeout) {
		return lastError
	}

	return err
}

func RetryOnTsuruEventLocked(ctx context.Context, fn LikelyRetriableFunc) error {
	return RetryWithBackoff(ctx, DefaultBackoff, IsTsuruEventLocked, fn)
}
