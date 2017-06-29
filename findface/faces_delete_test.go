package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFaceDeleteOptions_Path(t *testing.T) {
	tests := []struct {
		opt  *FaceDeleteOptions
		path string
	}{{
		opt:  &FaceDeleteOptions{FaceID: 203},
		path: "/faces/id/203",
	}, {
		opt:  &FaceDeleteOptions{Meta: "metadata"},
		path: "/faces/meta/metadata",
	}, {
		opt:  &FaceDeleteOptions{Meta: "metadata", GalleryName: "my_gallery"},
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

func TestFacesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/faces/id/203", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Face.Delete(context.Background(), &FaceDeleteOptions{FaceID: 203})
	if err != nil {
		t.Errorf("Face.Delete returned error: %v", err)
	}
}

func TestFacesService_DeleteParamsError(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Face.Delete(context.Background(), &FaceDeleteOptions{})
	if err == nil {
		t.Errorf("Expected error.")
	}

}
