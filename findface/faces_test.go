package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFaceListOptions_Path(t *testing.T) {
	tests := []struct {
		opt  *FaceListOptions
		path string
	}{{
		opt:  &FaceListOptions{},
		path: "/faces",
	}, {
		opt:  &FaceListOptions{Meta: "metadata"},
		path: "/faces/meta/metadata",
	}, {
		opt:  &FaceListOptions{GalleryName: "my_gallery"},
		path: "/faces/gallery/my_gallery",
	}, {
		opt:  &FaceListOptions{Meta: "metadata", GalleryName: "my_gallery"},
		path: "/faces/gallery/my_gallery/meta/metadata",
	}}

	for _, test := range tests {
		result, err := test.opt.Path()
		if err != nil {
			t.Error(err)
		}
		if result != test.path {
			t.Errorf("Real: %s, but expected: %s", result, test.path)
		}
	}
}

func TestFacesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/faces/gallery/my_gallery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		err := writeResponseFromFile(w, "faces/list_with_gallery.json")
		if err != nil {
			t.Error(err)
		}
	})

	opt := &FaceListOptions{
		GalleryName: "my_gallery",
	}

	result, err := client.Face.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Face.List returned error: %v", err)
	}
	face := result.Faces[0]

	wantedFace := &Face{
		ID:        2563,
		Age:       0,
		Meta:      "Angelina Jolie",
		Photo:     "http://static.findface.pro/sample2.jpg",
		PhotoHash: "dc7ac54590729669ca869a18d92cd05e",
		Thumbnail: "https://static.findface.pro/57726179d6946f02f3763824/9b1dd93259fe87df122cd678ce95b9f9_thumb.jpg",
		Timestamp: "2016-06-13T11:06:42.075754",
		BoundingBox: BoundingBox{
			X1: 225,
			X2: 307,
			Y1: 345,
			Y2: 428,
		},
	}

	testDeepEqual(t, face, wantedFace, "Face.Create")
}
