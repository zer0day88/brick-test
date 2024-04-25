package route

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/zer0day88/brick-test/external/bank_abc"
	"github.com/zer0day88/brick-test/internal/app/domain/repository"
	"github.com/zer0day88/brick-test/internal/app/handler"
	"github.com/zer0day88/brick-test/internal/app/service"
	"github.com/zer0day88/brick-test/pkg/config"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo,
	db *gorm.DB, log zerolog.Logger) {

	bankABCSrv := bank_abc.New(config.Key.External.BankABC.BaseUrl)
	pgRepo := repository.NewPgRepository(db)
	pgSrv := service.NewPGService(log, bankABCSrv, *pgRepo)
	pgHandler := handler.NewPgHandler(pgSrv)

	r := e.Group("/v1")

	pgRoute := r.Group("/pg")

	pgRoute.POST("/inquiry-account", pgHandler.InquiryAccount)
	pgRoute.POST("/transfer", pgHandler.Transfer)
	pgRoute.POST("/notify-transfer", pgHandler.NotifyTransfer)

}
