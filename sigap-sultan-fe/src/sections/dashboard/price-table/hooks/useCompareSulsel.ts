import { priceCompareProvinceApi } from "@/api/base/price/priceCompareProvince";
import { dataMapEmpty } from "@/constants/map/dataMap";
import { PriceTableState } from "@/sections/dashboard/price-table/components/PriceTableFilter";
import {
  PricePerCityCardData,
  PricePerCityDataMap,
} from "@/types/pricePerCityMap";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";

function useCompareSulsel(
  priceTableState: PriceTableState,
  sortBy: string | null
) {
  const { data: dataPriceLevelMap, isFetching: isFetchingPriceLevelMap } =
    useQuery({
      queryKey: [
        "price-compare-province-map",
        priceTableState.commodityType,
        priceTableState.date,
        priceTableState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceCompareProvinceApi.getPriceCompareProvinceMap({
            page: 1,
            limit: 100,
            sortBy: "price:asc",
            commodityId: priceTableState.commodityType
              ? priceTableState.commodityType.value
              : undefined,
            selectedDate: priceTableState.date
              ? format(priceTableState.date, "yyyy-MM-dd")
              : "",
          });

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch price level map");
        }

        try {
          return {
            priceLevel: result.priceLevel.reduce(
              (acc: PricePerCityDataMap, price) => {
                acc[price.city.name] = {
                  cityId: price.city.id,
                  city: price.city.name,
                  price: price.price,
                  color:
                    result.priceTier[price.tier]?.color ?? dataMapEmpty.color,
                };
                return acc;
              },
              {}
            ),
            priceTier: result.priceTierCode,
            priceTierMap: result.priceTier,
            summary: result.summary,
          };
        } catch (err) {
          console.error(err);
        }
      },
      enabled: !!priceTableState.commodityType && !!priceTableState.date,
    });

  const { data: dataPriceLevelList, isFetching: isFetchingPriceLevelList } =
    useQuery({
      queryKey: [
        "price-compare-province-list",
        priceTableState.commodityType,
        priceTableState.date,
        sortBy,
        priceTableState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceCompareProvinceApi.getPriceCompareProvinceList({
            page: 1,
            limit: 100,
            sortBy: sortBy ? `price:${sortBy}` : "",
            commodityId: priceTableState.commodityType
              ? priceTableState.commodityType.value
              : undefined,
            selectedDate: priceTableState.date
              ? format(priceTableState.date, "yyyy-MM-dd")
              : "",
          });

        if (error || !result) {
          throw new Error(
            displayMessage ?? "Failed to fetch compare sulsel list"
          );
        }

        try {
          const pricePerCity = (result.priceCity || []).map(
            (price, priceIndex) => ({
              cityId: price.city.id,
              cityNumber: priceIndex + 1,
              city: price.city.name,
              cityImage: price.city.assets?.assetsUrl,
              price: price.price,
              priceDiff: price.priceDiff,
              priceType: "ABSOLUTE",
            })
          ) as PricePerCityCardData[];
          const totalPage = 4;
          const totalElement = pricePerCity.length;
          const itemPerPage = Math.round(totalElement / totalPage);
          let pricePerCityPaginated: PricePerCityCardData[][] = [];
          Array.from({ length: totalPage }).map((_, pageIndex) => {
            const firstOffset = pageIndex * itemPerPage;
            const pricePerCityTemp = [...pricePerCity].splice(
              firstOffset,
              itemPerPage
            );
            pricePerCityPaginated.push(pricePerCityTemp);
          });
          let pricePerProvince: PricePerCityCardData | null = null;
          if (result.provincePrice) {
            pricePerProvince = {
              cityId: result.provincePrice.province.id,
              city: result.provincePrice.province.name,
              cityImage: result.provincePrice.province.assets?.assetsUrl,
              price: result.provincePrice.price,
              priceType: "ABSOLUTE",
            };
          }
          return {
            pricePerCity,
            pricePerCityPaginated,
            pricePerProvince,
          };
        } catch (err) {
          console.error(err);
        }
      },
      enabled: !!priceTableState.commodityType && !!priceTableState.date,
    });

  return {
    isLoadingPriceLevelMap: isFetchingPriceLevelMap,
    isLoadingPriceLevelList: isFetchingPriceLevelList,
    dataPriceLevelList,
    dataPriceLevelMap,
  };
}

export default useCompareSulsel;
