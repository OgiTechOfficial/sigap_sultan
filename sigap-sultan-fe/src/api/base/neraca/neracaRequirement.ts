import { baseFetcher } from "../fetcher";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import {
  apiNeracaRequirementCityHistory,
  apiNeracaRequirementCommodityHistory,
  apiNeracaRequirementByCommodityMap,
  apiNeracaRequirementByCommodityList,
  apiNeracaRequirementByCityMap,
} from "../api-paths";
import {
  StockTierType,
  StockTier,
  StockUnit,
  CityRequirementStockLocationSummary,
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
  cityStock: CityRequirementStockLocationSummary[] | null;
  stockTier: Record<string, StockTier>; // key: StockTierNeracaType
}

interface NeracaRequirementStockByCommodityMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaRequirementStockByCommodityMapResponse {
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

interface NeracaRequirementStockByCommodityListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

interface NeracaRequirementStockByCommodityListResponse {
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

interface NeracaRequirementByCityMapRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  selectedDate?: string;
}

interface NeracaRequirementByCityMapResponse {
  unit: StockUnit;
  unitDiff: StockTierUnit;
  city: {
    id: number;
    name: string;
    assets: Asset;
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
    tier: StockTierType;
  };
  commodityStock: CommodityAvailabilityStockCityMap[];
  stockTier: Record<string, StockTier>; // key: StockTierType
  stockTierCode: StockTierType[];
}

class NeracaRequirementApi {
  getNeracaRequirementStockByCommodityMap(
    request: NeracaRequirementStockByCommodityMapRequest
  ) {
    return baseFetcher.apiGet<NeracaRequirementStockByCommodityMapResponse>({
      path: `${apiNeracaRequirementByCommodityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaRequirementStockByCommodityList(
    request: NeracaRequirementStockByCommodityListRequest
  ) {
    return baseFetcher.apiGet<NeracaRequirementStockByCommodityListResponse>({
      path: `${apiNeracaRequirementByCommodityList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaRequirementCityHistory(request: NeracaLocationSummaryRequest) {
    return baseFetcher.apiGet<NeracaLocationSummaryResponse>({
      path: `${apiNeracaRequirementCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaRequirementCommodityHistory(request: NeracaMapSummaryRequest) {
    return baseFetcher.apiGet<NeracaMapSummaryResponse>({
      path: `${apiNeracaRequirementCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getNeracaRequirementByCityMap(request: NeracaRequirementByCityMapRequest) {
    return baseFetcher.apiGet<NeracaRequirementByCityMapResponse>({
      path: `${apiNeracaRequirementByCityMap}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const neracaRequirementApi = new NeracaRequirementApi();
