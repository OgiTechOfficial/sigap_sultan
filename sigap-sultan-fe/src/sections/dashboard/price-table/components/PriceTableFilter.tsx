"use client";

import { DateInput } from "@mantine/dates";
import { Button, Grid, Select, Skeleton, Stack, rem } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconSearch } from "@tabler/icons-react";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import useInformationTypeOptions from "@/hooks/options/useInformationTypeOptions";
import { isNotEmpty, useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useState } from "react";
import SelectDropdownNested from "@/app/components/SelectDropdownNested";
import { add, format, lastDayOfMonth } from "date-fns";
import { useQuery } from "@tanstack/react-query";
import { priceApi } from "@/api/base/price";

export type PriceTableState = {
  commodityType: OptionMap<number> | null;
  priceInfoType: OptionMap<string> | null;
  priceInfoSubType: OptionMap<string> | null;
  date: Date | null;
  requestTimestamp: number;
};

type Props = {
  priceTableSubmitted: PriceTableState;
  setPriceTableSubmitted: Dispatch<SetStateAction<PriceTableState>>;
  isLoading: boolean;
};

function PriceTableFilter(props: Props) {
  const { priceTableSubmitted, setPriceTableSubmitted, isLoading } = props;
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const form = useForm<PriceTableState>({
    mode: "controlled",
    initialValues: priceTableSubmitted,
    validate: {
      commodityType: isNotEmpty("Komoditas harus diisi"),
      priceInfoType: isNotEmpty("Jenis informasi harus diisi"),
      priceInfoSubType: isNotEmpty("Detail jenis informasi harus diisi"),
      date: isNotEmpty("Tanggal harus diisi"),
    },
  });

  const { commodities } = useCommodityOptions((commodities) => {
    form.setFieldValue("commodityType", commodities[0]);
    setPriceTableSubmitted((oldForm) => ({
      ...oldForm,
      commodityType: commodities[0],
    }));
  }, "price-table");
  const { informationType } = useInformationTypeOptions((informationType) => {
    form.setFieldValue("priceInfoType", informationType[1]);
    setPriceTableSubmitted((oldForm) => ({
      ...oldForm,
      priceInfoType: informationType[0],
    }));
    const priceInfoSubTypes = informationType[1].children;
    if (priceInfoSubTypes) {
      form.setFieldValue("priceInfoSubType", priceInfoSubTypes[0]);
      setPriceTableSubmitted((oldForm) => ({
        ...oldForm,
        priceInfoSubType: priceInfoSubTypes[0],
      }));
    }
  });

  const { data: dataPriceExist, isFetching: isLoadingPriceExist } = useQuery<
    Record<string, boolean>
  >({
    queryKey: [
      "price-exist",
      form.getValues().commodityType?.value,
      selectedDate,
    ],
    queryFn: async () => {
      let selectedDateParam: Date = selectedDate;

      const { result, error, displayMessage } = await priceApi.getExist({
        commodityId: form.getValues().commodityType?.value,
        startDate: format(selectedDateParam, "yyyy-MM-01"),
        endDate: format(lastDayOfMonth(selectedDateParam), "yyyy-MM-dd"),
      });

      if (error || !result) {
        throw new Error(displayMessage ?? "Failed to fetch price exist");
      }

      return result;
    },
    enabled: !!form.getValues().commodityType && !!selectedDate,
  });

  useQuery({
    queryKey: [
      "price-latest-date-exist",
      form.getValues().commodityType?.value,
    ],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await priceApi.getLatestDateExist({
          commodityId: form.getValues().commodityType?.value,
        });

      if (error || !result) {
        throw new Error(
          displayMessage ?? "Failed to fetch price latest date exist"
        );
      }

      form.getInputProps("date").onChange(new Date(result));
      setSelectedDate(new Date(result));
      onSearch({
        ...form.getValues(),
        date: new Date(result),
      });

      return result;
    },
    enabled: !!form.getValues().commodityType,
  });

  const onSearch = (form: PriceTableState) => {
    setPriceTableSubmitted({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
    <form onSubmit={form.onSubmit(onSearch)}>
      <Grid pb="sm">
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <SelectDropdownNested
            label="Komoditas"
            placeholder="Pilih komoditas"
            data={commodities}
            clearable
            searchable
            key={form.key("commodityType")}
            {...form.getInputProps("commodityType")}
            value={form.getInputProps("commodityType").value}
            onChange={form.getInputProps("commodityType").onChange}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <Select
            label="Jenis Informasi"
            data={informationType}
            key={form.key("priceInfoType")}
            {...form.getInputProps("priceInfoType")}
            value={form.getInputProps("priceInfoType").value?.value || null}
            onChange={(_, value) => {
              form.getInputProps("priceInfoType").onChange(value);
              const subTypeOptions =
                informationType.find(
                  (option) =>
                    option.value === form.getValues().priceInfoType?.value
                )?.children ?? [];
              form.setValues({ priceInfoSubType: subTypeOptions[0] });
            }}
            allowDeselect={false}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <Select
            label="Detail Jenis Informasi"
            data={
              informationType.find(
                (option) =>
                  option.value === form.getValues().priceInfoType?.value
              )?.children ?? []
            }
            key={form.key("priceInfoSubType")}
            {...form.getInputProps("priceInfoSubType")}
            value={form.getInputProps("priceInfoSubType").value?.value || null}
            onChange={(_, value) => {
              form.getInputProps("priceInfoSubType").onChange(value);
            }}
            allowDeselect={false}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <DateInput
            label="Tanggal"
            placeholder="Pilih tanggal"
            maxDate={add(new Date(), { days: -1 })}
            excludeDate={(date) => {
              if (isLoadingPriceExist) return true;
              return dataPriceExist
                ? !dataPriceExist[format(date, "yyyyMMdd")]
                : false;
            }}
            renderDay={(date) => {
              const day = date.getDate();
              if (!isLoadingPriceExist) {
                return day;
              }
              return (
                <Stack gap={0} justify="center" align="center">
                  <div>{day}</div>
                  <Skeleton height={4} width={15} radius="xl" />
                </Stack>
              );
            }}
            onNextMonth={(date) => setSelectedDate(date)}
            onPreviousMonth={(date) => setSelectedDate(date)}
            onNextYear={(date) => setSelectedDate(date)}
            onPreviousYear={(date) => setSelectedDate(date)}
            key={form.key("date")}
            {...form.getInputProps("date")}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }} pt="xl">
          <Button
            type="submit"
            variant="filled"
            bg={"#005395"}
            fullWidth
            leftSection={
              <IconSearch style={{ width: rem(16), height: rem(16) }} />
            }
            loading={isLoading}
          >
            Cari Filter
          </Button>
        </Grid.Col>
      </Grid>
    </form>
  );
}

export default PriceTableFilter;
