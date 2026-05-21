"use client";

import { useMemo, useState } from "react";
import {
  Button,
  Card,
  Text,
  rem,
  Group,
  Switch,
  Checkbox,
  Tooltip,
  Radio,
} from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconInfoCircle } from "@tabler/icons-react";
import { AreaChart } from "@mantine/charts";
import { neracaApi } from "@/api/base/neraca";
import { useQuery } from "@tanstack/react-query";
import { add, format } from "date-fns";
import { notifications } from "@mantine/notifications";
import { NeracaDetailFilterState } from "./NeracaDetailFilter";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { useMediaQuery } from "@mantine/hooks";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { FormatNumber } from "@/utils/currency";

type NeracaSummaryFilter = {
  isActiveStackedType: boolean;
  isCheckedNeraca: boolean;
  isCheckedAvailability: boolean;
  isCheckedRequired: boolean;
  timeFilter: "last_3_month" | "last_6_month" | "last_year";
};

type Props = {
  neracaDetailFilterSubmitted: NeracaDetailFilterState;
};

function NeracaSummaryChart(props: Props) {
  const { neracaDetailFilterSubmitted } = props;
  const [neracaSummaryFilter, setNeracaSummaryFilter] =
    useState<NeracaSummaryFilter>({
      isActiveStackedType: true,
      isCheckedNeraca: true,
      isCheckedAvailability: true,
      isCheckedRequired: true,
      timeFilter: "last_3_month",
    });
  const timeFilterOptions: OptionMap<string>[] = [
    { label: "3 Bulan", value: "last_3_month" },
    { label: "6 Bulan", value: "last_6_month" },
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
    data: dataNeracaAvailabilityChart,
    isFetching: isLoadingNeracaAvailabilityChart,
  } = useQuery({
    queryKey: [
      "neraca-availability-chart",
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
        await neracaApi.getNeracaAvailabilityByCityAndCommodityChart({
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
          displayMessage ?? "Failed to fetch neraca available chart"
        );
      }

      return result;
    },
    enabled:
      !!neracaDetailFilterSubmitted.city &&
      !!neracaDetailFilterSubmitted.commodityType &&
      !!neracaSummaryFilter.timeFilter,
  });

  const chartSeries = useMemo(() => {
    const series = [];
    if (neracaSummaryFilter.isCheckedNeraca) {
      series.push({
        name: "neraca",
        label: "Neraca",
        color: "#005395",
      });
    }

    if (neracaSummaryFilter.isCheckedAvailability) {
      series.push({
        name: "ketersediaan",
        label: "Ketersediaan",
        color: "#23A65F",
      });
    }

    if (neracaSummaryFilter.isCheckedRequired) {
      series.push({
        name: "kebutuhan",
        label: "Kebutuhan",
        color: "#B11016",
      });
    }

    return series;
  }, [neracaSummaryFilter]);

  return (
    <Card padding="md" radius="md" withBorder>
      <Text size="md" fw="bold" mb="md">
        {neracaDetailFilterSubmitted.commodityType?.label} {`/ ${unitSuffix}`} -{" "}
        {neracaDetailFilterSubmitted.city?.label}
      </Text>
      <Group gap="md" justify="space-between" mb="md">
        <Group gap="md">
          <Switch
            checked={neracaSummaryFilter.isActiveStackedType}
            onChange={(event) => {
              if (event.target.checked) {
                setNeracaSummaryFilter((oldForm) => ({
                  ...oldForm,
                  isActiveStackedType: event.target.checked,
                  isCheckedNeraca: true,
                  isCheckedAvailability: true,
                  isCheckedRequired: true,
                }));
              } else {
                setNeracaSummaryFilter((oldForm) => ({
                  ...oldForm,
                  isActiveStackedType: event.target.checked,
                  isCheckedNeraca: true,
                  isCheckedAvailability: false,
                  isCheckedRequired: false,
                }));
              }
            }}
            label="Tipe Stacked"
          />
          <Card withBorder p="xs">
            {!neracaSummaryFilter.isActiveStackedType && (
              <Radio
                checked={neracaSummaryFilter.isCheckedNeraca}
                onChange={() =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedNeraca: true,
                    isCheckedAvailability: false,
                    isCheckedRequired: false,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#005395"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Neraca</Text>
                  </Group>
                }
              />
            )}
            {neracaSummaryFilter.isActiveStackedType && (
              <Checkbox
                checked={neracaSummaryFilter.isCheckedNeraca}
                onChange={(event) =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedNeraca: event.target.checked,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#005395"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Neraca</Text>
                  </Group>
                }
              />
            )}
          </Card>
          <Card withBorder p="xs">
            {!neracaSummaryFilter.isActiveStackedType && (
              <Radio
                checked={neracaSummaryFilter.isCheckedAvailability}
                onChange={() =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedAvailability: true,
                    isCheckedNeraca: false,
                    isCheckedRequired: false,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#23A65F"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Ketersediaan</Text>
                  </Group>
                }
              />
            )}
            {neracaSummaryFilter.isActiveStackedType && (
              <Checkbox
                checked={neracaSummaryFilter.isCheckedAvailability}
                onChange={(event) =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedAvailability: event.target.checked,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#23A65F"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Ketersediaan</Text>
                  </Group>
                }
              />
            )}
          </Card>
          <Card withBorder p="xs">
            {!neracaSummaryFilter.isActiveStackedType && (
              <Radio
                checked={neracaSummaryFilter.isCheckedRequired}
                onChange={() =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedRequired: true,
                    isCheckedNeraca: false,
                    isCheckedAvailability: false,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#B11016"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Kebutuhan</Text>
                  </Group>
                }
              />
            )}
            {neracaSummaryFilter.isActiveStackedType && (
              <Checkbox
                checked={neracaSummaryFilter.isCheckedRequired}
                onChange={(event) =>
                  setNeracaSummaryFilter((oldForm) => ({
                    ...oldForm,
                    isCheckedRequired: event.target.checked,
                  }))
                }
                label={
                  <Group gap="xs">
                    <Card bg={"#B11016"} w={9} h={9} radius={50} p={0} />
                    <Text size="sm">Kebutuhan</Text>
                  </Group>
                }
              />
            )}
          </Card>
        </Group>
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
        isLoading={isLoadingNeracaAvailabilityChart}
        height={240}
      >
        <AreaChart
          h={240}
          data={dataNeracaAvailabilityChart?.stock ?? []}
          dataKey="period"
          series={chartSeries}
          curveType="natural"
          tickLine="none"
          gridAxis="none"
          valueFormatter={(value) => `${FormatNumber(value)} ${unitSuffix}`}
        />
      </LoadingPageContainer>
    </Card>
  );
}

export default NeracaSummaryChart;
