package main

import (
	"testing"

	"github.com/go-audio/audio"
)

func TestGetSegment(t *testing.T) {
	t.Run("endMs less than startMs", func(t *testing.T) {
		startMs := 2.0
		endMs := -1.5
		sampleRate := uint32(48000)
		pcm := &audio.IntBuffer{Data: []int{100, 200, 300, 400}}
		_, err := GetSegment(startMs, endMs, sampleRate, pcm)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("startMs or endMs are not positive numbers", func(t *testing.T) {
		startMs := -1.5
		endMs := 50.8
		sampleRate := uint32(48000)
		pcm := &audio.IntBuffer{Data: []int{100, 200, 300, 400}}
		_, err := GetSegment(startMs, endMs, sampleRate, pcm)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
