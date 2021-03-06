package findface

import (
	"context"
)

type FaceCreateOptions struct {
	// Url of the photo
	Photo string `json:"photo"`

	// Specifies behavior in case if multiple faces are detected on the photo; one of:
	// `reject` return an error and a list of faces if more than one face is detected on the provided photo
	// `biggest` (default) search using the biggest face on the photo
	// `all` search for each face found on the photo.
	MultiFaceSelector string `json:"mf_selector,omitempty"`

	// BoundingBoxes [optional] specifying faces coordinates on the photo.
	BoundingBox *BoundingBox `json:"bbox,omitempty"`

	// Metadata string that you can use to store any information associated with the face.
	Meta string `json:"meta,omitempty"`

	// List of gallery names to add face(s) to.
	Galleries []string `json:"galleries,omitempty"`

	// Set to true to extract emotion info from photo
	Emotions bool `json:"emotions,omitempty"`

	// Set to true to extract gender info from photo
	Gender bool `json:"gender,omitempty"`

	// Set to true to xtract age info from photo
	Age bool `json:"age,omitempty"`
}

type FaceCreateResult struct {
	FindFaceResponse
	Faces []*Face `json:"results"`
}

// Processes the provided URL, detects faces and adds the detected faces to the searchable dataset. If there are multiple faces on a photo, only the biggest face is added by default.
func (s *FacesService) Create(ctx context.Context, opt *FaceCreateOptions) (*FaceCreateResult, error) {
	req, err := s.client.NewRequest("POST", "face", opt)
	if err != nil {
		return nil, err
	}

	result := FaceCreateResult{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
