import { baseFetcher } from "../fetcher";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import {
  apiNeracaLastStockByCommodityMap,
  apiNeracaLastStockByCommodityList,
  apiNeracaLastStockCityHistory,
  apiNeracaLastStockCommodityHistory,
  apiNeracaAvailabilityByCityAndCommodityChart,
  apiNeracaLastStockByCityMap,
  apiNeracaLastStockByCityAndCommodity,
  apiNeracaCompareWithPriceAndCommodityHistory,
  apiNeracaExist,
  apiNeracaLatestDateExist,
} from "../api-paths";
import {
  StockTier,
  CityLastStockCommodityMap,
  ProvinceLastStockCommodityList,
  CityLastStockCommodityList,
  StockUnit,
  CityLastStockLocationSummary,
  StockTierUnit,
  CommodityAvailabilityStockCityMap,
  StockByPeriod,
  StockTierNeracaType,
} from "@/types/neraca";
import { Asset } from "@/types/asset";
import { generateEndpointWithParam } from "@/utils/api/endpoint-generator";

export interface NeracaExistRequest {
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface NeracaLatestDateExistRequest {
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

interface NeracaLastStockByCommodityMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaLastStockByCommodityMapResponse {
  unit: StockUnit;
  commodity: {
    id: number;
    name: string;
  };
  summary: Record<string, number>; // key: StockTierNeracaType
  provinceStock: {
    id: number;
    clientId: number;
    province: {
      id: number;
      name: string;
    };
    stock: number;
    tier: StockTierNeracaType;
  };
  cityStock: CityLastStockCommodityMap[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
  stockTierCode: StockTierNeracaType[];
}

interface NeracaLastStockByCommodityListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaLastStockByCommodityListResponse {
  unit: StockUnit;
  commodity: {
    id: number;
    name: string;
  };
  provinceStock: ProvinceLastStockCommodityList;
  cityStock: CityLastStockCommodityList[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
}

export interface NeracaLocationSummaryRequest {
  cityId: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

interface NeracaLocationSummaryResponse {
  unit: StockUnit;
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stockDiff: number;
  stock: {
    date: string;
    stock: number;
    stockFormat: string;
  }[];
  stockTier: Record<string, StockTier>; // key: StockTierType
}

export interface NeracaMapSummaryRequest {
  commodityId?: number;
  selectedDate?: string;
  status?: StockTierNeracaType;
}

interface NeracaMapSummaryResponse {
  unit: StockUnit;
  commodity: {
    id: number;
    name: string;
  };
  cityStock: CityLastStockLocationSummary[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
}

interface NeracaAvailabilityByCityAndCommodityChartRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

interface NeracaAvailabilityByCityAndCommodityChartResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stockDiff: number;
  stock: {
    date: string;
    neraca: number;
    ketersediaan: number;
    kebutuhan: number;
  }[];
}

interface NeracaCompareWithPriceAndCommodityHistoryRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

interface NeracaCompareWithPriceAndCommodityHistoryResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stockDiff: number;
  stock: {
    date: string;
    stock: number;
    stockFormat: string;
    price: number;
    priceFormat: string;
  }[];
}

interface NeracaLastStockByCityMapRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  selectedDate?: string;
}

interface NeracaLastStockByCityMapResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  summary: Record<string, number>; // key: StockTierNeracaType
  provinceStock: {
    id: number;
    clientId: number;
    province: {
      id: number;
      name: string;
    };
    stock: number;
    tier: StockTierNeracaType;
  };
  commodityStock: CommodityAvailabilityStockCityMap[];
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
  stockTierCode: StockTierNeracaType[];
}

interface NeracaLastStockByCityAndCommodityParam {
  cityId: number;
  commodityId: number;
}

interface NeracaLastStockByCityAndCommodityRequest {
  page?: number;
  limit?: number;
  startDate?: string;
  endDate?: string;
}

interface NeracaLastStockByCityAndCommodityResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stockDiff: number;
  stock: StockByPeriod[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierType
}

class NeracaApi {
  getExist(request: NeracaExistRequest) {
    return baseFetcher.apiGet<Record<string, boolean>>({
      path: `${apiNeracaExist}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getLatestDateExist(request: NeracaLatestDateExistRequest) {
    return baseFetcher.apiGet<string>({
      path: `${apiNeracaLatestDateExist}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaLastStockByCommodityMap(
    request: NeracaLastStockByCommodityMapRequest
  ) {
    return baseFetcher.apiGet<NeracaLastStockByCommodityMapResponse>({
      path: `${apiNeracaLastStockByCommodityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaLastStockByCommodityList(
    request: NeracaLastStockByCommodityListRequest
  ) {
    return baseFetcher.apiGet<NeracaLastStockByCommodityListResponse>({
      path: `${apiNeracaLastStockByCommodityList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaLastStockCityHistory(request: NeracaLocationSummaryRequest) {
    return baseFetcher.apiGet<NeracaLocationSummaryResponse>({
      path: `${apiNeracaLastStockCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaLastStockCommodityHistory(request: NeracaMapSummaryRequest) {
    return baseFetcher.apiGet<NeracaMapSummaryResponse>({
      path: `${apiNeracaLastStockCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaAvailabilityByCityAndCommodityChart(
    request: NeracaAvailabilityByCityAndCommodityChartRequest
  ) {
    return baseFetcher.apiGet<NeracaAvailabilityByCityAndCommodityChartResponse>(
      {
        path: `${apiNeracaAvailabilityByCityAndCommodityChart}`,
        config: {
          params: constructUrlSearchParams(request),
        },
      }
    );
  }
  getNeracaCompareWithPriceAndCommodityHistory(
    request: NeracaCompareWithPriceAndCommodityHistoryRequest
  ) {
    return baseFetcher.apiGet<NeracaCompareWithPriceAndCommodityHistoryResponse>(
      {
        path: `${apiNeracaCompareWithPriceAndCommodityHistory}`,
        config: {
          params: constructUrlSearchParams(request),
        },
      }
    );
  }
  getNeracaLastStockByCityMap(request: NeracaLastStockByCityMapRequest) {
    return baseFetcher.apiGet<NeracaLastStockByCityMapResponse>({
      path: `${apiNeracaLastStockByCityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaLastStockByCityAndCommodity(
    params: NeracaLastStockByCityAndCommodityParam,
    request: NeracaLastStockByCityAndCommodityRequest
  ) {
    return baseFetcher.apiGet<NeracaLastStockByCityAndCommodityResponse>({
      path: generateEndpointWithParam(
        `${apiNeracaLastStockByCityAndCommodity}`,
        params
      ),
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const neracaApi = new NeracaApi();
