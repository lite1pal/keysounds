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

		go func() {
			// create a player
			player := ctx.NewPlayer(bytes.NewReader(pcmData))

			player.Play()
		}()
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
	case 30:
		textKey = "KeyA"
	case 31:
		textKey = "KeyS"
	case 32:
		textKey = "KeyD"
	case 33:
		textKey = "KeyF"
	case 34:
		textKey = "KeyG"
	case 35:
		textKey = "KeyH"
	case 36:
		textKey = "KeyJ"
	case 37:
		textKey = "KeyK"
	case 38:
		textKey = "KeyL"
	case 39:
		textKey = "Semicolon"
	case 40:
		textKey = "Quote"
	case 43:
		textKey = "Backslash"
	case 44:
		textKey = "KeyZ"
	case 45:
		textKey = "KeyX"
	case 46:
		textKey = "KeyC"
	case 47:
		textKey = "KeyV"
	case 48:
		textKey = "KeyB"
	case 49:
		textKey = "KeyN"
	case 50:
		textKey = "KeyM"
	case 51:
		textKey = "Comma"
	case 52:
		textKey = "Period"
	case 53:
		textKey = "Slash"
	case 57:
		textKey = "Space"
	case 15:
		textKey = "Tab"
	case 58:
		textKey = "CapsLock"
	case 1:
		textKey = "Escape"
	case 28:
		textKey = "Enter"
	case 42:
		textKey = "ShiftLeft"
	case 54:
		textKey = "ShiftRight"
	case 2:
		textKey = "Digit1"
	case 3:
		textKey = "Digit2"
	case 4:
		textKey = "Digit3"
	case 5:
		textKey = "Digit4"
	case 6:
		textKey = "Digit5"
	case 7:
		textKey = "Digit6"
	case 8:
		textKey = "Digit7"
	case 9:
		textKey = "Digit8"
	case 10:
		textKey = "Digit9"
	case 11:
		textKey = "Digit0"
	case 13:
		textKey = "Equal"
	case 12:
		textKey = "Minus"
	}

	return textKey
}
