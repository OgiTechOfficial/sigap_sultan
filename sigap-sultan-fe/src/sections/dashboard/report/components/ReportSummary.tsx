"use client";

import { Card, Grid, Text, Group } from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";
import { format } from "date-fns";
import { ReportFilterState } from "@/sections/dashboard/report/components/ReportFilter";
import {
  reportPeriodOptions,
  reportTypeOptions,
} from "@/sections/dashboard/report/constants/reportOptions";

type Props = {
  submittedReportFilterState: ReportFilterState | null;
};

function ReportSummary(props: Props) {
  const { submittedReportFilterState } = props;

  return (
    <>
      <Group gap="md" justify="space-between" mb="md">
        <Text size="xl">Laporan Pangan</Text>
        <Image
          src={assetPrefix("/logo/logo_bi_sulsel.svg")}
          alt="BI dan Sulsel"
          height="100"
          width="100"
        />
      </Group>
      <Grid>
        <Grid.Col span={3}>
          <Text size="sm" mb="md">
            Jenis Laporan
          </Text>
        </Grid.Col>
        <Grid.Col span={9}>
          <Group gap="md">
            <Text size="sm" mb="md">
              :
            </Text>
            <Text size="sm" mb="md">
              {reportTypeOptions.find(
                (option) =>
                  option.value === submittedReportFilterState?.reportType
              )?.label || "-"}
            </Text>
          </Group>
        </Grid.Col>
      </Grid>
      <Grid>
        <Grid.Col span={3}>
          <Text size="sm" mb="md">
            Komoditas
          </Text>
        </Grid.Col>
        <Grid.Col span={9}>
          <Group gap="md">
            <Text size="sm" mb="md">
              :
            </Text>
            <Text size="sm" mb="md">
              {submittedReportFilterState?.commodityType?.label || "-"}
            </Text>
          </Group>
        </Grid.Col>
      </Grid>
      <Grid>
        <Grid.Col span={3}>
          <Text size="sm" mb="md">
            Daerah
          </Text>
        </Grid.Col>
        <Grid.Col span={9}>
          <Group gap="md">
            <Text size="sm" mb="md">
              :
            </Text>
            <Text size="sm" mb="md">
              {submittedReportFilterState?.city?.label || "-"}
            </Text>
          </Group>
        </Grid.Col>
      </Grid>
      <Grid>
        <Grid.Col span={3}>
          <Text size="sm" mb="md">
            Periode Laporan
          </Text>
        </Grid.Col>
        <Grid.Col span={9}>
          <Group gap="md">
            <Text size="sm" mb="md">
              :
            </Text>
            <Text size="sm" mb="md">
              {reportPeriodOptions.find(
                (option) =>
                  option.value === submittedReportFilterState?.reportPeriod
              )?.label || "-"}
            </Text>
          </Group>
        </Grid.Col>
      </Grid>
      <Grid>
        <Grid.Col span={3}>
          <Text size="sm" mb="md">
            Tanggal Mulai - Tanggal Akhir
          </Text>
        </Grid.Col>
        <Grid.Col span={9}>
          <Group gap="md">
            <Text size="sm" mb="md">
              :
            </Text>
            <Text size="sm" mb="md">
              {submittedReportFilterState?.startDate
                ? format(submittedReportFilterState.startDate, "MMM yyyy")
                : ""}{" "}
              -{" "}
              {submittedReportFilterState?.endDate
                ? format(submittedReportFilterState.endDate, "MMM yyyy")
                : ""}
            </Text>
          </Group>
        </Grid.Col>
      </Grid>
      <Grid>
        <Grid.Col span={12}>
          <Group justify="flex-start">
            <Card bg={"#F5FBFF"} p="xs" withBorder>
              <Text size="sm">Dalam Satuan Ton</Text>
            </Card>
          </Group>
        </Grid.Col>
      </Grid>
    </>
  );
}

export default ReportSummary;
