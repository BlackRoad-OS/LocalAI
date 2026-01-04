package backend

import (
	"github.com/BlackRoad-OS/LocalAI/core/config"

	"github.com/BlackRoad-OS/LocalAI/pkg/grpc"
	"github.com/BlackRoad-OS/LocalAI/pkg/model"
)

func StoreBackend(sl *model.ModelLoader, appConfig *config.ApplicationConfig, storeName string, backend string) (grpc.Backend, error) {
	if backend == "" {
		backend = model.LocalStoreBackend
	}
	sc := []model.Option{
		model.WithBackendString(backend),
		model.WithModel(storeName),
	}

	return sl.Load(sc...)
}
