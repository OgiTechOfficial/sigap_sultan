import { baseFetcher } from "../fetcher";
import { apiCommodities } from "../api-paths";
import { Commodity } from "@/types/commodity";
import constructUrlSearchParams from "@/utils/construct-url-search-params";

export interface CommodityListRequest {
  name?: string;
  moduleType?: "neraca";
}

class CommodityApi {
  getList(request: CommodityListRequest) {
    return baseFetcher.apiGet<Commodity[]>({
      path: `${apiCommodities}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const commodityApi = new CommodityApi();
