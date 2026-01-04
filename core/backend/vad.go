package backend

import (
	"context"

	"github.com/BlackRoad-OS/LocalAI/core/config"
	"github.com/BlackRoad-OS/LocalAI/core/schema"
	"github.com/BlackRoad-OS/LocalAI/pkg/grpc/proto"
	"github.com/BlackRoad-OS/LocalAI/pkg/model"
)

func VAD(request *schema.VADRequest,
	ctx context.Context,
	ml *model.ModelLoader,
	appConfig *config.ApplicationConfig,
	modelConfig config.ModelConfig) (*schema.VADResponse, error) {
	opts := ModelOptions(modelConfig, appConfig)
	vadModel, err := ml.Load(opts...)
	if err != nil {
		return nil, err
	}

	req := proto.VADRequest{
		Audio: request.Audio,
	}
	resp, err := vadModel.VAD(ctx, &req)
	if err != nil {
		return nil, err
	}

	segments := []schema.VADSegment{}
	for _, s := range resp.Segments {
		segments = append(segments, schema.VADSegment{Start: s.Start, End: s.End})
	}

	return &schema.VADResponse{
		Segments: segments,
	}, nil
}
