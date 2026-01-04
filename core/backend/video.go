package backend

import (
	"github.com/BlackRoad-OS/LocalAI/core/config"

	"github.com/BlackRoad-OS/LocalAI/pkg/grpc/proto"
	model "github.com/BlackRoad-OS/LocalAI/pkg/model"
)

func VideoGeneration(height, width int32, prompt, negativePrompt, startImage, endImage, dst string, numFrames, fps, seed int32, cfgScale float32, step int32, loader *model.ModelLoader, modelConfig config.ModelConfig, appConfig *config.ApplicationConfig) (func() error, error) {

	opts := ModelOptions(modelConfig, appConfig)
	inferenceModel, err := loader.Load(
		opts...,
	)
	if err != nil {
		return nil, err
	}

	fn := func() error {
		_, err := inferenceModel.GenerateVideo(
			appConfig.Context,
			&proto.GenerateVideoRequest{
				Height:         height,
				Width:          width,
				Prompt:         prompt,
				NegativePrompt: negativePrompt,
				StartImage:     startImage,
				EndImage:       endImage,
				NumFrames:      numFrames,
				Fps:            fps,
				Seed:           seed,
				CfgScale:       cfgScale,
				Step:           step,
				Dst:            dst,
			})
		return err
	}

	return fn, nil
}
