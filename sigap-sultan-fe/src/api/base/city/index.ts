import { baseFetcher } from '../fetcher';
import { apiCity } from '../api-paths';
import { City } from '@/types/city';
import constructUrlSearchParams from '@/utils/construct-url-search-params';

export interface CityListRequest {
  provinceId?: number;
}

class CityApi {
  getList(request: CityListRequest) {
    return baseFetcher.apiGet<City[]>({
      path: `${apiCity}`,
      config: {
        params: constructUrlSearchParams(request),
      }
    });
  }
}

export const cityApi = new CityApi();
