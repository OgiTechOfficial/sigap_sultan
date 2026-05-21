import { priceApi } from "@/api/base/price";
import { PriceLast5DaysState } from "@/sections/dashboard/price-table/components/PriceLast5DaysFilter";
import { usePagination } from "@mantine/hooks";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";

const PAGE_LIMIT = 9;

function usePriceLast5Days(priceLast5DaysState: PriceLast5DaysState) {
  const [page, onChange] = useState(1);
  const pagination = usePagination({ total: PAGE_LIMIT, page, onChange });

  const {
    data: dataLast5DaysByCommodity,
    isLoading: isLoadingLast5DaysByCommodity,
  } = useQuery({
    queryKey: [
      "price-last-5-days-by-commodity",
      priceLast5DaysState.commodityType,
      pagination.active,
      priceLast5DaysState.requestTimestamp,
    ],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await priceApi.getPriceLast5DaysByCommodity({
          commodityId: priceLast5DaysState.commodityType?.value ?? 0,
          page: pagination.active,
          limit: PAGE_LIMIT,
        });

      if (error || !result) {
        throw new Error(displayMessage ?? "Failed to fetch price list");
      }

      return result;
    },
    enabled: !!priceLast5DaysState.commodityType,
  });

  const { data: dataLast5DaysByCity, isLoading: isLoadingLast5DaysByCity } =
    useQuery({
      queryKey: [
        "price-last-5-days-by-city",
        priceLast5DaysState.city,
        pagination.active,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceApi.getPriceLast5DaysByCity({
            cityId: priceLast5DaysState.city
              ? Number(priceLast5DaysState.city.value)
              : 0,
            page: pagination.active,
            limit: PAGE_LIMIT,
          });

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch price list");
        }

        return result;
      },
      enabled: !!priceLast5DaysState.city,
    });

  return {
    isLoading: isLoadingLast5DaysByCommodity || isLoadingLast5DaysByCity,
    dataLast5DaysByCommodity,
    dataLast5DaysByCity,
    pagination,
  };
}

export default usePriceLast5Days;
