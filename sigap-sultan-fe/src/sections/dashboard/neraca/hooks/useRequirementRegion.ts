import { neracaRequirementApi } from "@/api/base/neraca/neracaRequirement";
import { NeracaFilterState } from "../components/NeracaFilter";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";

function useRequirementRegion(neracaState: NeracaFilterState) {
  const { data: dataLastStockCityMap, isFetching: isLoadingLastStockCityMap } =
    useQuery({
      queryKey: [
        "region-requirement-city-map",
        neracaState.commodityType,
        neracaState.date,
        neracaState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await neracaRequirementApi.getNeracaRequirementByCityMap({
            page: 1,
            limit: 100,
            cityId: neracaState.city
              ? Number(neracaState.city.value)
              : undefined,
            selectedDate: neracaState.date
              ? format(neracaState.date, "yyyy-MM-dd")
              : "",
          });

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch price level map");
        }

        try {
          return {
            commodityStock: result.commodityStock,
            provinceStock: result.provinceStock,
            stockTier: result.stockTier,
            stockTierCode: result.stockTierCode,
            summary: result.summary,
          };
        } catch (err) {
          console.error(err);
        }
      },
      enabled: !!neracaState.city && !!neracaState.date,
    });

  return {
    isLoading: isLoadingLastStockCityMap,
    dataLastStockCityMap,
  };
}

export default useRequirementRegion;
