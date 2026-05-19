package models

type AssetsResponse struct {
	Id
	AssetsType         string `json:"assetsType"`
	AssetsLocation     string `json:"assetsLocation"`
	AssetsLocationType string `json:"assetsLocationType"`
	AssetsMediaType    string `json:"assetsMediaType"`
	AssetsExt          string `json:"assetsExt"`
	AssetsName         string `json:"assetsName"`
	AssetsUrl          string `json:"assetsUrl"`
}

type AssetsRequestParams struct {
	AssetsLocation string `json:"assetsLocation" validate:"required"`
}
