// A program to play sounds of your choice as you type
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
)

var channels = 2 // mono - 1; stereo - 2
var configPath = "sound/config.json"

func Run() error {
	// read wav file
	fileBytes, err := os.ReadFile("sound/concatenated_audio.wav")
	if err != nil {
		return fmt.Errorf("read WAV file: %v", err)
	}

	// create a reader from the file bytes
	fileBytesReader := bytes.NewReader(fileBytes)

	// create a wav decoder
	decoderWav := wav.NewDecoder(fileBytesReader)
	if !decoderWav.IsValidFile() {
		return fmt.Errorf("invalid wav file")
	}

	// access all pcm data
	pcmBuffer, err := decoderWav.FullPCMBuffer()
	if err != nil {
		return fmt.Errorf("get full pcm buffer: %v", err)
	}

	sampleRate := decoderWav.SampleRate

	// create audio context
	ctx, err := CreateAudioCtx(sampleRate)
	if err != nil {
		return fmt.Errorf("create audio ctx: %v", err)
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("audio ctx error: %v", err)
	}

	// parse json config
	config, err := ParseJsonFile(configPath)
	if err != nil {
		return fmt.Errorf("parse json file: %v", err)
	}

	// start listening for keyboard events
	if err := ListenEvents(config, sampleRate, pcmBuffer, ctx); err != nil {
		return fmt.Errorf("listen events: %v", err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
