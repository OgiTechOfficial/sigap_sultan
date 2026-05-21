import { baseFetcher } from "../fetcher";
import { apiPosition } from "../api-paths";
import { Position } from "@/types/position";

class PositionApi {
  getList() {
    return baseFetcher.apiGet<Position[]>({
      path: `${apiPosition}`,
    });
  }
}

export const positionApi = new PositionApi();
