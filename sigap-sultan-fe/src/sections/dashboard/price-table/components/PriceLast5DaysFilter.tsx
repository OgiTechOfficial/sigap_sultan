"use client";

import { Button, Grid, Select, Text, rem } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconSearch } from "@tabler/icons-react";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { isNotEmpty, useForm } from "@mantine/form";
import useCityOptions from "@/hooks/options/useCityOptions";
import { Dispatch, SetStateAction } from "react";
import SelectDropdownNested from "@/app/components/SelectDropdownNested";

export type PriceLast5DaysState = {
  filterBy: string | null;
  city: OptionMap<string> | null;
  commodityType: OptionMap<number> | null;
  requestTimestamp: number;
};

type Props = {
  priceLast5DaysSubmitted: PriceLast5DaysState;
  setPriceLast5DaysSubmitted: Dispatch<SetStateAction<PriceLast5DaysState>>;
  isLoading: boolean;
};

function PriceLast5DaysFilter(props: Props) {
  const { priceLast5DaysSubmitted, setPriceLast5DaysSubmitted } = props;
  const form = useForm<PriceLast5DaysState>({
    mode: "controlled",
    initialValues: priceLast5DaysSubmitted,
    validate: {
      filterBy: isNotEmpty("Lihat berdasarkan harus diisi"),
      city: (value, values) =>
        values.filterBy === "daerah" && !value ? "Daerah harus diisi" : null,
      commodityType: (value, values) =>
        values.filterBy === "komoditas" && !value
          ? "Komoditas harus diisi"
          : null,
    },
  });

  const { commodities } = useCommodityOptions((commodities) => {
    if (form.getValues().filterBy === "komoditas") {
      form.setFieldValue("commodityType", commodities[0]);
      setPriceLast5DaysSubmitted((oldForm) => ({
        ...oldForm,
        commodityType: commodities[0],
      }));
    }
  }, "last-5-days");
  const { cities } = useCityOptions((cities) => {
    if (form.getValues().filterBy === "daerah") {
      form.setFieldValue("city", cities[0]);
      setPriceLast5DaysSubmitted((oldForm) => ({
        ...oldForm,
        city: cities[0],
      }));
    }
  });

  const seenByOptions: OptionMap<string>[] = [
    { label: "Komoditas", value: "komoditas" },
    { label: "Daerah", value: "daerah" },
  ];

  const onSearch = (form: PriceLast5DaysState) => {
    setPriceLast5DaysSubmitted({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
    <form onSubmit={form.onSubmit(onSearch)}>
      <Text size="lg">Filter</Text>
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
      </Grid>
      <Grid pb="sm">
        <Grid.Col span={{ sm: 12, md: 3 }} offset={{ md: 9 }}>
          <Button
            type="submit"
            variant="filled"
            bg={"#005395"}
            fullWidth
            leftSection={
              <IconSearch style={{ width: rem(16), height: rem(16) }} />
            }
          >
            Cari Filter
          </Button>
        </Grid.Col>
      </Grid>
    </form>
  );
}

export default PriceLast5DaysFilter;
