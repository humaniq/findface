package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Verify(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		err := writeResponseFromFile(w, "faces/verify.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceVerifyOptions{
		FirstPhoto:        "https://example.com/photo1.png",
		FirstBoundingBox:  &BoundingBox{X1: 610, X2: 796, Y1: 157, Y2: 342},
		SecondPhoto:       "https://example.com/photo2.png",
		SecondBoundingBox: &BoundingBox{X1: 584, X2: 807, Y1: 163, Y2: 386},
	}

	resultResponse, err := client.Face.Verify(context.Background(), opt)
	result := resultResponse.Results[0]
	if err != nil {
		t.Errorf("Face.Verify returned error: %v", err)
	}

	wanted := &FaceVerifyResult{
		FirstBoundingBox:  &BoundingBox{X1: 610, X2: 796, Y1: 157, Y2: 342},
		SecondBoundingBox: &BoundingBox{X1: 584, X2: 807, Y1: 163, Y2: 386},
		Confidence:        0.9222600758075714,
		Verified:          true,
	}

	testDeepEqual(t, result.FirstBoundingBox, wanted.FirstBoundingBox, "Face.Verify - FirstBoundingBox")
	testDeepEqual(t, result.SecondBoundingBox, wanted.SecondBoundingBox, "Face.Verify - SecondBoundingBox")
	testDeepEqual(t, result.Confidence, wanted.Confidence, "Face.Verify - Confidence")
	testDeepEqual(t, result.Verified, wanted.Verified, "Face.Verify - Verified")
}
