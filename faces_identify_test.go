package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Identify(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/identify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		err := writeResponseFromFile(w, "faces/identify.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceIdentifyOptions{
		Photo: "http://static.findface.pro/sample.jpg",
	}

	resp, err := client.Face.Identify(context.Background(), opt)
	if err != nil {
		t.Errorf("Face.Identify returned error: %v", err)
	}

	result := resp.ResultsMap["[610, 157, 796, 342]"][0]
	wantedResult := &FaceIdentifyResult{
		Confidence: 1,
		Face: &Face{
			ID:        2333,
			Meta:      "Sam Berry",
			Photo:     "http://static.findface.pro/sample.jpg",
			PhotoHash: "dc7ac54590729669ca869a18d92cd05e",
			Thumbnail: "https://static.findface.pro/57726179d6946f02f3763824/dc7ac54590729669ca869a18d92cd05e_thumb.jpg",
			Galleries: []string{"default", "ppl"},
			Timestamp: "2016-06-13T11:06:42.075754",
			BoundingBox: BoundingBox{
				X1: 225,
				X2: 307,
				Y1: 345,
				Y2: 428,
			},
		},
	}
	testDeepEqual(t, result, wantedResult, "Face.Identify")
}
