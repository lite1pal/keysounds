package main

import (
	"fmt"
	"github.com/ebitengine/oto/v3"
	"github.com/go-audio/audio"
	hook "github.com/robotn/gohook"
)

func ListenEvents(config *Config, sampleRate uint32, pcmBuffer *audio.IntBuffer, ctx *oto.Context) {
	events := hook.Start()
	defer hook.End()
	for event := range events {
		if event.Kind != hook.KeyDown {
			continue
		}
		fmt.Printf("%+v\n", event)

		// map event keys to config key strings
		key := EventToKey(int(event.Keycode))

		definition, ok := config.Definitions[key]

		if !ok {
			continue
		}

		// get a segment from the wav to play exact key sound
		startMs := definition.Timing[0][0]
		endMs := definition.Timing[0][1]
		pcmData := GetSegment(startMs, endMs, sampleRate, pcmBuffer)

		// play
		go CreateAndPlay(ctx, pcmData)
	}
}
