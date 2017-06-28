package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Detect(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/detect", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		err := writeResponseFromFile(w, "faces/detect.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceDetectOptions{
		Photo:    "http://example.com/any_image_url.jpg",
		Emotions: true,
		Gender:   true,
		Age:      true,
	}

	result, err := client.Face.Detect(context.Background(), opt)
	if err != nil {
		t.Errorf("Face.Detect returned error: %v", err)
	}
	face := result.Faces[0]
	wantedFace := &Face{
		ID:       0,
		Age:      36,
		Emotions: []string{"neutral", "happy"},
		Gender:   "female",
		BoundingBox: BoundingBox{
			X1: 236,
			X2: 311,
			Y1: 345,
			Y2: 419,
		},
	}

	testDeepEqual(t, face, wantedFace, "Face.Create")
}
