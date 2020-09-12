package tts

import (
	"context"
	"fmt"
	"sync"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type TtsService struct {
	client *texttospeech.Client
}

func New(ctx context.Context) (*TtsService, error) {
	cl, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &TtsService{cl}, nil
}

var cache map[string][]byte = make(map[string][]byte)
var loading map[string]struct{} = make(map[string]struct{})
var mu sync.Mutex

// getCache gets audio from cache. lock mutex before calling this and unlock after
func getCache(key string) []byte {
	if val, exists := cache[key]; exists {
		return val
	}
	// Check if we are already trying to load the same item
	if _, exists := loading[key]; exists {
		mu.Unlock()

		// Try to get from cache every second
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()
		tries := 0
		for range ticker.C {
			mu.Lock()
			// We already tried 3 times, probably something went wrong so remove the loading key
			if tries++; tries > 3 {
				delete(loading, key)
				return nil
			}
			val, exists := cache[key]
			if exists {
				return val
			}
			mu.Unlock()
		}
	}
	return nil
}

// Tts turns text into speech
func (tts *TtsService) Tts(ctx context.Context, text string, lang string) ([]byte, error) {
	cacheKey := lang + text
	mu.Lock()
	cachedVal := getCache(cacheKey)
	if cachedVal != nil {
		mu.Unlock()
		return cachedVal, nil
	}
	loading[cacheKey] = struct{}{}
	mu.Unlock()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: lang,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := tts.client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("error synthesizing speech: %v", err)
	}

	mu.Lock()
	cache[cacheKey] = resp.AudioContent
	mu.Unlock()

	return resp.AudioContent, nil
}
