package findface

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGalleriesService_List(t *testing.T) {
	setup()
	defer teardown()

	data := `{"results": ["default", "gal1"]}`

	mux.HandleFunc("/galleries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	resultResponse, err := client.Galleries.List(context.Background())
	if err != nil {
		t.Errorf("Galleries.List returned error: %v", err)
	}
	wanted := []string{"default", "gal1"}

	testDeepEqual(t, resultResponse.Galleries, wanted, "Galleries.List")
}
