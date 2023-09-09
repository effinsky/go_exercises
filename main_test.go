package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type httpStatus = int

func TestRootHandler(t *testing.T) {
	type check func(t *testing.T, act httpStatus, msg string) error

	hasExpCode := func(exp httpStatus) check {
		return func(t *testing.T, act httpStatus, _ string) error {
			t.Helper()
			if diff := cmp.Diff(exp, act); diff != "" {
				return fmt.Errorf("unexpected diff: %s", diff)
			}

			return nil
		}
	}

	hasExpMessage := func(exp string) check {
		return func(t *testing.T, _ httpStatus, act string) error {
			t.Helper()
			if !strings.Contains(exp, act) {
				return fmt.Errorf("expected message missing")
			}

			return nil
		}
	}

	for _, tc := range []struct {
		desc string
		req  *http.Request
		ch   []check
	}{
		{
			desc: "happy path",
			req:  httptest.NewRequest(http.MethodGet, "/root", nil),
			ch:   []check{hasExpCode(http.StatusOK)},
		},
		{
			desc: "invalid path",
			req:  httptest.NewRequest(http.MethodGet, "/roots", nil),
			ch:   []check{hasExpCode(http.StatusBadRequest)},
		},
		{
			desc: "invalid method",
			req:  httptest.NewRequest(http.MethodPut, "/root", nil),
			ch:   []check{hasExpCode(http.StatusMethodNotAllowed)},
		},
		{
			desc: "invalid body",
			req: httptest.NewRequest(
				http.MethodGet,
				"/root",
				strings.NewReader("unbelievable"),
			),
			ch: []check{
				hasExpCode(http.StatusForbidden),
				hasExpMessage("request body must not contain word 'unbelievable'"),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			tc := tc
			rec := httptest.NewRecorder()
			GetRoot(rec, tc.req)

			recCode := rec.Code
			recBody, err := io.ReadAll(rec.Body)
			if err != nil {
				t.Fatalf("reading bytes from recorder body: %v", err)
			}

			for _, c := range tc.ch {
				if err := c(t, recCode, string(recBody)); err != nil {
					log.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}
