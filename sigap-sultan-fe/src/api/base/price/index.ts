import { baseFetcher } from "../fetcher";
import {
  apiPrice,
  apiPriceLast5DaysByCommodity,
  apiPriceLast5DaysByCity,
  apiPriceExist,
  apiPriceLatestDateExist,
} from "../api-paths";
import {
  PriceLevelMap,
  PriceTier,
  PriceLast5DaysByCommodity,
  PriceLast5DaysByCity,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";
import { PaginationData } from "@/types/api";

export type PriceDiffBy = "province" | "national";

export interface PriceListRequest {
  commodityId?: number;
  cityId?: number;
  provinceId?: number;
  startDate?: string;
  endDate?: string;
  diffBy?: PriceDiffBy;
}

export interface PriceListResponse {
  commodity: {
    id: number;
    name: string;
  };
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  priceDiff: number;
  priceDiffFormat: string;
  price: {
    date: string;
    price: number;
    priceDiff: number;
    priceRupiahFormat: string;
  }[];
}

export interface PriceExistRequest {
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceLatestDateExistRequest {
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface PriceLevelMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceLevelMapResponse {
  priceLevel: PriceLevelMap[];
  priceMin: number;
  priceMax: number;
  priceDiff: number;
  priceBarCategory: number;
  priceTier: PriceTier[];
}

export interface PriceLevelListRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceLast5DaysByCommodityRequest {
  page?: number;
  limit?: number;
  commodityId: number;
}

export interface PriceLast5DaysByCommodityResponse extends PaginationData {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  cities: PriceLast5DaysByCommodity[];
}

export interface PriceLast5DaysByCityRequest {
  page?: number;
  limit?: number;
  cityId: number;
}

export interface PriceLast5DaysByCityResponse extends PaginationData {
  id: number;
  clientId: number;
  commodity: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodities: PriceLast5DaysByCity[];
}

class PriceApi {
  getList(request: PriceListRequest) {
    return baseFetcher.apiGet<PriceListResponse>({
      path: `${apiPrice}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getExist(request: PriceExistRequest) {
    return baseFetcher.apiGet<Record<string, boolean>>({
      path: `${apiPriceExist}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getLatestDateExist(request: PriceLatestDateExistRequest) {
    return baseFetcher.apiGet<string>({
      path: `${apiPriceLatestDateExist}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceLast5DaysByCommodity(request: PriceLast5DaysByCommodityRequest) {
    return baseFetcher.apiGet<PriceLast5DaysByCommodityResponse>({
      path: `${apiPriceLast5DaysByCommodity}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceLast5DaysByCity(request: PriceLast5DaysByCityRequest) {
    return baseFetcher.apiGet<PriceLast5DaysByCityResponse>({
      path: `${apiPriceLast5DaysByCity}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceApi = new PriceApi();
