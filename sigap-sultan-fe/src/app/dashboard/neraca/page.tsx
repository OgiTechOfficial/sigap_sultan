"use client";

import { useState } from "react";
import { Grid, Text, Container, Group } from "@mantine/core";
import NeracaContainer, {
  NeracaType,
} from "@/sections/dashboard/neraca/NeracaContainer";
import NeracaSummaryChart from "@/sections/dashboard/neraca/components/NeracaSummaryChart";
import NeracaSummaryByPeriode from "@/sections/dashboard/neraca/components/NeracaSummaryByPeriode";
import NeracaComparisonChart from "@/sections/dashboard/neraca/components/NeracaComparisonChart";
import NeracaFilter, {
  NeracaFilterState,
} from "@/sections/dashboard/neraca/components/NeracaFilter";
import NeracaDetailFilter, {
  NeracaDetailFilterState,
} from "@/sections/dashboard/neraca/components/NeracaDetailFilter";
import { add, format } from "date-fns";

export default function NeracaTable() {
  const [neracaFilterSubmitted, setNeracaFilterSubmitted] =
    useState<NeracaFilterState>({
      filterBy: "komoditas",
      neracaInfoType: { label: "Neraca (Stok Akhir)", value: "stok_akhir" },
      date: add(new Date(), { days: -1 }),
      commodityType: null,
      city: null,
      requestTimestamp: new Date().getTime(),
    });
  const [neracaDetailFilterSubmitted, setNeracaDetailFilterSubmitted] =
    useState<NeracaDetailFilterState>({
      city: null,
      commodityType: null,
      requestTimestamp: new Date().getTime(),
      date: add(new Date(), { days: -1 }),
    });
  const [isLoading, setIsLoading] = useState<boolean>(false);

  return (
    <Container fluid p="lg" bg={"#F9FAFB"}>
      <NeracaFilter
        neracaFilterSubmitted={neracaFilterSubmitted}
        setNeracaFilterSubmitted={setNeracaFilterSubmitted}
        isLoading={isLoading}
      />
      <NeracaContainer
        neracaState={neracaFilterSubmitted}
        neracaType={
          `${neracaFilterSubmitted.filterBy}_${neracaFilterSubmitted.neracaInfoType?.value}` as NeracaType
        }
      />
      <NeracaDetailFilter
        neracaDetailFilterSubmitted={neracaDetailFilterSubmitted}
        setNeracaDetailFilterSubmitted={setNeracaDetailFilterSubmitted}
        isLoading={isLoading}
      />
      <Group gap="md" justify="space-between" mb="md">
        <Text size="md" fw="bold">
          Detail Data Berseri Neraca Pangan
        </Text>
        {neracaFilterSubmitted.date && (
          <Text size="md" c={"#667085"}>
            {`(${format(neracaFilterSubmitted.date, "MMMM yyyy")})`}
          </Text>
        )}
      </Group>
      <Grid>
        <Grid.Col span={12}>
          <NeracaSummaryChart
            neracaDetailFilterSubmitted={neracaDetailFilterSubmitted}
          />
        </Grid.Col>
        <Grid.Col span={12}>
          <NeracaSummaryByPeriode
            neracaDetailFilterSubmitted={neracaDetailFilterSubmitted}
          />
        </Grid.Col>
        <Grid.Col span={12}>
          <NeracaComparisonChart
            neracaDetailFilterSubmitted={neracaDetailFilterSubmitted}
          />
        </Grid.Col>
      </Grid>
    </Container>
  );
}
