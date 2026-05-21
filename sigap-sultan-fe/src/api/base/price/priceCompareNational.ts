import { baseFetcher } from "../fetcher";
import {
  apiPriceCompareNational,
  apiPriceCompareNationalCityHistory,
  apiPriceCompareNationalCommodityHistory,
  apiPriceCompareNationalList,
} from "../api-paths";
import {
  PriceTier,
  PriceTierType,
  PriceCompareNationalCityMap,
  PriceCompareNationalList,
  PriceCompareNational,
  PriceCompareNationalProvince,
  PriceCompareNationalProvinceMap,
  PriceCompareNationalCityHistory,
  PriceCompareNationalCommodityHistory,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export interface PriceCompareNationalMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceCompareNationalMapResponse {
  commodity: {
    id: number;
    name: string;
  };
  summary: Record<string, number>; // key: PriceTierType
  provincePrice: PriceCompareNationalProvinceMap;
  priceLevel: PriceCompareNationalCityMap[];
  priceTier: Record<string, PriceTier>; // key: PriceTierType
  priceTierCode: PriceTierType[];
}

export interface PriceCompareNationalListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceCompareNationalListResponse {
  commodity: {
    id: number;
    name: string;
  };
  nationalPrice: PriceCompareNational | null;
  provincePrice: PriceCompareNationalProvince | null;
  cityPrice: PriceCompareNationalList[] | null;
}

export interface PriceCompareNationalCityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  provinceId?: number;
  countryId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceCompareNationalCityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  priceDiff: PriceCompareNationalCityHistory[];
}

export interface PriceCompareNationalCommodityHistoryRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  cityId?: number;
  commodityId?: number;
  selectedDate?: string;
  status?: PriceTierType;
}

export interface PriceCompareNationalCommodityHistoryResponse {
  commodity: {
    id: number;
    name: string;
  };
  priceDiff: PriceCompareNationalCommodityHistory[];
}

class PriceCompareNationalApi {
  getPriceCompareNationalMap(request: PriceCompareNationalMapRequest) {
    return baseFetcher.apiGet<PriceCompareNationalMapResponse>({
      path: `${apiPriceCompareNational}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareNationalList(request: PriceCompareNationalListRequest) {
    return baseFetcher.apiGet<PriceCompareNationalListResponse>({
      path: `${apiPriceCompareNationalList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareNationalCityHistory(
    request: PriceCompareNationalCityHistoryRequest
  ) {
    return baseFetcher.apiGet<PriceCompareNationalCityHistoryResponse>({
      path: `${apiPriceCompareNationalCityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceCompareNationalCommodityHistory(
    request: PriceCompareNationalCommodityHistoryRequest
  ) {
    return baseFetcher.apiGet<PriceCompareNationalCommodityHistoryResponse>({
      path: `${apiPriceCompareNationalCommodityHistory}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceCompareNationalApi = new PriceCompareNationalApi();
