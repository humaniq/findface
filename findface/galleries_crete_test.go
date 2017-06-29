package findface

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGalleriesService_Create(t *testing.T) {
	setup()
	defer teardown()

	var wantedData = `{"name":"my_gallery"}`

	mux.HandleFunc("/galleries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		if body == nil {
			t.Error("Body shuld be present")
		}

		cleanData := strings.TrimSpace(string(body))
		if cleanData != wantedData {
			t.Errorf("Result: %+v, Wanted: %+v", cleanData, wantedData)
		}

		w.WriteHeader(http.StatusCreated)
	})

	resultResponse, err := client.Galleries.Create(context.Background(), "my_gallery")
	if err != nil {
		t.Error(err)
	}

	testDeepEqual(t, resultResponse.Response.StatusCode, http.StatusCreated, "Status")
}
