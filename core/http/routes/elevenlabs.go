package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/BlackRoad-OS/LocalAI/core/config"
	"github.com/BlackRoad-OS/LocalAI/core/http/endpoints/elevenlabs"
	"github.com/BlackRoad-OS/LocalAI/core/http/middleware"
	"github.com/BlackRoad-OS/LocalAI/core/schema"
	"github.com/BlackRoad-OS/LocalAI/pkg/model"
)

func RegisterElevenLabsRoutes(app *echo.Echo,
	re *middleware.RequestExtractor,
	cl *config.ModelConfigLoader,
	ml *model.ModelLoader,
	appConfig *config.ApplicationConfig) {

	// Elevenlabs
	ttsHandler := elevenlabs.TTSEndpoint(cl, ml, appConfig)
	app.POST("/v1/text-to-speech/:voice-id",
		ttsHandler,
		re.BuildFilteredFirstAvailableDefaultModel(config.BuildUsecaseFilterFn(config.FLAG_TTS)),
		re.SetModelAndConfig(func() schema.LocalAIRequest { return new(schema.ElevenLabsTTSRequest) }))

	soundGenHandler := elevenlabs.SoundGenerationEndpoint(cl, ml, appConfig)
	app.POST("/v1/sound-generation",
		soundGenHandler,
		re.BuildFilteredFirstAvailableDefaultModel(config.BuildUsecaseFilterFn(config.FLAG_SOUND_GENERATION)),
		re.SetModelAndConfig(func() schema.LocalAIRequest { return new(schema.ElevenLabsSoundGenerationRequest) }))

}
