package findface

import (
	"context"
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

	var result *FaceDetectResult
	resp, err := s.client.Do(ctx, req, &result)
	result.Response = resp
	return result, err
}
