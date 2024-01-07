# Video Streaming Service

This is a video streaming service implemented in Go.

## How to Run

TODO: Add instructions on how to build and run the services.

```
/video-streaming-service
|-- /origin
|   |-- main.go (code for setting up the origin source of the video)
|-- /transformation-service
|   |-- main.go (code for the Transformation Service)
|-- /display-service
|   |-- main.go (code for the final display origin)
|-- /common
|   |-- /utils
|       |-- network.go (common networking functions)
|       |-- video.go (common video processing functions)
|-- /assets
|   |-- video.mp4 (sample video file for testing)
|-- README.md (project documentation)
|-- .gitignore (specifies which files Git should ignore)
```