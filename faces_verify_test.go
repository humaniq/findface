package findface

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	firstPhoto  = "https://example.com/photo1.png"
	secondPhoto = "https://example.com/photo2.png"
)

func TestFacesService_Verify(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		defer r.Body.Close()
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		data := map[string]interface{}{}
		err = json.Unmarshal(requestBody, &data)
		if err != nil {
			t.Error(err)
		}

		if data["photo1"] != firstPhoto {
			t.Errorf("Expect photo1 to be %s, got %s", firstPhoto, data["photo1"])
		}
		if data["photo2"] != secondPhoto {
			t.Errorf("Expect photo2 to be %s, got %s", secondPhoto, data["photo2"])
		}

		firstBoundingBoxes := data["bbox1"].([]interface{})
		if len(firstBoundingBoxes) != 1 {
			t.Errorf("Expect bbox1 to have 1 element")
		}

		firstBoundingBox := firstBoundingBoxes[0].(map[string]interface{})
		expectedFirstBoundedBox := map[string]interface{}{
			"x1": 610.0, "y1": 157.0, "x2": 796.0, "y2": 342.0,
		}
		testDeepEqual(t, firstBoundingBox, expectedFirstBoundedBox, "Face.Verify - Request bbox1")

		secondBoundingBoxes := data["bbox2"].([]interface{})
		if len(secondBoundingBoxes) != 1 {
			t.Errorf("Expect bbox2 to have 1 element")
		}

		secondBoundingBox := secondBoundingBoxes[0].(map[string]interface{})
		expectedSecondBoundedBox := map[string]interface{}{
			"x1": 584.0, "y1": 163.0, "x2": 807.0, "y2": 386.0,
		}
		testDeepEqual(t, secondBoundingBox, expectedSecondBoundedBox, "Face.Verify - Request bbox2")

		err = writeResponseFromFile(w, "faces/verify.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceVerifyOptions{
		FirstPhoto:          firstPhoto,
		FirstBoundingBoxes:  []*BoundingBox{&BoundingBox{X1: 610, X2: 796, Y1: 157, Y2: 342}},
		SecondPhoto:         secondPhoto,
		SecondBoundingBoxes: []*BoundingBox{&BoundingBox{X1: 584, X2: 807, Y1: 163, Y2: 386}},
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

func TestFacesService_VerifyWithoutBoundaryBoxesParams(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		defer r.Body.Close()
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		data := map[string]interface{}{}
		err = json.Unmarshal(requestBody, &data)
		if err != nil {
			t.Error(err)
		}

		_, firstBoundingBoxExists := data["bbox1"]
		if firstBoundingBoxExists {
			t.Error("Expect bbox1 to be missed")
		}

		_, secondBoundingBoxExists := data["bbox2"]
		if secondBoundingBoxExists {
			t.Error("Expect bbox2 to be missed")
		}

		err = writeResponseFromFile(w, "faces/verify.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceVerifyOptions{
		FirstPhoto:  firstPhoto,
		SecondPhoto: secondPhoto,
	}

	_, err := client.Face.Verify(context.Background(), opt)
	if err != nil {
		t.Errorf("Face.Verify returned error: %v", err)
	}
}
