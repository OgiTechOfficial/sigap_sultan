import { useState } from "react";
import { Select, Text, Stack, Group, Modal, Box, Button } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { AreaChart } from "@mantine/charts";
import { PricePerCityData } from "@/types/pricePerCityMap";
import { useQuery } from "@tanstack/react-query";
import { notifications } from "@mantine/notifications";
import { add, format } from "date-fns";
import {
  priceMTMApi,
  PriceMTMCityHistoryRequest,
} from "@/api/base/price/priceMTM";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { PriceTableState } from "./PriceTableFilter";
import { RegionType } from "@/types/city";

type Props = {
  opened: boolean;
  close: () => void;
  title: string;
  selectedCity: PricePerCityData;
  priceTableState: PriceTableState;
  regionType?: RegionType;
};

function MTMPerCityChartModal(props: Props) {
  const {
    title,
    opened,
    close,
    selectedCity,
    priceTableState,
    regionType = "CITY",
  } = props;
  const [timeFilter, setTimeFilter] = useState<string | null>("last_week");

  const timeFilterOptions: OptionMap<string>[] = [
    { label: "1 Minggu Terakhir", value: "last_week" },
    { label: "1 Bulan Terakhir", value: "last_month" },
    { label: "3 Bulan Terakhir", value: "last_3_month" },
    { label: "1 Tahun Terakhir", value: "last_year" },
  ];

  const { data: dataPricePerCity, isFetching: isLoadingPricePerCity } =
    useQuery({
      queryKey: [
        "price-per-city-chart-mtm",
        selectedCity.city,
        priceTableState.commodityType,
        priceTableState.date,
        timeFilter,
        priceTableState.requestTimestamp,
        regionType,
      ],
      queryFn: async () => {
        const defaultDate =
          priceTableState.date ?? add(new Date(), { days: -1 });
        let startDate: Date = defaultDate;
        let endDate: Date = defaultDate;
        if (timeFilter === "last_week") {
          startDate = add(defaultDate, { days: -7 });
          endDate = defaultDate;
        } else if (timeFilter === "last_month") {
          startDate = add(defaultDate, { months: -1 });
          endDate = defaultDate;
        } else if (timeFilter === "last_3_month") {
          startDate = add(defaultDate, { months: -3 });
          endDate = defaultDate;
        } else if (timeFilter === "last_year") {
          startDate = add(defaultDate, { years: -1 });
          endDate = defaultDate;
        }
        const { result, error, displayMessage } =
          await priceMTMApi.getPriceMTMCityHistory({
            cityId: regionType === "CITY" && selectedCity.cityId,
            provinceId: regionType === "PROVINCE" && selectedCity.cityId,
            commodityId: priceTableState.commodityType?.value,
            startDate: format(startDate, "yyyy-MM-dd"),
            endDate: format(endDate, "yyyy-MM-dd"),
          } as PriceMTMCityHistoryRequest);

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch price list");
        }

        return result;
      },
    });

  return (
    <Modal
      size={1178}
      opened={opened}
      onClose={close}
      withCloseButton={false}
      centered
    >
      <Group justify="space-between" gap="md" mb="xl">
        <Stack gap={0}>
          <Text size="lg">{selectedCity.city}</Text>
          <Text size="md">{title}</Text>
        </Stack>
        <Select
          data={timeFilterOptions}
          value={timeFilter}
          onChange={(value) => setTimeFilter(value)}
        />
      </Group>
      <LoadingPageContainer isLoading={isLoadingPricePerCity}>
        <Box px="md" mb="md">
          <AreaChart
            h={495}
            data={dataPricePerCity?.commodityInflations ?? []}
            dataKey="date"
            series={[
              {
                name: "inflation",
                color: "#667085",
                label: "Harga",
              },
            ]}
            curveType="natural"
            tickLine="none"
            xAxisLabel="Bulan"
            gridAxis="none"
            xAxisProps={{}}
            valueFormatter={(value) => `${value} %`}
          />
        </Box>
      </LoadingPageContainer>
      <Button fullWidth onClick={close}>
        Close
      </Button>
    </Modal>
  );
}

export default MTMPerCityChartModal;
