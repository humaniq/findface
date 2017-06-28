package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/face", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		err := writeResponseFromFile(w, "faces/create_success.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceCreateOptions{
		Photo:     "http://example.com/any_image_url.jpg",
		Meta:      "Sam Berry",
		Galleries: []string{"default", "ppl"},
	}

	faceCreateResult, err := client.Face.Create(context.Background(), opt)
	if err != nil {
		t.Errorf("Face.Create returned error: %v", err)
	}
	result := faceCreateResult.Results[0]
	wantedFace := Face{
		ID:        2333,
		Meta:      "Sam Berry",
		PhotoHash: "dc7ac54590729669ca869a18d92cd05e",
		Timestamp: "2016-06-13T11:06:42.075754",
		Galleries: []string{"default", "ppl"},
		BoundingBox: BoundingBox{
			X1: 225,
			X2: 307,
			Y1: 345,
			Y2: 428,
		},
	}

	testDeepEqual(t, result.Face, wantedFace, "Face.Create")
	testDeepEqual(t, result.Age, 40, "Face.Create")
	testDeepEqual(t, result.Emotions, []string{"neutral", "surprised"}, "Face.Create")
}
