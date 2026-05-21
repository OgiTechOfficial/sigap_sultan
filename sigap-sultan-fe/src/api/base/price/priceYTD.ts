import { baseFetcher } from "../fetcher";
import {
  apiPriceYTD,
  apiPriceYTDCityHistory,
  apiPriceYTDCommodityHistory,
  apiPriceYTDList,
} from "../api-paths";
import {
  PriceTier,
  PriceTierType,
  PriceMTMProvince,
  PriceMTMList,
  PriceMTMProvinceMap,
  PriceMTMCityMap,
  PriceYTDCityHistory,
  PriceYTDCommodityHistory,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export interface PriceYTDMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceYTDMapResponse {
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

export interface PriceYTDListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceYTDListResponse {
  commodity: {
    id: number;
    name: string;
  };
  provincePrice: PriceMTMProvince | null;
  priceCity: PriceMTMList[] | null;
}

export interface PriceYTDCityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  provinceId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceYTDCityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodityInflations: PriceYTDCityHistory[];
}

export interface PriceYTDCommodityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  commodityId?: number;
  selectedDate?: string;
  status?: PriceTierType;
}

export interface PriceYTDCommodityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  cityInflations: PriceYTDCommodityHistory[];
}

class PriceYTDApi {
  getPriceYTDMap(request: PriceYTDMapRequest) {
    return baseFetcher.apiGet<PriceYTDMapResponse>({
      path: `${apiPriceYTD}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTDList(request: PriceYTDListRequest) {
    return baseFetcher.apiGet<PriceYTDListResponse>({
      path: `${apiPriceYTDList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTDCityHistory(request: PriceYTDCityHistoryRequest) {
    return baseFetcher.apiGet<PriceYTDCityHistoryResponse>({
      path: `${apiPriceYTDCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceYTDCommodityHistory(request: PriceYTDCommodityHistoryRequest) {
    return baseFetcher.apiGet<PriceYTDCommodityHistoryResponse>({
      path: `${apiPriceYTDCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceYTDApi = new PriceYTDApi();
