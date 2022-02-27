package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMainProgram(t *testing.T) {
	cases := []struct {
		Name string
		Args []string
	}{
		{"flags set", []string{"", "-parallel", "1", "google.com"}},
		{"flags set with default value", []string{"", "google.com", "facebook.com"}},
	}
	for _, tc := range cases {
		flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
		os.Args = tc.Args
		main()
	}
}

func TestMakeParallelGetRequests(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}))

	defer svr.Close()
	expectedResult := fmt.Sprintf("%s %s", svr.URL, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	realResult := ParallelGetRequests([]string{svr.URL}, 10)

	if realResult[0] != expectedResult {
		t.Errorf("%s != %s", realResult, expectedResult)
	}
}
