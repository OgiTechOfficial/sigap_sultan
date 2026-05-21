import { priceLevelApi } from "@/api/base/price/priceLevel";
import { PriceTableState } from "@/sections/dashboard/price-table/components/PriceTableFilter";
import {
  PricePerCityCardData,
  PricePerCityDataMap,
} from "@/types/pricePerCityMap";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";

function usePriceLevel(
  priceTableState: PriceTableState,
  sortBy: string | null
) {
  const { data: dataPriceLevelMap, isFetching: isFetchingPriceLevelMap } =
    useQuery({
      queryKey: [
        "price-level-map",
        priceTableState.commodityType,
        priceTableState.date,
        priceTableState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceLevelApi.getPriceLevelMap({
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
            priceLevel: result.cityPrice.reduce(
              (acc: PricePerCityDataMap, price) => {
                acc[price.city.name] = {
                  cityId: price.city.id,
                  city: price.city.name,
                  price: price.price,
                };
                return acc;
              },
              {}
            ),
            provincePrice: result.provincePrice,
            priceTier: result.priceTier ?? [],
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
        "price-level-list",
        priceTableState.commodityType,
        priceTableState.date,
        sortBy,
        priceTableState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceLevelApi.getPriceLevelList({
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
          throw new Error(displayMessage ?? "Failed to fetch price level list");
        }

        try {
          const pricePerCity = (result.priceCity || []).map(
            (price, priceIndex) => ({
              cityId: price.city.id,
              cityNumber: priceIndex + 1,
              city: price.city.name,
              cityImage: price.city.assets?.assetsUrl,
              price: price.price,
              priceTier: price.priceTier[price.tier]?.title,
              priceTierColor: price.priceTier[price.tier]?.color,
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
          if (result.priceProvince) {
            pricePerProvince = {
              cityId: result.priceProvince.province.id,
              city: result.priceProvince.province.name,
              cityImage: result.priceProvince.province.assets?.assetsUrl,
              price: result.priceProvince.price,
              priceTier:
                result.priceCity?.[0].priceTier[result.priceProvince.tier]
                  ?.title,
              priceTierColor:
                result.priceCity?.[0].priceTier[result.priceProvince.tier]
                  ?.color,
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
    dataPriceLevelMap,
    dataPriceLevelList,
  };
}

export default usePriceLevel;
