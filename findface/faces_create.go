package findface

import (
	"context"
)

const (
	faceCreatePath = "/face"
)

type FaceCreateOptions struct {
	// Url of the photo
	Photo string `json:"photo"`

	// Specifies behavior in case if multiple faces are detected on the photo; one of:
	// `reject` return an error and a list of faces if more than one face is detected on the provided photo
	// `biggest` (default) search using the biggest face on the photo
	// `all` search for each face found on the photo.
	MulitFaceSelector string `json:"mf_selector"`

	// BoundingBoxes [optional] specifying faces coordinates on the photo.
	BoundingBox *BoundingBox `json:"bbox"`

	// Metadata string that you can use to store any information associated with the face.
	Meta string `json:"meta"`

	// List of gallery names to add face(s) to.
	Galleries []string `json:"galleries"`

	// Set to true to extract emotion info from photo
	Emotions bool `json:"emotions"`

	// Set to true to extract gender info from photo
	Gender bool `json:"gender"`

	// Set to true to xtract age info from photo
	Age bool `json:"age"`
}

type FaceCreateResult struct {
	FindFaceResponse
	Results []struct {
		Face
		Age      int      `json:"age"`
		Emotions []string `json:"emotions"`
		Gender   string   `json:"gender"`
	} `json:"results"`
}

// Processes the provided URL, detects faces and adds the detected faces to the searchable dataset. If there are multiple faces on a photo, only the biggest face is added by default.
func (s *FacesService) Create(ctx context.Context, opt *FaceCreateOptions) (*FaceCreateResult, error) {
	req, err := s.client.NewRequest("POST", faceCreatePath, opt)
	if err != nil {
		return nil, err
	}

	faceCreateResult := &FaceCreateResult{}
	resp, err := s.client.Do(ctx, req, &faceCreateResult)
	faceCreateResult.Response = resp
	return faceCreateResult, err
}
