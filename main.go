package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func main() {
	videoStream := NewEventStream()
	videoStream.Subscribe(VideoHandler, videoFrameEventCode)

	gpsStream := NewEventStream()
	gpsStream.Subscribe(GPSHandler, gpsEventCode)

	go videoStream.Run()
	go gpsStream.Run()

	var wg sync.WaitGroup
	wg.Add(1); go CaptureVideo(videoStream)
	wg.Add(1); go CaptureGPS(gpsStream)


	wg.Wait()
}

const (
	videoFrameEventCode = iota
	gpsEventCode
)

type VideoEvent struct {
	Frame []byte
}

func VideoHandler(e Event) {
	if e, ok := e.(VideoEvent); ok {
		fmt.Printf("new frame: %d bytes\n", len(e.Frame))
	}
}

type GPSEvent struct {
	Lat  float32
	Long float32
}

func GPSHandler(e Event) {
	if e, ok := e.(GPSEvent); ok {
		fmt.Printf("new coordinates: %f, %f\n", e.Lat, e.Long)
	}
}

func CaptureVideo(stream *EventStream) {
	for {
		frame := EventObject{
			EventType: videoFrameEventCode,
			Event: VideoEvent{
				Frame: make([]byte, rand.Int31n(100)),
			},
		}

		stream.AddEvent(frame)
		time.Sleep(time.Duration(100 * rand.Int31n(10)) * time.Millisecond)
	}
}

func CaptureGPS(stream *EventStream) {
	for {

		coordinates := EventObject{
			EventType: gpsEventCode,
			Event: GPSEvent{
				Lat: rand.Float32() * 180,
				Long: rand.Float32() * 180,
			},
		}

		stream.AddEvent(coordinates)
		time.Sleep(time.Duration(100 * rand.Int31n(10)) * time.Millisecond)
	}
}