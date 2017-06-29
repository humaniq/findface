package findface

import (
	"context"
	"net/http"
	"testing"
)

func TestFacesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/faces/id/203", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Face.Delete(context.Background(), 203)
	if err != nil {
		t.Errorf("Face.Delete returned error: %v", err)
	}
}

func TestFacesService_DeleteParamsError(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Face.Delete(context.Background(), 0)
	if err == nil {
		t.Errorf("Expected error.")
	}

}
