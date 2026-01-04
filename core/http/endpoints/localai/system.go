package localai

import (
	"github.com/labstack/echo/v4"
	"github.com/BlackRoad-OS/LocalAI/core/config"
	"github.com/BlackRoad-OS/LocalAI/core/schema"
	"github.com/BlackRoad-OS/LocalAI/pkg/model"
)

// SystemInformations returns the system informations
// @Summary Show the LocalAI instance information
// @Success 200 {object} schema.SystemInformationResponse "Response"
// @Router /system [get]
func SystemInformations(ml *model.ModelLoader, appConfig *config.ApplicationConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		availableBackends := []string{}
		loadedModels := ml.ListLoadedModels()
		for b := range appConfig.ExternalGRPCBackends {
			availableBackends = append(availableBackends, b)
		}
		for b := range ml.GetAllExternalBackends(nil) {
			availableBackends = append(availableBackends, b)
		}

		sysmodels := []schema.SysInfoModel{}
		for _, m := range loadedModels {
			sysmodels = append(sysmodels, schema.SysInfoModel{ID: m.ID})
		}
		return c.JSON(200,
			schema.SystemInformationResponse{
				Backends: availableBackends,
				Models:   sysmodels,
			},
		)
	}
}
