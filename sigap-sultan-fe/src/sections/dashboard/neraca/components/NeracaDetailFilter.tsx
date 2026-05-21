"use client";

import { Button, Grid, Select, Text, rem } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconSearch } from "@tabler/icons-react";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { isNotEmpty, useForm } from "@mantine/form";
import { Dispatch, SetStateAction } from "react";
import useCityOptions from "@/hooks/options/useCityOptions";
import SelectDropdownNested from "@/app/components/SelectDropdownNested";
import { useQuery } from "@tanstack/react-query";
import { neracaApi } from "@/api/base/neraca";

export type NeracaDetailFilterState = {
  commodityType: OptionMap<number> | null;
  city: OptionMap<string> | null;
  requestTimestamp: number;
  date: Date | null;
};

type Props = {
  neracaDetailFilterSubmitted: NeracaDetailFilterState;
  setNeracaDetailFilterSubmitted: Dispatch<
    SetStateAction<NeracaDetailFilterState>
  >;
  isLoading: boolean;
};

function NeracaDetailFilter(props: Props) {
  const {
    neracaDetailFilterSubmitted,
    setNeracaDetailFilterSubmitted,
    isLoading,
  } = props;
  const form = useForm<NeracaDetailFilterState>({
    mode: "controlled",
    initialValues: neracaDetailFilterSubmitted,
    validate: {
      commodityType: isNotEmpty("Komoditas harus diisi"),
      city: isNotEmpty("Kota harus diisi"),
    },
  });

  const { commodities } = useCommodityOptions(
    (commodities) => {
      form.setFieldValue("commodityType", commodities[0]);
      setNeracaDetailFilterSubmitted((oldForm) => ({
        ...oldForm,
        commodityType: commodities[0],
      }));
    },
    "neraca-detail-filter",
    "neraca"
  );

  const { cities } = useCityOptions(
    (cities) => {
      form.setFieldValue("city", cities[0]);
      setNeracaDetailFilterSubmitted((oldForm) => ({
        ...oldForm,
        city: cities[0],
      }));
    },
    "neraca-detail-filter",
    "neraca"
  );

  useQuery({
    queryKey: [
      "price-detail-latest-date-exist",
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
      onSearch({
        ...form.getValues(),
        date: new Date(result),
      });

      return result;
    },
    enabled: !!form.getValues().commodityType,
  });

  const onSearch = (form: NeracaDetailFilterState) => {
    setNeracaDetailFilterSubmitted({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
    <form onSubmit={form.onSubmit(onSearch)}>
      <Text size="lg">Filter</Text>
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

export default NeracaDetailFilter;
