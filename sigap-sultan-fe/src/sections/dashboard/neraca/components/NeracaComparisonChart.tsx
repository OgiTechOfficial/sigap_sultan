"use client";

import { Button, Card, Group, rem, Text, Tooltip } from "@mantine/core";
import { AreaChart } from "@mantine/charts";
import { add, format } from "date-fns";
import { NeracaDetailFilterState } from "./NeracaDetailFilter";
import { useState } from "react";
import { OptionMap } from "@/types/option";
import { useQuery } from "@tanstack/react-query";
import { neracaApi } from "@/api/base/neraca";
import { notifications } from "@mantine/notifications";
import { useMediaQuery } from "@mantine/hooks";
import { IconInfoCircle } from "@tabler/icons-react";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { FormatNumber } from "@/utils/currency";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";

type NeracaSummaryFilter = {
  timeFilter: "last_3_month" | "last_6_month" | "last_year";
};

type Props = {
  neracaDetailFilterSubmitted: NeracaDetailFilterState;
};

function NeracaComparisonChart(props: Props) {
  const { neracaDetailFilterSubmitted } = props;
  const [neracaSummaryFilter, setNeracaSummaryFilter] =
    useState<NeracaSummaryFilter>({
      timeFilter: "last_year",
    });
  const timeFilterOptions: OptionMap<string>[] = [
    { label: "1 Tahun", value: "last_year" },
  ];
  const matches = useMediaQuery("(min-width: 56.25em)");
  const { commodityUnitMap } = useCommodityOptions(
    () => {},
    "neraca-filter",
    "neraca"
  );
  const unitSuffix = neracaDetailFilterSubmitted.commodityType?.value
    ? commodityUnitMap[neracaDetailFilterSubmitted.commodityType?.value]
    : "";

  const {
    data: dataNeracaComparisonChart,
    isFetching: isLoadingNeracaComparisonChart,
  } = useQuery({
    queryKey: [
      "neraca-comparison-chart",
      neracaSummaryFilter.timeFilter,
      neracaDetailFilterSubmitted.city,
      neracaDetailFilterSubmitted.commodityType,
      neracaDetailFilterSubmitted.requestTimestamp,
    ],
    queryFn: async () => {
      const defaultDate = neracaDetailFilterSubmitted.date || new Date();
      let startDate: Date = defaultDate;
      let endDate: Date = defaultDate;
      if (neracaSummaryFilter.timeFilter === "last_3_month") {
        startDate = add(defaultDate, { months: -2 });
        endDate = defaultDate;
      } else if (neracaSummaryFilter.timeFilter === "last_6_month") {
        startDate = add(defaultDate, { months: -5 });
        endDate = defaultDate;
      } else if (neracaSummaryFilter.timeFilter === "last_year") {
        startDate = add(defaultDate, { years: -1 });
        endDate = defaultDate;
      }

      const { result, error, displayMessage } =
        await neracaApi.getNeracaCompareWithPriceAndCommodityHistory({
          page: 1,
          limit: 100,
          cityId: neracaDetailFilterSubmitted.city
            ? Number(neracaDetailFilterSubmitted.city.value)
            : undefined,
          commodityId: neracaDetailFilterSubmitted.commodityType
            ? neracaDetailFilterSubmitted.commodityType.value
            : undefined,
          startDate: format(startDate, "yyyy-MM-dd"),
          endDate: format(endDate, "yyyy-MM-dd"),
        });

      if (error || !result) {
        throw new Error(
          displayMessage ?? "Failed to fetch neraca comparison chart"
        );
      }

      return result;
    },
    enabled:
      !!neracaDetailFilterSubmitted.city &&
      !!neracaDetailFilterSubmitted.commodityType &&
      !!neracaSummaryFilter.timeFilter,
  });

  const chartSeries = [
    {
      name: "stock",
      label: `Neraca (${unitSuffix})`,
      color: "#23A65F",
    },
    {
      name: "price",
      label: `Harga(Rp/${unitSuffix}) - Rhs`,
      color: "#DC4531",
      yAxisId: "right",
    },
  ];

  return (
    <Card padding="md" radius="md" withBorder>
      <Text size="md" fw="bold" mb="md">
        {neracaDetailFilterSubmitted.commodityType?.label} {`/ ${unitSuffix}`} -{" "}
        {neracaDetailFilterSubmitted.city?.label}
      </Text>
      <Group gap="md" justify="flex-end" mb="md">
        <Button.Group>
          {timeFilterOptions.map((option, index) => {
            const selected = neracaSummaryFilter.timeFilter === option.value;
            return (
              <Button
                variant="default"
                key={index}
                bg={selected ? "#EAECF0" : "#FFFFFF"}
                onClick={() =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    timeFilter:
                      option.value as NeracaSummaryFilter["timeFilter"],
                  }))
                }
              >
                <Group gap="xs">
                  <Text size="sm">{option.label}</Text>
                  {matches && (
                    <Tooltip
                      multiline
                      w={220}
                      withArrow
                      transitionProps={{ duration: 200 }}
                      label={`${option.label} sebelumnya dari tanggal saat ini`}
                    >
                      <IconInfoCircle
                        style={{ width: rem(15), height: rem(15) }}
                        color={"#344054"}
                      />
                    </Tooltip>
                  )}
                </Group>
              </Button>
            );
          })}
        </Button.Group>
      </Group>
      <LoadingPageContainer
        isLoading={isLoadingNeracaComparisonChart}
        height={240}
      >
        <AreaChart
          h={240}
          data={dataNeracaComparisonChart?.stock ?? []}
          dataKey="date"
          withRightYAxis
          yAxisLabel={`Neraca (${unitSuffix})`}
          yAxisProps={{
            width: 100,
          }}
          rightYAxisLabel={`Harga (Rp/${unitSuffix})`}
          rightYAxisProps={{
            width: 100,
          }}
          series={chartSeries}
          curveType="natural"
          tickLine="none"
          gridAxis="none"
          withLegend
          valueFormatter={(value) => `${FormatNumber(value)}`}
        />
      </LoadingPageContainer>
    </Card>
  );
}

export default NeracaComparisonChart;
