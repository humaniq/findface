package findface

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestFacesService_Update(t *testing.T) {
	setup()
	defer teardown()

	var wantedData = `{"meta":"New meta","galleries":{"add":["list","of","galleries","to","add"]}}`

	mux.HandleFunc("/faces/id/203", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
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
	})

	opt := &FaceUpdateOptions{
		FaceID: 203,
		Meta:   "New meta",
		Galleries: map[string][]string{
			"add": []string{"list", "of", "galleries", "to", "add"},
		},
	}
	_, err := client.Face.Update(context.Background(), opt)
	if err != nil {
		t.Error(err)
	}
}
