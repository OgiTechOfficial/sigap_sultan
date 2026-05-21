"use client";

import { useMemo, useState } from "react";
import { Card, Text, Table, Pagination, Group, Stack } from "@mantine/core";
import { ReportFilterState } from "@/sections/dashboard/report/components/ReportFilter";
import { usePagination } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { reportApi } from "@/api/base/report";
import { notifications } from "@mantine/notifications";
import { add, differenceInMonths, format, lastDayOfMonth } from "date-fns";
import EmptyPage from "@/app/components/EmptyPage";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { FormatNumber } from "@/utils/currency";

type ReportDateRange = {
  label: string;
  value: string;
};

type Props = {
  submittedReportFilterState: ReportFilterState | null;
};

function ReportResultNeraca(props: Props) {
  const { submittedReportFilterState } = props;
  const [page, onChange] = useState(1);
  const pagination = usePagination({ total: 10, page, onChange });

  const { data: dataReportNeraca, isFetching: isLoadingReportNeraca } =
    useQuery({
      queryKey: [
        "report-neraca",
        submittedReportFilterState?.city?.value,
        submittedReportFilterState?.commodityType?.value,
        submittedReportFilterState?.startDate,
        submittedReportFilterState?.endDate,
        submittedReportFilterState?.requestTimestamp,
        pagination.active,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await reportApi.getReportNeraca({
            cityId: submittedReportFilterState?.city
              ? submittedReportFilterState?.city.value
              : undefined,
            commodityId: submittedReportFilterState?.commodityType
              ? submittedReportFilterState?.commodityType.value
              : undefined,
            startDate: submittedReportFilterState?.startDate
              ? format(submittedReportFilterState?.startDate, "yyyy-MM-01")
              : "",
            endDate: submittedReportFilterState?.endDate
              ? format(
                  lastDayOfMonth(submittedReportFilterState?.endDate),
                  "yyyy-MM-dd"
                )
              : "",
            page: pagination.active,
            limit: 10,
          });

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch report neraca");
        }

        return result;
      },
      enabled: !!submittedReportFilterState,
    });

  const reportDateRanges = useMemo(() => {
    const startDate =
      submittedReportFilterState?.startDate ?? add(new Date(), { days: -1 });
    const endDate =
      submittedReportFilterState?.endDate ?? add(new Date(), { days: -1 });
    const monthRange = differenceInMonths(endDate, startDate);
    if (monthRange === 0) return [];
    const dateRanges: ReportDateRange[] = Array.from({
      length: monthRange + 1,
    }).map((_, monthIndex) => {
      const selectedDate = add(startDate, { months: monthIndex });
      return {
        label: format(selectedDate, "MMMM yyyy"),
        value: format(selectedDate, "MMyyyy"),
      };
    });
    return dateRanges;
  }, [
    submittedReportFilterState?.startDate,
    submittedReportFilterState?.endDate,
  ]);

  const [prevSubmittedReportFilterState, setPrevSubmittedReportFilterState] =
    useState<ReportFilterState | null>(null);
  if (
    JSON.stringify(submittedReportFilterState) !==
    JSON.stringify(prevSubmittedReportFilterState)
  ) {
    setPrevSubmittedReportFilterState(submittedReportFilterState);

    pagination.setPage(1);
  }

  if (!isLoadingReportNeraca && dataReportNeraca?.data.length === 0) {
    return (
      <EmptyPage title="Ups! Laporan  yang kamu cari tidak ada. Silahkan pilih filter dan lihat laporan" />
    );
  }

  return (
    <Stack>
      <Card withBorder p={0} mt="md">
        <LoadingPageContainer isLoading={isLoadingReportNeraca}>
          <Table.ScrollContainer minWidth={500}>
            <Table>
              <Table.Thead bg={"#F9FAFB"}>
                <Table.Tr>
                  <Table.Th>Nama Daerah</Table.Th>
                  <Table.Th>Jenis Informasi</Table.Th>
                  {reportDateRanges.map((reportDateRange) => (
                    <Table.Th key={reportDateRange.value}>
                      {reportDateRange.label}
                    </Table.Th>
                  ))}
                </Table.Tr>
              </Table.Thead>
              <Table.Tbody>
                {dataReportNeraca?.data.map((reportData, reportIndex) => (
                  <>
                    {dataReportNeraca?.informationTypes.map(
                      (informationType, informationTypeIndex) => (
                        <Table.Tr
                          key={`${reportIndex}.${informationTypeIndex}`}
                        >
                          {informationTypeIndex %
                            dataReportNeraca?.informationTypes.length ===
                            0 && (
                            <Table.Td
                              rowSpan={
                                dataReportNeraca?.informationTypes.length
                              }
                            >
                              <Text size="md">{reportData.title}</Text>
                            </Table.Td>
                          )}
                          <Table.Td>
                            <Text size="md">{informationType}</Text>
                          </Table.Td>
                          {reportDateRanges.map((reportDateRange) => {
                            const displayValue = (
                              reportData.stocks[informationType] as Record<
                                string,
                                number
                              >
                            )[reportDateRange.value];

                            return (
                              <Table.Td key={reportDateRange.value}>
                                <Text size="md">
                                  {displayValue
                                    ? FormatNumber(displayValue)
                                    : "N/A"}
                                </Text>
                              </Table.Td>
                            );
                          })}
                        </Table.Tr>
                      )
                    )}
                  </>
                ))}
              </Table.Tbody>
            </Table>
          </Table.ScrollContainer>
        </LoadingPageContainer>
      </Card>
      <Pagination.Root
        value={pagination.active}
        total={dataReportNeraca?.totalPage ?? 0}
        onChange={(page) => pagination.setPage(page)}
      >
        <Group gap={5} justify="space-between">
          <Pagination.Previous />
          <Group gap={5} justify="center">
            <Pagination.Items />
          </Group>
          <Pagination.Next />
        </Group>
      </Pagination.Root>
    </Stack>
  );
}

export default ReportResultNeraca;
