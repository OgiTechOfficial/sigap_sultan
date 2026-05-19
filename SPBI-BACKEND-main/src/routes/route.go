package routes

import (
	"sigap-sultan-be/src/app/controllers"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                        *fiber.App
	AssetConroller             *controllers.AssetsController
	NeracaController           *controllers.NeracaController
	PriceController            *controllers.PriceController
	TmCommodityController      *controllers.TmCommodityController
	TmCityController           *controllers.TmCityController
	TmPositionController       *controllers.TmPositionController
	TmMenuController           *controllers.TmMenuController
	TmJenisInformasiController *controllers.TmJenisInformasiController
	ReportController           *controllers.ReportController
	TmProvinceController       *controllers.TmProvinceController
	LoginController            *controllers.LoginController
}

func (r RouteConfig) Setup() {
	r.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Root")
	})

	r.App.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})

	r.App.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("Logout")
	})

	// MASTER DATA
	r.App.Get("/city", r.TmCityController.Get)
	r.App.Get("/province", r.TmProvinceController.Get)
	r.App.Get("/province/:id", r.TmProvinceController.GetById)
	r.App.Get("/commodities", r.TmCommodityController.Get)
	r.App.Get("/commodities-grouping", r.TmCommodityController.GetGrouping)

	r.App.Get("/jenis-informasi", r.TmJenisInformasiController.Get)
	r.App.Get("/jenis-informasi/:id", r.TmJenisInformasiController.GetById)

	r.App.Get("/menu", r.TmMenuController.Get)
	r.App.Get("/menu/:id", r.TmMenuController.GetById)
	r.App.Post("/menu", r.TmMenuController.Insert)
	r.App.Put("/menu/:id", r.TmMenuController.Update)
	r.App.Delete("/menu/:id", r.TmMenuController.Delete)

	r.App.Get("/position", r.TmPositionController.Get)
	r.App.Get("/position/:id", r.TmPositionController.GetById)
	r.App.Post("/position", r.TmPositionController.Insert)
	r.App.Delete("/position/:id", r.TmPositionController.Delete)

	// CMS
	r.App.Post("/price", r.PriceController.Upload)
	r.App.Get("/price/upload-history", r.PriceController.UploadHistory)
	r.App.Post("/neraca", r.NeracaController.Upload)
	// r.App.Get("/neraca/upload-history", r.NeracaController.UploadHistory)

	// PRICE TABLE
	r.App.Get("/assets", r.AssetConroller.GetAssets)
	r.App.Get("/assets/not-found", r.AssetConroller.Get404)

	r.App.Get("/price/level-harga/map", r.PriceController.GetLevelHarga)
	r.App.Get("/price/level-harga/map/new", r.PriceController.GetLevelHargaNew)
	r.App.Get("/price/level-harga/list", r.PriceController.GetLevelHargaList)

	// PRICE TABLE - LAST 5 DAYS
	r.App.Get("/price/last-5-days-by-city", r.PriceController.GetLastFiveDays)
	r.App.Get("/price/last-5-days-by-commodity", r.PriceController.GetLastFiveDays)

	// PRICE TABLE - POP UP MAP
	//r.App.Get("/price-by-city-and-commodity", r.PriceController.Get)
	r.App.Get("/price", r.PriceController.Get)

	r.App.Get("/price/compare-province", r.PriceController.GetDibandingkanProvince)
	r.App.Get("/price/compare-province/list", r.PriceController.GetDibandingkanProvinceList)
	r.App.Get("/price/compare-province/city/history", r.PriceController.GetCompareProvinceCityHistory)
	r.App.Get("/price/compare-province/commodity/history", r.PriceController.GetCompareProvinceCommodityHistory)

	r.App.Get("/price/compare-national", r.PriceController.GetDibandingkanNasional)
	r.App.Get("/price/compare-national/list", r.PriceController.GetDibandingkanNasionalList)
	r.App.Get("/price/compare-national/city/history", r.PriceController.GetCompareNationalCityHistory)
	r.App.Get("/price/compare-national/commodity/history", r.PriceController.GetCompareNationalCommodityHistory)

	r.App.Get("/price/mtm", r.PriceController.GetMtM)
	r.App.Get("/price/mtm/list", r.PriceController.GetMtMList)
	r.App.Get("/price/mtm/city/history", r.PriceController.GetMtMCityHistory)
	r.App.Get("/price/mtm/commodity/history", r.PriceController.GetMtMCommodityHistory)

	r.App.Get("/price/ytd", r.PriceController.GetYtd)
	r.App.Get("/price/ytd/list", r.PriceController.GetYtdList)
	r.App.Get("/price/ytd/city/history", r.PriceController.GetYtdCityHistory)
	r.App.Get("/price/ytd/commodity/history", r.PriceController.GetYtdCommodityHistory)

	r.App.Get("/price/yty", r.PriceController.GetYty)
	r.App.Get("/price/yty/list", r.PriceController.GetYtyList)
	r.App.Get("/price/yty/city/history", r.PriceController.GetYtYCityHistory)
	r.App.Get("/price/yty/commodity/history", r.PriceController.GetYtYCommodityHistory)

	r.App.Get("/price/exist", r.PriceController.Exist)
	r.App.Get("/price/latest-date-exist", r.PriceController.LatestDateExist)

	r.App.Get("/neraca/stock-akhir-by-commodity/map", r.NeracaController.GetStockAkhir)
	r.App.Get("/neraca/stock-akhir-by-commodity/list", r.NeracaController.GetStockAkhirByCommodityList)
	r.App.Get("/neraca/stock-akhir-by-commodity/commodity/history", r.NeracaController.GetStockAkhirByCommodityHistory)
	r.App.Get("/neraca/stock-akhir-by-commodity/city/history", r.NeracaController.GetStockAkhirByCommodityCityHistory)
	r.App.Get("/neraca/stock-akhir-by-city/map", r.NeracaController.GetStockAkhirByCity)
	r.App.Get("/neraca/stock-akhir-by-city/list", r.NeracaController.GetStockAkhirByCityList)

	r.App.Get("/neraca/ketersediaan-by-commodity/map", r.NeracaController.GetKetesediaanMap)
	r.App.Get("/neraca/ketersediaan-by-commodity/list", r.NeracaController.GetKetersediaanList)
	r.App.Get("/neraca/ketersediaan-by-commodity/commodity/history", r.NeracaController.GetKetersediaanByCommodityHistory)
	r.App.Get("/neraca/ketersediaan-by-commodity/city/history", r.NeracaController.GetKetersediaanByCommodityCityHistory)
	r.App.Get("/neraca/ketersediaan-by-city/map", r.NeracaController.GetKetersediaanByCity)
	r.App.Get("/neraca/ketersediaan-by-city/list", r.NeracaController.GetKetersediaanByCityList)

	r.App.Get("/neraca/kebutuhan-by-commodity/map", r.NeracaController.GetKebutuhanMap)
	r.App.Get("/neraca/kebutuhan-by-commodity/list", r.NeracaController.GetKebutuhanList)
	r.App.Get("/neraca/kebutuhan-by-commodity/commodity/history", r.NeracaController.GetKebutuhanByCommodityHistory)
	r.App.Get("/neraca/kebutuhan-by-commodity/city/history", r.NeracaController.GetKebutuhanByCommodityCityHistory)
	r.App.Get("/neraca/kebutuhan-by-city/map", r.NeracaController.GetKebutuhanByCity)
	r.App.Get("/neraca/kebutuhan-by-city/list", r.NeracaController.GetKebutuhanByCityList)

	// NERACA - CHART DI BAWAH MAP DAN LIST TABLE
	r.App.Get("/neraca/ketersediaan-by-city-and-commodity/chart", r.NeracaController.GetStockAkhirByCityAndCommodityIdChart)
	r.App.Get("/neraca/stock-akhir/city/:cityId/commodity/:commodityId", r.NeracaController.GetStockAkhirByCityAndCommodityId)
	r.App.Get("/report/price", r.ReportController.GetPriceReport)
	r.App.Get("/report/price/download", r.ReportController.GetPriceReportDownload)
	r.App.Get("/report/neraca", r.ReportController.GetNeracaReport)
	r.App.Get("/report/neraca/download", r.ReportController.GetNeracaReportDownload)

	r.App.Get("/neraca/compare-with-price/commodity/history", r.NeracaController.CompareWithPriceCommodityHistory)

	r.App.Get("/neraca/exist", r.NeracaController.Exist)
	r.App.Get("/neraca/latest-date-exist", r.NeracaController.LatestDateExist)

	//FORGOT
	r.App.Post("/forgot-password", r.LoginController.Forgot)
	r.App.Post("/verify-token", r.LoginController.CheckKode)
	r.App.Post("/reset-password", r.LoginController.ResetPassword)
	r.App.Get("/profile", r.LoginController.Profile)
	r.App.Post("/change-password", r.LoginController.ChangePassword)
	r.App.Post("/profile", r.LoginController.UpdateProfile)

	//r.App.Get("/price/compare-national", r.PriceController.GetDibandingkanNasional)
	//
	//// PRICE TABLE - LAST 5 DAYS
	//r.App.Get("/price/last-5-days-by-city", r.PriceController.GetLastFiveDays)
	//r.App.Get("/price/last-5-days-by-commodity", r.PriceController.GetLastFiveDays)
	//
	//// PRICE TABLE - POP UP MAP
	//r.App.Get("/price-by-city-and-commodity", r.PriceController.Get)
	//
	//// PRICE TABLE - POP UP SUMMARY
	//r.App.Get("/price-commodity", r.PriceController.GetByCommodity)

	//Login
	r.App.Post("/login", r.LoginController.Login)
}
