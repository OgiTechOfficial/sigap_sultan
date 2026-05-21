"use client";

import { MonthPickerInput } from "@mantine/dates";
import { Button, Card, Grid, Select, rem } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { IconDownload, IconFileInvoice } from "@tabler/icons-react";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import useCityOptions from "@/hooks/options/useCityOptions";
import { useForm } from "@mantine/form";
import {
  reportPeriodOptions,
  reportTypeOptions,
  seenByOptions,
} from "../constants/reportOptions";
import {
  isAfter,
  isBefore,
  differenceInCalendarMonths,
  add,
  format,
  lastDayOfMonth,
} from "date-fns";
import { notifications } from "@mantine/notifications";
import SelectDropdownNested from "@/app/components/SelectDropdownNested";
import {
  apiReportNeracaDownload,
  apiReportPriceDownload,
} from "@/api/base/api-paths";
import constructUrlSearchParams from "@/utils/construct-url-search-params";
import { apiHostConfig } from "@/config";
import { useMutation } from "@tanstack/react-query";
import {
  reportApi,
  ReportNeracaRequest,
  ReportPriceRequest,
} from "@/api/base/report";
import { APIResult } from "@/utils/api/handle-response";

export type ReportFilterState = {
  reportType: string | null;
  filterBy: string | null;
  commodityType: OptionMap<number> | null;
  city: OptionMap<number> | null;
  reportPeriod: string | null;
  startDate: Date | null;
  endDate: Date | null;
  requestTimestamp: number;
};

const defaultDate = add(new Date(), { days: -1 });

const initialReportFilterState: ReportFilterState = {
  reportType: "neraca",
  filterBy: "komoditas",
  commodityType: null,
  city: null,
  reportPeriod: "monthly",
  startDate: add(defaultDate, { months: -3 }),
  endDate: defaultDate,
  requestTimestamp: defaultDate.getTime(),
};

type Props = {
  submittedReportFilterState: ReportFilterState | null;
  handleSearch: (form: ReportFilterState) => void;
};

function ReportFilter(props: Props) {
  const { submittedReportFilterState, handleSearch } = props;

  const form = useForm<ReportFilterState>({
    mode: "controlled",
    initialValues: initialReportFilterState,
    validate: {
      startDate: (value, values) => {
        if (!value) return "Periode mulai harus diisi";
        if (values.endDate && isAfter(value, values.endDate))
          return "Periode mulai harus kurang dari periode selesai";
        if (
          values.endDate &&
          differenceInCalendarMonths(value, values.endDate) > 12
        )
          return "Maksimal filter 12 bulan";
      },
      endDate: (value, values) => {
        if (!value) return "Periode mulai harus diisi";
        if (values.startDate && isBefore(value, values.startDate))
          return "Periode selesai harus lebih dari periode mulai";
        if (
          values.startDate &&
          differenceInCalendarMonths(value, values.startDate) > 12
        )
          return "Maksimal filter 12 bulan";
      },
      city: (value, values) =>
        values.filterBy === "daerah" && !value ? "Daerah harus diisi" : null,
      commodityType: (value, values) =>
        values.filterBy === "komoditas" && !value
          ? "Komoditas harus diisi"
          : null,
    },
  });

  const { commodities } = useCommodityOptions(
    () => {},
    "report-neraca-filter",
    form.getInputProps("reportType").value === "neraca" ? "neraca" : undefined
  );
  const { cities } = useCityOptions();

  const { mutateAsync: mutateAsyncNeraca } = useMutation<
    APIResult<Blob>,
    Error,
    ReportNeracaRequest
  >({
    mutationKey: ["report-neraca"],
    mutationFn: (payload) => reportApi.downloadReportNeraca(payload),
  });

  const { mutateAsync: mutateAsyncPrice } = useMutation<
    APIResult<Blob>,
    Error,
    ReportPriceRequest
  >({
    mutationKey: ["report-price"],
    mutationFn: (payload) => reportApi.downloadReportPrice(payload),
  });

  const handleDownload = async () => {
    if (!submittedReportFilterState) {
      return notifications.show({
        message: "Download report gagal. Silahkan isi filter laporan",
        color: "red",
        autoClose: 5000,
      });
    }

    if (submittedReportFilterState.reportType === "neraca") {
      const request = {
        cityId: submittedReportFilterState?.city
          ? submittedReportFilterState?.city.value
          : 0,
        commodityId: submittedReportFilterState?.commodityType
          ? submittedReportFilterState?.commodityType.value
          : 0,
        startDate: submittedReportFilterState?.startDate
          ? format(submittedReportFilterState?.startDate, "yyyy-MM-01")
          : "",
        endDate: submittedReportFilterState?.endDate
          ? format(
              lastDayOfMonth(submittedReportFilterState?.endDate),
              "yyyy-MM-dd"
            )
          : "",
      };

      const { displayMessage } = await mutateAsyncNeraca(request);

      if (displayMessage) {
        return notifications.show({
          message: displayMessage ?? "Failed to download report",
          color: "red",
          autoClose: 5000,
        });
      }

      const params = constructUrlSearchParams(request);
      const queryString = new URLSearchParams(params).toString();
      const url = `${apiHostConfig.baseApiUrl}${apiReportNeracaDownload}?${queryString}`;
      window.open(url, "_blank");
    } else if (submittedReportFilterState.reportType === "price") {
      const request = {
        cityId: submittedReportFilterState?.city
          ? submittedReportFilterState?.city.value
          : 0,
        commodityId: submittedReportFilterState?.commodityType
          ? submittedReportFilterState?.commodityType.value
          : 0,
        startDate: submittedReportFilterState?.startDate
          ? format(submittedReportFilterState?.startDate, "yyyy-MM-01")
          : "",
        endDate: submittedReportFilterState?.endDate
          ? format(
              lastDayOfMonth(submittedReportFilterState?.endDate),
              "yyyy-MM-dd"
            )
          : "",
      };
      const { displayMessage } = await mutateAsyncPrice(request);

      if (displayMessage) {
        return notifications.show({
          message: displayMessage ?? "Failed to download report",
          color: "red",
          autoClose: 5000,
        });
      }

      const params = constructUrlSearchParams(request);
      const queryString = new URLSearchParams(params).toString();
      const url = `${apiHostConfig.baseApiUrl}${apiReportPriceDownload}?${queryString}`;
      window.open(url, "_blank");
    }
  };

  return (
    <form onSubmit={form.onSubmit(handleSearch)}>
      <Card padding="lg" radius="md" mb="lg">
        <Grid pb="sm">
          <Grid.Col span={12}>
            <Select
              label="Jenis Laporan"
              data={reportTypeOptions}
              placeholder="Pilih jenis laporan"
              key={form.key("reportType")}
              {...form.getInputProps("reportType")}
              value={form.getInputProps("reportType").value}
              onChange={(value) => {
                form.getInputProps("reportType").onChange(value);
                form.setFieldValue("city", null);
                form.setFieldValue("commodityType", null);
              }}
              allowDeselect={false}
            />
          </Grid.Col>
          <Grid.Col span={12}>
            <Select
              label="Laporan Berdasarkan"
              data={seenByOptions}
              placeholder="Pilih laporan berdasarkan"
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
          <Grid.Col span={12}>
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
          <Grid.Col span={12}>
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
            />
          </Grid.Col>
          <Grid.Col span={12}>
            <Select
              label="Periode Laporan"
              data={reportPeriodOptions}
              placeholder="Pilih periode laporan"
              key={form.key("reportPeriod")}
              {...form.getInputProps("reportPeriod")}
              value={form.getInputProps("reportPeriod").value}
              onChange={form.getInputProps("reportPeriod").onChange}
              allowDeselect={false}
            />
          </Grid.Col>
          <Grid.Col span={12}>
            <MonthPickerInput
              label="Periode Mulai"
              placeholder="Pilih periode mulai"
              key={form.key("startDate")}
              {...form.getInputProps("startDate")}
              value={form.getInputProps("startDate").value}
              onChange={form.getInputProps("startDate").onChange}
            />
          </Grid.Col>
          <Grid.Col span={12}>
            <MonthPickerInput
              label="Periode Selesai"
              placeholder="Pilih periode selesai"
              key={form.key("endDate")}
              {...form.getInputProps("endDate")}
              value={form.getInputProps("endDate").value}
              onChange={form.getInputProps("endDate").onChange}
            />
          </Grid.Col>
          <Grid.Col span={12}>
            <Button
              variant="default"
              fullWidth
              leftSection={
                <IconFileInvoice style={{ width: rem(16), height: rem(16) }} />
              }
              type="submit"
              loading={false}
            >
              Lihat Laporan
            </Button>
          </Grid.Col>
          <Grid.Col span={12}>
            <Button
              fullWidth
              variant="filled"
              bg={"#005395"}
              leftSection={
                <IconDownload style={{ width: rem(16), height: rem(16) }} />
              }
              disabled={!submittedReportFilterState}
              onClick={handleDownload}
            >
              Download Laporan
            </Button>
          </Grid.Col>
        </Grid>
      </Card>
    </form>
  );
}

export default ReportFilter;
