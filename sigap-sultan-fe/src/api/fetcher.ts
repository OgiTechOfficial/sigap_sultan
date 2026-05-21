import axios, { AxiosRequestConfig } from "axios";
import handleResponse, { APIResult } from "@/utils/api/handle-response";
import handleInterceptorAPI from "@/utils/api/handle-interceptor-api";

export type ApiGetProps = {
  path: string,
  baseUrl: string,
  config?: AxiosRequestConfig,
}

export type ApiPostProps = {
  path: string,
  baseUrl: string,
  payload?: Object,
  config?: AxiosRequestConfig,
}

export type ApiPatchProps = {
  path: string,
  baseUrl: string,
  payload?: Object,
  config?: AxiosRequestConfig,
}

export type ApiPutProps = {
  path: string,
  baseUrl: string,
  payload?: Object,
  config?: AxiosRequestConfig,
}

export type ApiDeleteProps = {
  path: string,
  baseUrl: string,
  config?: AxiosRequestConfig,
}

class Fetcher {
  apiGet<T>(props: ApiGetProps): Promise<APIResult<T>> {
    const { path, baseUrl, config } = props;
    const url = baseUrl + path;
    const instance = handleInterceptorAPI(axios);

    return handleResponse<T>(
      instance.get<APIResult<T>>(url, config)
    );
  }

  apiPost<T>(props: ApiPostProps): Promise<APIResult<T>> {
    const { path, baseUrl, payload, config } = props;
    const url = baseUrl + path;
    const instance = handleInterceptorAPI(axios);

    return handleResponse<T>(
      instance.post<APIResult<T>>(url, payload, config)
    );
  }

  apiPatch<T>(props: ApiPatchProps): Promise<APIResult<T>> {
    const { path, baseUrl, payload, config } = props;
    const url = baseUrl + path;
    const instance = handleInterceptorAPI(axios);

    return handleResponse<T>(
      instance.patch<APIResult<T>>(url, payload, config)
    );
  }

  apiPut<T>(props: ApiPutProps): Promise<APIResult<T>> {
    const { path, baseUrl, payload, config } = props;
    const url = baseUrl + path;
    const instance = handleInterceptorAPI(axios);

    return handleResponse<T>(
      instance.put<APIResult<T>>(url, payload, config)
    );
  }

  apiDelete<T>(props: ApiDeleteProps): Promise<APIResult<T>> {
    const { path, baseUrl, config } = props;
    const url = `${baseUrl}` + path;
    const instance = handleInterceptorAPI(axios);

    return handleResponse<T>(
      instance.delete<APIResult<T>>(url, config)
    );
  }
}

const fetcher = new Fetcher();

export default fetcher;