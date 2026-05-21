import {
  PriceCompareProvince,
  PriceCompareProvinceCityHistory,
  PriceCompareProvinceCommodityHistory,
  PriceCompareProvinceList,
  PriceCompareProvinceMap,
  PriceTier,
  PriceTierType,
} from "@/types/price";
import { baseFetcher } from "../fetcher";
import {
  apiPriceCompareProvince,
  apiPriceCompareProvinceCityHistory,
  apiPriceCompareProvinceCommodityHistory,
  apiPriceCompareProvinceList,
} from "../api-paths";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export interface PriceCompareProvinceMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceCompareProvinceMapResponse {
  summary: Record<string, number>; // key: PriceTierType
  priceLevel: PriceCompareProvinceMap[];
  priceTier: Record<string, PriceTier>; // key: PriceTierType
  priceTierCode: PriceTierType[];
}

export interface PriceCompareProvinceListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceCompareProvinceListResponse {
  commodity: {
    id: number;
    name: string;
  };
  provincePrice: PriceCompareProvince | null;
  priceCity: PriceCompareProvinceList[];
}

export interface PriceCompareProvinceCityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  provinceId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceCompareProvinceCityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  priceDiff: PriceCompareProvinceCityHistory[];
}

export interface PriceCompareProvinceCommodityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  commodityId?: number;
  selectedDate?: string;
  status?: PriceTierType;
}

export interface PriceCompareProvinceCommodityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  priceDiff: PriceCompareProvinceCommodityHistory[];
}

class PriceCompareProvinceApi {
  getPriceCompareProvinceMap(request: PriceCompareProvinceMapRequest) {
    return baseFetcher.apiGet<PriceCompareProvinceMapResponse>({
      path: `${apiPriceCompareProvince}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareProvinceList(request: PriceCompareProvinceListRequest) {
    return baseFetcher.apiGet<PriceCompareProvinceListResponse>({
      path: `${apiPriceCompareProvinceList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareProvinceCityHistory(
    request: PriceCompareProvinceCityHistoryRequest
  ) {
    return baseFetcher.apiGet<PriceCompareProvinceCityHistoryResponse>({
      path: `${apiPriceCompareProvinceCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareProvinceCommodityHistory(
    request: PriceCompareProvinceCommodityHistoryRequest
  ) {
    return baseFetcher.apiGet<PriceCompareProvinceCommodityHistoryResponse>({
      path: `${apiPriceCompareProvinceCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceCompareProvinceApi = new PriceCompareProvinceApi();
