import { baseFetcher } from "../fetcher";
import { apiPriceCommodity } from "../api-paths";
import { PriceCommodityCity } from "@/types/price";
import constructUrlSearchParams from "@/utils/construct-url-search-params";

export interface PriceCommodityListRequest {
  commodityId?: string | null;
  tier?: string | null;
}

export interface PriceCommodityListResponse {
  commodity: {
    id: number;
    name: string;
  };
  cityPrice: PriceCommodityCity[] | null;
}

class PriceCommodityApi {
  getList(request: PriceCommodityListRequest) {
    return baseFetcher.apiGet<PriceCommodityListResponse>({
      path: `${apiPriceCommodity}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const priceCommodityApi = new PriceCommodityApi();
