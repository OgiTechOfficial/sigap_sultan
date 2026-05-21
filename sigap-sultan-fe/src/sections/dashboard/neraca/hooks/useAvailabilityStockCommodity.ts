import { neracaAvailabilityApi } from "@/api/base/neraca/neracaAvailability";
import { NeracaFilterState } from "../components/NeracaFilter";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import {
  StockPerCityCardData,
  StockPerCityDataMap,
} from "@/types/stockPerCityMap";

function useAvailabilityStockCommodity(
  neracaState: NeracaFilterState,
  sortBy: string | null
) {
  const {
    data: dataAvailabilityStockCommodityMap,
    isFetching: isLoadingAvailabilityStockCommodityMap,
  } = useQuery({
    queryKey: [
      "availability-stock-commodity-map",
      neracaState.commodityType,
      neracaState.date,
      neracaState.requestTimestamp,
    ],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await neracaAvailabilityApi.getNeracaAvailabilityStockByCommodityMap({
          page: 1,
          limit: 100,
          sortBy: "id:asc",
          commodityId: neracaState.commodityType
            ? neracaState.commodityType.value
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
          cityStock: (result.cityStock || []).reduce(
            (acc: StockPerCityDataMap, price) => {
              acc[price.city.name] = {
                cityId: price.city.id,
                city: price.city.name,
                stock: price.stock,
                tier: price.tier,
              };
              return acc;
            },
            {}
          ),
          provinceStock: result.provinceStock,
          stockTier: result.stockTier,
          stockTierCode: result.stockTierCode,
          summary: result.summary,
        };
      } catch (err) {
        console.error(err);
      }
    },
    enabled: !!neracaState.commodityType && !!neracaState.date,
  });

  const {
    data: dataAvailabilityStockCommodityList,
    isFetching: isLoadingAvailabilityStockCommodityList,
  } = useQuery({
    queryKey: [
      "availability-stock-commodity-list",
      neracaState.commodityType,
      neracaState.date,
      sortBy,
      neracaState.requestTimestamp,
    ],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await neracaAvailabilityApi.getNeracaAvailabilityStockByCommodityList({
          page: 1,
          limit: 100,
          sortBy: sortBy ? `neraca:${sortBy}` : "",
          commodityId: neracaState.commodityType
            ? neracaState.commodityType.value
            : undefined,
          selectedDate: neracaState.date
            ? format(neracaState.date, "yyyy-MM-dd")
            : "",
        });

      if (error || !result) {
        throw new Error(
          displayMessage ?? "Failed to fetch availability stock commodity"
        );
      }

      try {
        const stockPerCity = (result.cityStock || []).map(
          (stock, stockIndex) => ({
            cityId: stock.city.id,
            cityNumber: stockIndex + 1,
            city: stock.city.name,
            cityImage: stock.city.assets?.assetsUrl,
            stock: stock.stock,
            stockDiff: stock.stockDiff,
            tier: stock.tier,
            stockType: "ABSOLUTE",
          })
        ) as StockPerCityCardData[];
        const totalPage = 4;
        const totalElement = stockPerCity.length;
        const itemPerPage = Math.round(totalElement / totalPage);
        let stockPerCityPaginated: StockPerCityCardData[][] = [];
        Array.from({ length: totalPage }).map((_, pageIndex) => {
          const firstOffset = pageIndex * itemPerPage;
          const stockPerCityTemp = [...stockPerCity].splice(
            firstOffset,
            itemPerPage
          );
          stockPerCityPaginated.push(stockPerCityTemp);
        });
        return {
          stockPerCity,
          stockPerCityPaginated,
          stockPerProvince: {
            cityId: result.provinceStock.province.id,
            city: result.provinceStock.province.name,
            cityImage: result.provinceStock.province.assets?.assetsUrl,
            stock: result.provinceStock.stock,
            stockDiff: result.provinceStock.stockDiff,
            tier: result.provinceStock.tier,
            stockType: "ABSOLUTE",
          } as StockPerCityCardData,
          stockTier: result.stockTier,
        };
      } catch (err) {
        console.error(err);
      }
    },
    enabled: !!neracaState.commodityType && !!neracaState.date,
  });

  return {
    isLoadingAvailabilityStockCommodityList,
    isLoadingAvailabilityStockCommodityMap,
    dataAvailabilityStockCommodityMap,
    dataAvailabilityStockCommodityList,
  };
}

export default useAvailabilityStockCommodity;
