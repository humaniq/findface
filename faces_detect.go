package findface

import (
	"context"
	"encoding/json"
	"fmt"
)

type FaceDetectOptions struct {
	// Url of the photo
	Photo string `json:"photo"`
	// Set to true to extract emotion info from photo
	Emotions bool `json:"emotions"`

	// Set to true to extract gender info from photo
	Gender bool `json:"gender"`

	// Set to true to xtract age info from photo
	Age bool `json:"age"`
}

type FaceDetectResult struct {
	FindFaceResponse
	Error       *FindFaceError
	Faces       []*Face
	Orientation int `json:"orientation"`
}

// This method detects faces on the provided image.
// You shuld provide an URL, which the API will use to fetch the image.
func (s *FacesService) Detect(ctx context.Context, opt *FaceDetectOptions) (*FaceDetectResult, error) {
	req, err := s.client.NewRequest("POST", "/detect", opt)
	if err != nil {
		return nil, err
	}

	result := &FaceDetectResult{}
	resp, rawResp, dErr := s.client.Do(ctx, req)
	fErr := &FindFaceError{}
	switch resp.StatusCode {
	case 200:
		unErr := json.Unmarshal(rawResp, &result)
		if unErr != nil {
			return nil, unErr
		}
	case 400:
		unErr := json.Unmarshal(rawResp, &fErr)
		if unErr != nil {
			return nil, unErr
		}
	default:
		err = fmt.Errorf("FindFace returned an unhandled status: %s, body: %s", resp.Status, string(rawResp))
	}
	result.Response = resp
	result.RawResponseBody = rawResp
	result.Error = fErr
	return result, dErr
}
