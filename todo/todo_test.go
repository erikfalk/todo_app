package todo

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	os.Chdir("../")
}

func TestListTodos(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the ListTodos function with the recorder and request
	ListTodos(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Read the expected response from file
	expected_res, err := os.ReadFile("todo/test_responses/index_response.html")
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body
	diff := cmp.Diff(string(expected_res), recorder.Body.String())
	if diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}
}
