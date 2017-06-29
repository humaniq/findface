package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/faces/id/203/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		err := writeResponseFromFile(w, "faces/get.json")
		if err != nil {
			t.Error(err)
		}
	})

	result, err := client.Face.Get(context.Background(), 203)
	if err != nil {
		t.Errorf("Face.Get returned error: %v", err)
	}

	face := result.Faces[0]
	wantedFace := &Face{
		ID:        2333,
		Meta:      "Sam Berry",
		Photo:     "http://static.findface.pro/sample.jpg",
		PhotoHash: "dc7ac54590729669ca869a18d92cd05e",
		Galleries: []string{"default", "ppl"},
		Thumbnail: "https://static.findface.pro/57726179d6946f02f3763824/dc7ac54590729669ca869a18d92cd05e_thumb.jpg",
		Timestamp: "2016-06-13T11:06:42.075754",
		BoundingBox: BoundingBox{
			X1: 225,
			X2: 307,
			Y1: 345,
			Y2: 428,
		},
	}

	testDeepEqual(t, face, wantedFace, "Face.Create")
}
