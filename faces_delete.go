package findface

import (
	"context"
	"fmt"
	"path"
	"strconv"
)

type FacesDeleteResponse struct {
	FindFaceResponse
	Error *FindFaceError
}

type FaceDeleteOptions struct {
	// Face ID
	FaceID int
	// Gallery name
	GalleryName string
	// Metadata string to filter faces by.
	Meta string
}

func (o *FaceDeleteOptions) Path() (string, error) {
	var urlPath string

	switch {
	case o.FaceID > 0:
		urlPath = path.Join("/faces/id", strconv.Itoa(o.FaceID))
	case o.Meta != "" && o.GalleryName != "":
		urlPath = path.Join("/faces/gallery", o.GalleryName, "meta", o.Meta)
	case o.Meta != "":
		urlPath = path.Join("/faces/meta", o.Meta)
	default:
		return "", fmt.Errorf("FaceDeleteOptions is invalid.")
	}
	return urlPath, nil
}

func (s *FacesService) Delete(ctx context.Context, opt *FaceDeleteOptions) (*FacesDeleteResponse, error) {
	urlPath, err := opt.Path()
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("DELETE", urlPath, nil)
	if err != nil {
		return nil, err
	}

	var result = &FacesDeleteResponse{}
	resp, rawResp, err := s.client.Do(ctx, req)
	result.Response = resp
	result.RawResponseBody = rawResp
	return result, err
}
