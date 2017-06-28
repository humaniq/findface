package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestMetaService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/meta/gallery/my_gallery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		err := writeResponseFromFile(w, "meta/list.json")
		if err != nil {
			t.Error(err)
		}
	})

	resultResponse, err := client.Meta.List(context.Background(), "my_gallery")
	if err != nil {
		t.Errorf("Face.Verify returned error: %v", err)
	}
	firstResult := resultResponse.Results[0]

	wanted := &MetaData{
		Count: 1,
		Face: &Face{
			ID:          2333,
			Galleries:   []string{"default", "ppl"},
			Photo:       "http://static.findface.pro/sample.jpg",
			PhotoHash:   "dc7ac54590729669ca869a18d92cd05e",
			Thumbnail:   "https://static.findface.pro/57726179d6946f02f3763824/dc7ac54590729669ca869a18d92cd05e_thumb.jpg",
			Timestamp:   "2016-06-13T11:06:42.075754",
			BoundingBox: BoundingBox{X1: 225, X2: 307, Y1: 345, Y2: 428},
			Meta:        "Sam Berry",
		},
		Meta: "Sam Berry",
	}

	testDeepEqual(t, firstResult.Face, wanted.Face, "Meta.List")
	testDeepEqual(t, firstResult.Meta, wanted.Meta, "Meta.List")
}
