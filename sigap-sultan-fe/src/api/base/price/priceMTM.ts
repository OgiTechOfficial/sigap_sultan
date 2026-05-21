import { baseFetcher } from "../fetcher";
import {
  apiPriceMTM,
  apiPriceMTMCityHistory,
  apiPriceMTMCommodityHistory,
  apiPriceMTMList,
} from "../api-paths";
import {
  PriceTier,
  PriceTierType,
  PriceMTMProvince,
  PriceMTMList,
  PriceMTMProvinceMap,
  PriceMTMCityMap,
  PriceMTMCityHistory,
  PriceMTMCommodityHistory,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export interface PriceMTMMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceMTMMapResponse {
  commodity: {
    id: number;
    name: string;
  };
  summary: Record<string, number>; // key: PriceTierType
  provincePrice: PriceMTMProvinceMap;
  priceLevel: PriceMTMCityMap[];
  priceTier: Record<string, PriceTier>; // key: PriceTierType
  priceTierCode: PriceTierType[];
}

export interface PriceMTMListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceMTMListResponse {
  commodity: {
    id: number;
    name: string;
  };
  provincePrice: PriceMTMProvince | null;
  priceCity: PriceMTMList[] | null;
}

export interface PriceMTMCityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  provinceId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceMTMCityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodityInflations: PriceMTMCityHistory[];
}

export interface PriceMTMCommodityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  commodityId?: number;
  selectedDate?: string;
  status?: PriceTierType;
}

export interface PriceMTMCommodityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  cityInflations: PriceMTMCommodityHistory[];
}

class PriceMTMApi {
  getPriceMTMMap(request: PriceMTMMapRequest) {
    return baseFetcher.apiGet<PriceMTMMapResponse>({
      path: `${apiPriceMTM}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceMTMList(request: PriceMTMListRequest) {
    return baseFetcher.apiGet<PriceMTMListResponse>({
      path: `${apiPriceMTMList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceMTMCityHistory(request: PriceMTMCityHistoryRequest) {
    return baseFetcher.apiGet<PriceMTMCityHistoryResponse>({
      path: `${apiPriceMTMCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceMTMCommodityHistory(request: PriceMTMCommodityHistoryRequest) {
    return baseFetcher.apiGet<PriceMTMCommodityHistoryResponse>({
      path: `${apiPriceMTMCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceMTMApi = new PriceMTMApi();
