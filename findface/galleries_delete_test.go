package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestGalleriesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/galleries/my_gallery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resultResponse, err := client.Galleries.Delete(context.Background(), "my_gallery")
	if err != nil {
		t.Error(err)
	}

	testDeepEqual(t, resultResponse.Response.StatusCode, 204, "Status")
}
