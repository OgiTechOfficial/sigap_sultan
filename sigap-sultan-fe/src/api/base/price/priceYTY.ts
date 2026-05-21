import { baseFetcher } from "../fetcher";
import {
  apiPriceYTY,
  apiPriceYTYCityHistory,
  apiPriceYTYCommodityHistory,
  apiPriceYTYList,
} from "../api-paths";
import {
  PriceTier,
  PriceTierType,
  PriceMTMProvince,
  PriceMTMList,
  PriceMTMProvinceMap,
  PriceMTMCityMap,
  PriceYTYCommodityHistory,
  PriceYTYCityHistory,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export interface PriceYTYMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceYTYMapResponse {
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

export interface PriceYTYListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceYTYListResponse {
  commodity: {
    id: number;
    name: string;
  };
  provincePrice: PriceMTMProvince | null;
  priceCity: PriceMTMList[] | null;
}

export interface PriceYTYCityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  provinceId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceYTYCityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodityInflations: PriceYTYCityHistory[];
}

export interface PriceYTYCommodityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  commodityId?: number;
  selectedDate?: string;
  status?: PriceTierType;
}

export interface PriceYTYCommodityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  cityInflations: PriceYTYCommodityHistory[];
}

class PriceYTYApi {
  getPriceYTYMap(request: PriceYTYMapRequest) {
    return baseFetcher.apiGet<PriceYTYMapResponse>({
      path: `${apiPriceYTY}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTYList(request: PriceYTYListRequest) {
    return baseFetcher.apiGet<PriceYTYListResponse>({
      path: `${apiPriceYTYList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTYCityHistory(request: PriceYTYCityHistoryRequest) {
    return baseFetcher.apiGet<PriceYTYCityHistoryResponse>({
      path: `${apiPriceYTYCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTYCommodityHistory(request: PriceYTYCommodityHistoryRequest) {
    return baseFetcher.apiGet<PriceYTYCommodityHistoryResponse>({
      path: `${apiPriceYTYCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceYTYApi = new PriceYTYApi();
