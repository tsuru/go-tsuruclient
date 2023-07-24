// Copyright 2023 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrTsuruTargetNotProvided       = errors.New("Tsuru target not provided")
	ErrTsuruTokenNotProvided        = errors.New("Tsuru token not provided")
	ErrTsuruServiceNameNotProvided  = errors.New("Tsuru service name not provided")
	ErrTsuruInstanceNameNotProvided = errors.New("Tsuru instance name not provided")

	_ http.RoundTripper = (*TsuruProxyTransport)(nil)
)

type TsuruProxyTransport struct {
	Target   string
	Token    string
	Service  string
	Instance string

	Base http.RoundTripper
}

func (tpt *TsuruProxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if err := tpt.checkRequiredFields(); err != nil {
		return nil, err
	}

	u := fmt.Sprintf("%s/1.0/services/%s/proxy/%s?callback=%s", tpt.Target, tpt.Service, tpt.Instance, r.URL.Path)
	if r.URL.RawQuery != "" {
		u += fmt.Sprintf("&%s", r.URL.RawQuery)
	}

	u2, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	r2 := r.Clone(r.Context())
	r2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tpt.Token))
	r2.URL = u2

	return tpt.base().RoundTrip(r2)
}

func (tpt *TsuruProxyTransport) base() http.RoundTripper {
	if tpt.Base == nil {
		return http.DefaultTransport
	}

	return tpt.Base
}

func (tpt *TsuruProxyTransport) checkRequiredFields() error {
	if tpt.Target == "" {
		return ErrTsuruTargetNotProvided
	}

	if tpt.Token == "" {
		return ErrTsuruTokenNotProvided
	}

	if tpt.Service == "" {
		return ErrTsuruServiceNameNotProvided
	}

	if tpt.Instance == "" {
		return ErrTsuruInstanceNameNotProvided
	}

	return nil
}
