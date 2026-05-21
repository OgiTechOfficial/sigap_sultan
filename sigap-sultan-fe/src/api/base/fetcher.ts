import { APIResult } from "@/utils/api/handle-response";
import fetcher, { ApiDeleteProps, ApiGetProps, ApiPatchProps, ApiPostProps, ApiPutProps } from "../fetcher";
import { apiHostConfig } from "@/config";

type BaseApiGetProps = Omit<ApiGetProps, 'baseUrl'>

type BaseApiPostProps = Omit<ApiPostProps, 'baseUrl'>

type BaseApiPatchProps = Omit<ApiPatchProps, 'baseUrl'>

type BaseApiPutProps = Omit<ApiPutProps, 'baseUrl'>

type BaseApiDeleteProps = Omit<ApiDeleteProps, 'baseUrl'>

class BaseFetcher {
  apiGet<T>(props: BaseApiGetProps): Promise<APIResult<T>> {
    return fetcher.apiGet<T>({
      ...props,
      baseUrl: apiHostConfig.baseApiUrl,
    });
  }

  apiPost<T>(props: BaseApiPostProps): Promise<APIResult<T>> {
    return fetcher.apiPost<T>({
      ...props,
      baseUrl: apiHostConfig.baseApiUrl,
    });
  }

  apiPatch<T>(props: BaseApiPatchProps): Promise<APIResult<T>> {
    return fetcher.apiPatch<T>({
      ...props,
      baseUrl: apiHostConfig.baseApiUrl,
    });
  }

  apiPut<T>(props: BaseApiPutProps): Promise<APIResult<T>> {
    return fetcher.apiPut<T>({
      ...props,
      baseUrl: apiHostConfig.baseApiUrl,
    });
  }

  apiDelete<T>(props: BaseApiDeleteProps): Promise<APIResult<T>> {
    return fetcher.apiDelete<T>({
      ...props,
      baseUrl: apiHostConfig.baseApiUrl,
    });
  }
}

export const baseFetcher = new BaseFetcher();