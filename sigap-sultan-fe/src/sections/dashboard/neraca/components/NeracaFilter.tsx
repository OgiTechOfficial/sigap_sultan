"use client";

import { MonthPickerInput } from "@mantine/dates";
import { Button, Grid, Select, rem } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconSearch } from "@tabler/icons-react";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { isNotEmpty, useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useState } from "react";
import useCityOptions from "@/hooks/options/useCityOptions";
import SelectDropdownNested from "@/app/components/SelectDropdownNested";
import { add, format, lastDayOfYear } from "date-fns";
import { useQuery } from "@tanstack/react-query";
import { neracaApi } from "@/api/base/neraca";

export type NeracaFilterState = {
  filterBy: string | null;
  neracaInfoType: OptionMap<string> | null;
  date: Date | null;
  commodityType: OptionMap<number> | null;
  city: OptionMap<string> | null;
  requestTimestamp: number;
};

type Props = {
  neracaFilterSubmitted: NeracaFilterState;
  setNeracaFilterSubmitted: Dispatch<SetStateAction<NeracaFilterState>>;
  isLoading: boolean;
};

function NeracaFilter(props: Props) {
  const { neracaFilterSubmitted, setNeracaFilterSubmitted, isLoading } = props;
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const form = useForm<NeracaFilterState>({
    mode: "controlled",
    initialValues: neracaFilterSubmitted,
    validate: {
      neracaInfoType: isNotEmpty("Jenis informasi harus diisi"),
      date: isNotEmpty("Tanggal harus diisi"),
      city: (value, values) =>
        values.filterBy === "daerah" && !value ? "Daerah harus diisi" : null,
      commodityType: (value, values) =>
        values.filterBy === "komoditas" && !value
          ? "Komoditas harus diisi"
          : null,
    },
  });

  const { commodities } = useCommodityOptions(
    (commodities) => {
      if (form.getValues().filterBy === "komoditas") {
        form.setFieldValue("commodityType", commodities[0]);
        setNeracaFilterSubmitted((oldForm) => ({
          ...oldForm,
          commodityType: commodities[0],
        }));
      }
    },
    "neraca-filter",
    "neraca"
  );
  const { cities } = useCityOptions(
    (cities) => {
      if (form.getValues().filterBy === "daerah") {
        form.setFieldValue("city", cities[0]);
        setNeracaFilterSubmitted((oldForm) => ({
          ...oldForm,
          city: cities[0],
        }));
      }
    },
    "neraca-filter",
    "neraca"
  );

  const seenByOptions: OptionMap<string>[] = [
    { label: "Komoditas", value: "komoditas" },
    { label: "Daerah", value: "daerah" },
  ];

  const informationTypeOptions: OptionMap<string>[] = [
    {
      label: "Neraca (Stok Akhir)",
      value: "stok_akhir",
    },
    {
      label: "Ketersediaan",
      value: "ketersediaan",
    },
    {
      label: "Kebutuhan",
      value: "kebutuhan",
    },
  ];

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

      const { result, error, displayMessage } = await neracaApi.getExist({
        commodityId: form.getValues().commodityType?.value,
        startDate: format(selectedDateParam, "yyyy-01-01"),
        endDate: format(lastDayOfYear(selectedDateParam), "yyyy-MM-dd"),
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
        await neracaApi.getLatestDateExist({
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

  const onSearch = (form: NeracaFilterState) => {
    setNeracaFilterSubmitted({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
    <form onSubmit={form.onSubmit(onSearch)}>
      <Grid pb="sm">
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <Select
            label="Lihat Berdasarkan"
            data={seenByOptions}
            placeholder="Pilih lihat berdasarkan"
            key={form.key("filterBy")}
            {...form.getInputProps("filterBy")}
            value={form.getInputProps("filterBy").value}
            onChange={(value) => {
              form.getInputProps("filterBy").onChange(value);
              form.setFieldValue("city", null);
              form.setFieldValue("commodityType", null);
            }}
            allowDeselect={false}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <Select
            label="Jenis Informasi"
            data={informationTypeOptions}
            key={form.key("neracaInfoType")}
            {...form.getInputProps("neracaInfoType")}
            value={form.getInputProps("neracaInfoType").value?.value || null}
            onChange={(_, value) => {
              form.getInputProps("neracaInfoType").onChange(value);
            }}
            allowDeselect={false}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <MonthPickerInput
            label="Bulan"
            placeholder="Pilih Bulan"
            maxDate={add(new Date(), { days: -1 })}
            getYearControlProps={(date) => {
              if (isLoadingPriceExist) return { disabled: true };
              if (!dataPriceExist) return { disabled: false };
              return {
                disabled: Array.from({ length: 12 }).every((_, index) => {
                  let monthNumber = String(index + 1);
                  if (monthNumber.length === 1) {
                    monthNumber = "0" + monthNumber;
                  }
                  return !dataPriceExist[format(date, `yyyy${monthNumber}01`)];
                }),
              };
            }}
            getMonthControlProps={(date) => {
              if (isLoadingPriceExist) return { disabled: true };
              if (!dataPriceExist) return { disabled: false };
              return { disabled: !dataPriceExist[format(date, "yyyyMM01")] };
            }}
            key={form.key("date")}
            {...form.getInputProps("date")}
            value={form.getInputProps("date").value}
            onChange={form.getInputProps("date").onChange}
            onNextYear={(date) => setSelectedDate(date)}
            onPreviousYear={(date) => setSelectedDate(date)}
            onNextDecade={(date) => setSelectedDate(date)}
          />
        </Grid.Col>
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
            disabled={form.getInputProps("filterBy").value !== "komoditas"}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }}>
          <Select
            label="Daerah"
            placeholder="Pilih daerah"
            data={cities}
            clearable
            searchable
            key={form.key("city")}
            {...form.getInputProps("city")}
            value={form.getInputProps("city").value?.value || null}
            onChange={(_, value) => {
              form.getInputProps("city").onChange(value);
            }}
            disabled={form.getInputProps("filterBy").value !== "daerah"}
            allowDeselect={false}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: "auto" }} pt={{ sm: "md", md: "xl" }}>
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

export default NeracaFilter;
