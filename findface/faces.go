package findface

type BoundingBox struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

type Face struct {
	BoundingBox

	ID int `json:"id"`

	// Metadata string that you can use to store any information associated with the face.
	Meta string `json:"meta"`

	// Age
	Age int `json:"age"`

	// List of emotions
	Emotions []string `json:"emotions"`
	// Gender
	Gender string `json:"gender"`

	// Url of the photo
	Photo     string `json:"photo"`
	PhotoHash string `json:"photo_hash"`
	Thumbnail string `json:"thumbnail"`
	Timestamp string `json:"timestamp"`

	// List of gallery names to add face(s) to.
	Galleries []string `json:"galleries"`
}

type FacesService service
