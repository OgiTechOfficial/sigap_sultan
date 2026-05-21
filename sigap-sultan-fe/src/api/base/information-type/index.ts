import { baseFetcher } from '../fetcher';
import { apiInformationType } from '../api-paths';
import { InformationType } from '@/types/informationType';
import constructUrlSearchParams from '@/utils/construct-url-search-params';

export interface InformationTypeListRequest {
  page?: number;
  limit?: number;
}

class InformationTypeApi {
  getList(request: InformationTypeListRequest) {
    return baseFetcher.apiGet<InformationType[]>({
      path: `${apiInformationType}`,
      config: {
        params: constructUrlSearchParams(request),
      }
    });
  }
}

export const informationTypeApi = new InformationTypeApi();
