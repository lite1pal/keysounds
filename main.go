package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ebitengine/oto/v3"
	"github.com/go-audio/wav"
	hook "github.com/robotn/gohook"
	"log"
	"os"
	"time"
)

type Definition struct {
	Timing [][]float64 `json:"timing"`
}

type Config struct {
	Definitions map[string]Definition `json:"definitions"`
}

func main() {
	fileBytes, err := os.ReadFile("sound/audio.wav")
	if err != nil {
		log.Fatalf("read WAV file: %v", err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decoderWav := wav.NewDecoder(fileBytesReader)
	if !decoderWav.IsValidFile() {
		log.Fatal("invalid wav file")
	}

	pcmBuffer, err := decoderWav.FullPCMBuffer()
	if err != nil {
		log.Fatalf("get full pcm buffer: %v", err)
	}

	options := &oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	ctx, ready, err := oto.NewContext(options)
	if err != nil {
		log.Fatalf("initialize audio: %v", err)
	}

	<-ready

	if err := ctx.Err(); err != nil {
		log.Fatalf("audio context error: %v", err)
	}

	// read json config
	jsonFileBytes, err := os.ReadFile("sound/config.json")
	if err != nil {
		log.Fatalf("read json file: %v", err)
	}
	var config Config
	err = json.Unmarshal(jsonFileBytes, &config)
	if err != nil {
		log.Fatalf("parse json file: %v", err)
	}

	channels := 2
	sampleRate := 48000.0

	events := hook.Start()
	defer hook.End()
	for event := range events {
		if event.Kind != hook.KeyDown {
			continue
		}
		fmt.Printf("%+v\n", event)

		key := eventToKey(int(event.Keycode))

		definition, ok := config.Definitions[key]

		if !ok {
			continue
		}

		startMs := definition.Timing[0][0]
		endMs := definition.Timing[0][1]
		startFrame := int(startMs * float64(sampleRate) / 1000)
		endFrame := int(endMs * float64(sampleRate) / 1000)

		startSample := startFrame * channels
		endSample := endFrame * channels

		pcmSamples := pcmBuffer.Data[startSample:endSample]
		pcmData := make([]byte, len(pcmSamples)*2)
		for i, sample := range pcmSamples {
			binary.LittleEndian.PutUint16(
				pcmData[i*2:],
				uint16(int16(sample)),
			)
		}

		// create a player
		player := ctx.NewPlayer(bytes.NewReader(pcmData))

		player.Play()

		for player.IsPlaying() {
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func eventToKey(key int) string {
	var textKey string
	switch key {
	case 16:
		textKey = "KeyQ"
	case 17:
		textKey = "KeyW"
	case 18:
		textKey = "KeyE"
	case 19:
		textKey = "KeyR"
	case 20:
		textKey = "KeyT"
	case 21:
		textKey = "KeyY"
	case 22:
		textKey = "KeyU"
	case 23:
		textKey = "KeyI"
	case 24:
		textKey = "KeyO"
	case 25:
		textKey = "KeyP"
	case 26:
		textKey = "BracketLeft"
	case 27:
		textKey = "BracketRight"
	case 14:
		textKey = "Backspace"
	}

	return textKey
}
