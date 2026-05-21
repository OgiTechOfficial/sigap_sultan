import { baseFetcher } from "../fetcher";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import {
  apiNeracaAvailabilityCityHistory,
  apiNeracaAvailabilityCommodityHistory,
  apiNeracaAvailabilityByCommodityMap,
  apiNeracaAvailabilityByCommodityList,
  apiNeracaAvailabilityByCityMap,
} from "../api-paths";
import {
  StockTierType,
  StockTier,
  StockUnit,
  CityAvailabilityStockLocationSummary,
  CityAvailabilityStockCommodityMap,
  ProvinceAvailabilityStockCommodityList,
  CityAvailabilityStockCommodityList,
  StockTierUnit,
  StockTierNeracaType,
  CommodityAvailabilityStockCityMap,
} from "@/types/neraca";
import { Asset } from "@/types/asset";

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
  cityStock: CityAvailabilityStockLocationSummary[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
}

interface NeracaAvailabilityStockByCommodityMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaAvailabilityStockByCommodityMapResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  commodity: {
    id: number;
    name: string;
  };
  summary: Record<string, number>; // key: StockTierType
  provinceStock: {
    id: number;
    clientId: number;
    province: {
      id: number;
      name: string;
    };
    stock: number;
    stockDiff: number;
    tier: StockTierType;
  };
  cityStock: CityAvailabilityStockCommodityMap[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierType
  stockTierCode: StockTierType[];
}

interface NeracaAvailabilityStockByCommodityListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaAvailabilityStockByCommodityListResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  commodity: {
    id: number;
    name: string;
  };
  provinceStock: ProvinceAvailabilityStockCommodityList;
  cityStock: CityAvailabilityStockCommodityList[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierType
}

interface NeracaAvailabilityByCityMapRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  selectedDate?: string;
}

interface NeracaAvailabilityByCityMapResponse {
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

class NeracaAvailabilityApi {
  getNeracaAvailabilityStockByCommodityMap(
    request: NeracaAvailabilityStockByCommodityMapRequest
  ) {
    return baseFetcher.apiGet<NeracaAvailabilityStockByCommodityMapResponse>({
      path: `${apiNeracaAvailabilityByCommodityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaAvailabilityStockByCommodityList(
    request: NeracaAvailabilityStockByCommodityListRequest
  ) {
    return baseFetcher.apiGet<NeracaAvailabilityStockByCommodityListResponse>({
      path: `${apiNeracaAvailabilityByCommodityList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaAvailabilityCityHistory(request: NeracaLocationSummaryRequest) {
    return baseFetcher.apiGet<NeracaLocationSummaryResponse>({
      path: `${apiNeracaAvailabilityCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaAvailabilityCommodityHistory(request: NeracaMapSummaryRequest) {
    return baseFetcher.apiGet<NeracaMapSummaryResponse>({
      path: `${apiNeracaAvailabilityCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaAvailabilityByCityMap(request: NeracaAvailabilityByCityMapRequest) {
    return baseFetcher.apiGet<NeracaAvailabilityByCityMapResponse>({
      path: `${apiNeracaAvailabilityByCityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const neracaAvailabilityApi = new NeracaAvailabilityApi();
