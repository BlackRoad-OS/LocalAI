package backend

import (
	"github.com/BlackRoad-OS/LocalAI/core/config"
	"github.com/BlackRoad-OS/LocalAI/core/schema"
	"github.com/BlackRoad-OS/LocalAI/pkg/grpc"
	"github.com/BlackRoad-OS/LocalAI/pkg/model"
)

func ModelTokenize(s string, loader *model.ModelLoader, modelConfig config.ModelConfig, appConfig *config.ApplicationConfig) (schema.TokenizeResponse, error) {

	var inferenceModel grpc.Backend
	var err error

	opts := ModelOptions(modelConfig, appConfig)
	inferenceModel, err = loader.Load(opts...)
	if err != nil {
		return schema.TokenizeResponse{}, err
	}

	predictOptions := gRPCPredictOpts(modelConfig, loader.ModelPath)
	predictOptions.Prompt = s

	// tokenize the string
	resp, err := inferenceModel.TokenizeString(appConfig.Context, predictOptions)
	if err != nil {
		return schema.TokenizeResponse{}, err
	}

	if resp.Tokens == nil {
		resp.Tokens = make([]int32, 0)
	}

	return schema.TokenizeResponse{
		Tokens: resp.Tokens,
	}, nil

}
