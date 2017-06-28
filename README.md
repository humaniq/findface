# findface #
Golang client for FindFace cloud API

findface requires Go version 1.7 or greater.

## Usage ##

```go
import "github.com/humaniq/findface/findface"
```

Create new Findface.pro client, then use the various services on the client to
access different parts of the FindFace.pro API. For example:

You shuld request Authentication token.

```go
client := findface.NewAuthClient(token, nil)

// list all faces
result, err := client.Face.List(context.Background(), &FaceListOptions{})
if err != nil {
  log.Error(err)
}
```

You can specify options:
```go

// List all faces from `my_gallery`
opt := &FaceListOptions{
  GalleryName: "my_gallery",
}
result, err := client.Face.List(context.Background(), opt)
if err != nil {
  log.Error(err)
}

```
