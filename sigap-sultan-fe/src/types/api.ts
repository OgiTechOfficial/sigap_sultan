
export interface ResultData<T> {
  contents: T;
  currentPage: number;
  totalPages: number;
  totalElements: number;
  hasNext: boolean;
  hasPrevious: boolean;
  firstPage: boolean;
  lastPage: boolean;
}

export interface ResultSubmit {
  status: boolean;
  rows_affected: number;
  last_insert_id: number;
}

export interface DefaultResponse {
  status: number;
  message: string;
  errors: string;
  serverTime: string;
}

export interface ErrorResponse {
  error: string;
  error_description: string;
}

export interface APIResponseList<T> extends DefaultResponse {
  data: ResultData<T>;
}

export interface APIResponse<T> extends DefaultResponse {
  data: T;
}

export interface DefaultRequest {
  client_id?: string;
  client_secret?: string;
}

export interface DefaultQueryRequest extends DefaultRequest {
  search?: string;
}

export interface DefaultQueryPageRequest extends DefaultQueryRequest {
  page: number;
  limit: number;
  sortDirection: 'ASC' | 'DESC';
}

export interface PaginationData {
  page: number;
  totalData: number;
  totalPage: number;
}

export interface ResponseData {
  id: string;
  createdDate: string;
  modifiedDate: string;
}

export interface DefaultResponseData<T> {
  code: 'OK' | string;
  message: string;
  data: T;
  errors: string | null;
  serverTime: number;
}

export interface DefaultResponseDataPagination<T> {
  code: 'OK' | string;
  message: string;
  data: T;
  errors: string | null;
  serverTime: number;
  currentPage: number;
  totalElements: number;
  totalPages: number;
}