import { useState } from "react";
import { Select, Text, Stack, Group, Modal, Box, Button } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { AreaChart } from "@mantine/charts";
import { StockPerCityData } from "@/types/stockPerCityMap";
import { useQuery } from "@tanstack/react-query";
import {
  NeracaLocationSummaryRequest,
  neracaAvailabilityApi,
} from "@/api/base/neraca/neracaAvailability";
import { add, format } from "date-fns";
import { notifications } from "@mantine/notifications";
import { NeracaFilterState } from "@/sections/dashboard/neraca/components/NeracaFilter";
import { LoadingPageContainer } from "../../../../app/components/LoadingPage";

type Props = {
  opened: boolean;
  close: () => void;
  title: string;
  selectedCity: StockPerCityData;
  neracaState: NeracaFilterState;
  unitSuffix: string;
};

function NeracaAvailabilityPerCityChartModal(props: Props) {
  const { title, opened, close, selectedCity, neracaState, unitSuffix } = props;
  const [timeFilter, setTimeFilter] = useState<string | null>("last_3_month");

  const timeFilterOptions: OptionMap<string>[] = [
    { label: "3 Bulan Terakhir", value: "last_3_month" },
    { label: "1 Tahun Terakhir", value: "last_year" },
  ];

  const { data: dataPricePerCity, isFetching: isLoadingPricePerCity } =
    useQuery({
      queryKey: [
        "neraca-map-summary-availability",
        selectedCity.city,
        timeFilter,
        neracaState.commodityType,
        neracaState.requestTimestamp,
      ],
      queryFn: async () => {
        const defaultDate = neracaState.date || new Date();
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
          await neracaAvailabilityApi.getNeracaAvailabilityCityHistory({
            cityId: selectedCity.cityId,
            commodityId: neracaState.commodityType?.value,
            startDate: format(startDate, "yyyy-MM-dd"),
            endDate: format(endDate, "yyyy-MM-dd"),
          } as NeracaLocationSummaryRequest);

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
      <LoadingPageContainer isLoading={isLoadingPricePerCity} height={495}>
        <Box px="md" mb="md">
          <AreaChart
            h={495}
            data={dataPricePerCity?.stock ?? []}
            dataKey="date"
            series={[
              {
                name: "stock",
                color: "#23A65F",
                label: "Stock",
              },
            ]}
            curveType="natural"
            tickLine="none"
            xAxisLabel="Bulan"
            gridAxis="none"
            xAxisProps={{}}
            valueFormatter={(value) => `${value} ${unitSuffix}`}
          />
        </Box>
      </LoadingPageContainer>
      <Button fullWidth onClick={close}>
        Close
      </Button>
    </Modal>
  );
}

export default NeracaAvailabilityPerCityChartModal;
