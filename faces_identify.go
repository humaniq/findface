package findface

import (
	"context"
	"encoding/json"
	"fmt"
)

type FaceIdentifyOptions struct {
	// Url of the photo
	Photo string `json:"photo"`

	GalleryName string `json:"-"` // Optional

	// Specifies behavior in case if multiple faces are detected on the photo; one of:
	// `reject` return an error and a list of faces if more than one face is detected on the provided photo
	// `biggest` (default) search using the biggest face on the photo
	// `all` search for each face found on the photo.
	MultiFaceSelector string `json:"mf_selector,omitempty"`

	// BoundingBoxes [optional] specifying faces coordinates on the photo.
	BoundingBox *BoundingBox `json:"bbox,omitempty"`

	// [optional]: one of "strict", "medium", "low" [default], "none" or a value between 0 and 1
	// Example: "0.75"
	Threshold string `json:"threshold,omitempty"`

	N int `json:"n,omitempty"`
}

type FaceIdentifyResult struct {
	Confidence float64 `json:"confidence"`
	Face       *Face   `json:"face"`
}

type FaceIdentifyResponse struct {
	FindFaceResponse
	Error      *FindFaceError
	ResultsMap map[string][]*FaceIdentifyResult `json:"results"`
}

// This method searches through the face dataset. The method returns at most n faces (one by default), which are the most similar to the specified face, and the similarity confidence is above the specified threshold.
func (s *FacesService) Identify(ctx context.Context, opt *FaceIdentifyOptions) (*FaceIdentifyResponse, error) {
	var path = "identify"

	if opt.GalleryName != "" {
		path = "face/gallery/" + opt.GalleryName + "/" + path
	}

	req, err := s.client.NewRequest("POST", path, opt)
	if err != nil {
		return nil, err
	}
	result := &FaceIdentifyResponse{}
	resp, rawResp, err := s.client.Do(ctx, req)
	var fErr *FindFaceError
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
	return result, err
}
