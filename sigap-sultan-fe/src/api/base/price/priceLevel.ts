import { baseFetcher } from "../fetcher";
import { apiPriceLevelList, apiPriceLevelMapNew } from "../api-paths";
import {
  PriceLevelMap,
  PriceLevelList,
  PriceTier,
  PriceLevelProvince,
  PriceLevelMapProvince,
} from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { Asset } from "@/types/asset";

export type PriceDiffBy = "province" | "national";

export interface PriceListRequest {
  commodityId?: number;
  cityId: number;
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

export interface PriceLevelMapRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  commodityId?: number;
  selectedDate?: string;
}

export interface PriceLevelMapResponse {
  provincePrice: PriceLevelMapProvince;
  cityPrice: PriceLevelMap[];
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

export interface PriceLevelListResponse {
  commodity: {
    id: number;
    name: string;
  };
  priceProvince: PriceLevelProvince | null;
  priceCity: PriceLevelList[] | null;
}

class PriceLevelApi {
  getPriceLevelMap(request: PriceLevelMapRequest) {
    return baseFetcher.apiGet<PriceLevelMapResponse>({
      path: `${apiPriceLevelMapNew}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getPriceLevelList(request: PriceLevelListRequest) {
    return baseFetcher.apiGet<PriceLevelListResponse>({
      path: `${apiPriceLevelList}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceLevelApi = new PriceLevelApi();
