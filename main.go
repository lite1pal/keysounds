// A program to play sounds of your choice as you type
package main

import (
	"bytes"
	"log"
	"os"

	"github.com/go-audio/wav"
)

var channels = 2 // mono - 1; stereo - 2
var configPath = "sound/config.json"

func main() {
	// read wav file
	fileBytes, err := os.ReadFile("sound/audio.wav")
	if err != nil {
		log.Fatalf("read WAV file: %v", err)
	}

	// create a reader from the file bytes
	fileBytesReader := bytes.NewReader(fileBytes)

	// create a wav decoder
	decoderWav := wav.NewDecoder(fileBytesReader)
	if !decoderWav.IsValidFile() {
		log.Fatal("invalid wav file")
	}

	// access all pcm data
	pcmBuffer, err := decoderWav.FullPCMBuffer()
	if err != nil {
		log.Fatalf("get full pcm buffer: %v", err)
	}

	sampleRate := decoderWav.SampleRate

	// create audio context
	ctx, err := CreateAudioCtx(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
	if err := ctx.Err(); err != nil {
		log.Fatalf("audio context error: %v", err)
	}

	// parse json config
	config, err := ParseJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// start listening for keyboard events
	ListenEvents(config, sampleRate, pcmBuffer, ctx)
}
