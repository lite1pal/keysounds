package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ebitengine/oto/v3"
	"github.com/go-audio/audio"
)

func CreateAudioCtx(sampleRate uint32) (*oto.Context, error) {
	options := &oto.NewContextOptions{
		SampleRate:   int(sampleRate),
		ChannelCount: channels,
		Format:       oto.FormatSignedInt16LE,
	}

	// init audio
	ctx, ready, err := oto.NewContext(options)
	if err != nil {
		return nil, fmt.Errorf("initialize audio: %v", err)
	}

	<-ready

	return ctx, nil
}

func GetSegment(startMs float64, endMs float64, sampleRate uint32, pcm *audio.IntBuffer) ([]byte, error) {
	if startMs < 0 || endMs < 0 {
		return nil, fmt.Errorf("startMs and endMs must be positive numbers")
	}
	if startMs > endMs {
		return nil, fmt.Errorf("endMs must be higher than startMs")
	}
	startFrame := int(startMs * float64(sampleRate) / 1000)
	endFrame := int(endMs * float64(sampleRate) / 1000)
	startSample := startFrame * channels
	endSample := endFrame * channels
	pcmSamples := pcm.Data[startSample:endSample]
	pcmData := make([]byte, len(pcmSamples)*2)
	for i, sample := range pcmSamples {
		binary.LittleEndian.PutUint16(
			pcmData[i*2:],
			uint16(int16(sample)),
		)
	}
	return pcmData, nil
}

func CreateAndPlay(ctx *oto.Context, pcmData []byte) {
	player := ctx.NewPlayer(bytes.NewReader(pcmData))
	// play
	player.Play()
}
