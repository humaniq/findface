# findface #
Golang client for FindFace cloud API

[![GoDoc](https://godoc.org/github.com/humaniq/findface?status.svg)](https://godoc.org/github.com/humaniq/findface)
[![Vexor status](https://ci.vexor.io/projects/eaac14f5-b552-4fd8-8f66-70b5cff44115/status.svg)](https://ci.vexor.io/ui/projects/eaac14f5-b552-4fd8-8f66-70b5cff44115/builds)

findface requires Go version 1.7 or greater.

This is the Go wrapper for the [FindFace PRO Facial Recognition API](https://findface.pro/en/). The wrapper was written for [Humaniq.co](https://humaniq.co/) project needs, during the facial image research process. 

## Contributors
This  wrapper was developed by:

* [Alexander Kirillov](https://github.com/saratovsource)


## Usage ##

```go
import "github.com/humaniq/findface"
```

Create new Findface.pro client, then use the various services on the client to
access different parts of the FindFace.pro API. For example:

You should request an Authentication token.

```go
client := findface.NewClient(token, nil, "")

// list all faces
result, err := client.Face.List(context.Background(), &FaceListOptions{})
if err != nil {
  log.Error(err)
}
```

You can specify the options:
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
