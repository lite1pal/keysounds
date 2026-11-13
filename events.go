package main

import (
	"fmt"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/go-audio/audio"
	hook "github.com/robotn/gohook"
)

type Segments map[string][]byte

func ListenEvents(config *Config, sampleRate uint32, pcmBuffer *audio.IntBuffer, ctx *oto.Context) error {
	events := hook.Start()
	defer hook.End()

	// get all segments before any events to remove key sound delay
	segments := make(Segments)

	for key, definition := range config.Definitions {
		startMs := float64(definition.Timing[0][0])
		endMs := float64(definition.Timing[0][1])
		pcmData, err := GetSegment(startMs, endMs, sampleRate, pcmBuffer)
		if err != nil {
			return fmt.Errorf("get segment: %v", err)
		}

		segments[key] = pcmData
	}

	for event := range events {
		if event.Kind != hook.KeyDown {
			continue
		}

		HandleEvent(event, segments, ctx)
	}

	fmt.Println("we're out!")

	return nil
}

func HandleEvent(event hook.Event, segments Segments, ctx *oto.Context) {
	// map event keys to config key strings
	key := EventToKey(int(event.Keycode))
	pcmData, ok := segments[key]
	if !ok {
		return
	}
	// play
	go CreateAndPlay(ctx, pcmData)
}
