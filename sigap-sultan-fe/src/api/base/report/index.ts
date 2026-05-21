import { baseFetcher } from "../fetcher";
import {
  apiReportNeraca,
  apiReportNeracaDownload,
  apiReportPrice,
  apiReportPriceDownload,
} from "../api-paths";
import { ReportNeraca, ReportPrice } from "@/types/report";
import constructUrlSearchParams from "@/utils/construct-url-search-params";

export interface ReportNeracaRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

export interface ReportPriceRequest {
  page?: number;
  limit?: number;
  cityId?: number;
  commodityId?: number;
  startDate?: string;
  endDate?: string;
}

class ReportApi {
  getReportNeraca(request: ReportNeracaRequest) {
    return baseFetcher.apiGet<ReportNeraca>({
      path: `${apiReportNeraca}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  downloadReportNeraca(request: ReportNeracaRequest) {
    return baseFetcher.apiGet<Blob>({
      path: `${apiReportNeracaDownload}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  getReportPrice(request: ReportPriceRequest) {
    return baseFetcher.apiGet<ReportPrice>({
      path: `${apiReportPrice}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
  downloadReportPrice(request: ReportPriceRequest) {
    return baseFetcher.apiGet<Blob>({
      path: `${apiReportPriceDownload}`,
      config: {
        params: constructUrlSearchParams(request),
      },
    });
  }
}

export const reportApi = new ReportApi();
