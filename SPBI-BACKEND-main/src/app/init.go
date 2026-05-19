package app

import (
	"sigap-sultan-be/src/app/controllers"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"sigap-sultan-be/src/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type InitConfig struct {
	App       *fiber.App
	Db        *pgxpool.Pool
	Redis     *redis.Client
	EnvConfig *common.EnvConfig
}

func InitModules(initConfig InitConfig) {
	validate := validator.New()
	xValidator := &common.XValidator{
		Validator: validate,
	}

	assetsRepository := repositories.NewAssetsRepository(initConfig.Db)
	priceRepository := repositories.NewPriceRepository(initConfig.Db)
	tmCommodityRepository := repositories.NewTmCommodityRepository(initConfig.Db, assetsRepository)
	tmCityRepository := repositories.NewTmCityRepository(initConfig.Db)
	txFileUploadHistoryRepository := repositories.NewTxFileUploadHistoryRepository(initConfig.Db)
	tmProvinceRepository := repositories.NewTmProvinceRepository(initConfig.Db)
	loginRepository := repositories.NewLoginRepository(initConfig.Db)
	positionRepository := repositories.NewTmPositionRepository(initConfig.Db)
	menuRepository := repositories.NewTmMenuRepository(initConfig.Db)
	neracaRepository := repositories.NewNeracaRepository(initConfig.Db, priceRepository)
	jenisInformasiRepository := repositories.NewTmJenisInformasiRepository(initConfig.Db)

	assetsController := controllers.NewAssetsController(xValidator)

	commodityService := services.NewTmCommodityService(tmCommodityRepository)
	commodityController := controllers.NewTmCommodityController(commodityService)

	cityService := services.NewTmCityService(tmCityRepository)
	cityController := controllers.NewTmCityController(cityService)

	jenisInformasiService := services.NewTmJenisInformasiService(jenisInformasiRepository)
	jenisInformasiController := controllers.NewTmJenisInformasiController(jenisInformasiService)

	menuService := services.NewTmMenuService(menuRepository)
	menuController := controllers.NewTmMenuController(menuService)

	neracaServices := services.NewNeracaService(neracaRepository, tmCommodityRepository, tmCityRepository, txFileUploadHistoryRepository, tmProvinceRepository)
	neracaController := controllers.NewNeracaController(xValidator, neracaServices)

	positionService := services.NewTmPositionService(positionRepository)
	positionController := controllers.NewTmPositionController(positionService)

	priceService := services.NewPriceService(priceRepository, tmCommodityRepository, tmCityRepository, txFileUploadHistoryRepository, tmProvinceRepository, initConfig.Redis)
	priceController := controllers.NewPriceController(xValidator, priceService)

	reportService := services.NewReportService(priceRepository, tmCommodityRepository, tmCityRepository, txFileUploadHistoryRepository, tmProvinceRepository, neracaRepository)
	reportController := controllers.NewReportController(xValidator, reportService, cityService, commodityService)

	tmProvinceService := services.NewTmProvinceService(tmProvinceRepository)
	tmProvinceController := controllers.NewTmProvinceController(tmProvinceService)
	loginService := services.NewLoginService(loginRepository)
	loginController := controllers.NewLoginController(loginService)

	routeConfig := routes.RouteConfig{
		App:                        initConfig.App,
		AssetConroller:             assetsController,
		NeracaController:           neracaController,
		PriceController:            priceController,
		TmCommodityController:      commodityController,
		TmCityController:           cityController,
		TmPositionController:       positionController,
		TmMenuController:           menuController,
		TmJenisInformasiController: jenisInformasiController,
		ReportController:           reportController,
		TmProvinceController:       tmProvinceController,
		LoginController:            loginController,
	}

	routeConfig.Setup()
}
