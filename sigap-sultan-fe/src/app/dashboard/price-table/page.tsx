"use client";

import { PriceTier } from "@/types/price";
import { Card, Grid, Text, Container, Group, Pagination } from "@mantine/core";
import PriceTableContainer, {
  PriceTableType,
} from "@/sections/dashboard/price-table/PriceTableContainer";
import EmptyPage from "@/app/components/EmptyPage";
import {
  PricePerCityCardData,
  PricePerCityDataMap,
} from "@/types/pricePerCityMap";
import PriceTableFilter, {
  PriceTableState,
} from "@/sections/dashboard/price-table/components/PriceTableFilter";
import PriceLast5DaysFilter, {
  PriceLast5DaysState,
} from "@/sections/dashboard/price-table/components/PriceLast5DaysFilter";
import usePriceLast5Days from "@/sections/dashboard/price-table/hooks/usePriceLast5Days";
import { useState } from "react";
import PriceLast5Days from "@/app/components/PriceLast5Day";
import { add, format } from "date-fns";
import LoadingPage from "@/app/components/LoadingPage";

export type PriceState = {
  priceLevel: PricePerCityDataMap;
  priceTier: PriceTier[];
  pricePerCity: PricePerCityCardData[];
};

export default function PriceTable() {
  const [priceTableSubmitted, setPriceTableSubmitted] =
    useState<PriceTableState>({
      commodityType: null,
      priceInfoType: null,
      priceInfoSubType: null,
      date: add(new Date(), { days: -1 }),
      requestTimestamp: new Date().getTime(),
    });
  const [priceLast5DaysSubmitted, setPriceLast5DaysSubmitted] =
    useState<PriceLast5DaysState>({
      filterBy: "komoditas",
      city: null,
      commodityType: null,
      requestTimestamp: new Date().getTime(),
    });
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const {
    isLoading: isLoadingLast5Days,
    dataLast5DaysByCommodity,
    dataLast5DaysByCity,
    pagination,
  } = usePriceLast5Days(priceLast5DaysSubmitted);

  return (
    <Container fluid p="lg" bg={"#F9FAFB"}>
      <PriceTableFilter
        priceTableSubmitted={priceTableSubmitted}
        setPriceTableSubmitted={setPriceTableSubmitted}
        isLoading={isLoading}
      />
      <PriceTableContainer
        priceTableState={priceTableSubmitted}
        priceTableType={
          priceTableSubmitted.priceInfoSubType?.value as PriceTableType
        }
      />
      <PriceLast5DaysFilter
        priceLast5DaysSubmitted={priceLast5DaysSubmitted}
        setPriceLast5DaysSubmitted={setPriceLast5DaysSubmitted}
        isLoading={isLoadingLast5Days}
      />
      <Group gap="md" justify="space-between" mb="md">
        <Text size="md" fw="bold">
          Trend & Perubahan Harga 5 Hari Terakhir
        </Text>
        <Text size="md" c={"#667085"}>
          {`(${format(
            add(priceTableSubmitted.date || new Date(), { days: -4 }),
            "d MMMM yyyy"
          )} - ${format(
            priceTableSubmitted.date || new Date(),
            "d MMMM yyyy"
          )})`}
        </Text>
      </Group>
      {isLoadingLast5Days && <LoadingPage />}
      {dataLast5DaysByCommodity?.cities?.length === 0 &&
        dataLast5DaysByCity?.commodities?.length === 0 && (
          <EmptyPage title="Ups! data yang kamu cari tidak ada, Harap ubah pencarian kamu" />
        )}
      {priceLast5DaysSubmitted.filterBy === "komoditas" && (
        <>
          <Grid>
            {dataLast5DaysByCommodity?.cities?.map((last5DaysData, index) => (
              <Grid.Col span={{ sm: 12, md: 4 }} key={index}>
                <PriceLast5Days last5DaysData={last5DaysData} />
              </Grid.Col>
            ))}
          </Grid>
          <Group justify="center" mt="md">
            <Pagination
              total={dataLast5DaysByCommodity?.totalData ?? 0}
              defaultValue={pagination.active}
            />
          </Group>
        </>
      )}
      {priceLast5DaysSubmitted.filterBy === "daerah" && (
        <>
          <Grid>
            {dataLast5DaysByCity?.commodities?.map((last5DaysData, index) => (
              <Grid.Col span={{ sm: 12, md: 4 }} key={index}>
                <PriceLast5Days last5DaysData={last5DaysData} />
              </Grid.Col>
            ))}
          </Grid>
          <Group justify="center" mt="md">
            <Pagination
              total={dataLast5DaysByCity?.totalData ?? 0}
              defaultValue={pagination.active}
            />
          </Group>
        </>
      )}
    </Container>
  );
}
