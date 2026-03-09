package auth_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApi(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{input: http.Header{"Authorization": []string{""}}, want: ""},
		{input: http.Header{"Authorization": []string{"ApiKey"}}, want: ""},
		{input: http.Header{"Authorization": []string{"ApiKey someApiKey"}}, want: "someApiKey"},
	}

	for _, tc := range tests {
		output, _ := auth.GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, output) {
			t.Fatalf("expected: %v, got: %v", tc.want, output)
		}
	}
}
