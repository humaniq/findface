package findface

import (
	"context"
	"fmt"
	"path"
	"strconv"
)

type FaceUpdateOptions struct {
	// Face ID [required]
	FaceID int `json:"-"`
	// New metadata string
	Meta string `json:"meta,omitempty"`
	// JSON dictionary with one key and one value
	// Examples:
	//   Add:
	//     map["add"][]string{"list", "of", "galleries", "to", "add"}
	//   Del:
	//     map["del"][]string{"list", "of", "galleries", "to", "delete", "face", "from"}
	//   Set:
	//     map["set"][]string{"list", "of", "galleries", "to", "replace", "current", "list"}
	Galleries map[string][]string `json:"galleries,omitempty"`
}

type FaceUpdateResponse struct {
	FindFaceResponse
}

func (s *FacesService) Update(ctx context.Context, opt *FaceUpdateOptions) (*FaceUpdateResponse, error) {
	if opt.FaceID <= 0 {
		return nil, fmt.Errorf("FaceID shuld be > 0, but was: %d", opt.FaceID)
	}

	urlPath := path.Join("/faces/id", strconv.Itoa(opt.FaceID))
	req, err := s.client.NewRequest("PUT", urlPath, opt)
	if err != nil {
		return nil, err
	}

	var result = &FaceUpdateResponse{}
	resp, err := s.client.Do(ctx, req, result)
	result.Response = resp
	return result, err
}
