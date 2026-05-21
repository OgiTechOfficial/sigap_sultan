"use client";

import { type UIEvent, useCallback, useMemo, useRef } from "react";
import { Card, Stack } from "@mantine/core";
import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
  type MRT_RowVirtualizer,
} from "mantine-react-table";
import { useInfiniteQuery } from "@tanstack/react-query";
import { ReportFilterState } from "@/sections/dashboard/report/components/ReportFilter";
import { reportApi } from "@/api/base/report";
import { notifications } from "@mantine/notifications";
import { add, differenceInMonths, format, lastDayOfMonth } from "date-fns";

type ReportDateRange = {
  label: string;
  value: string;
};

type Props = {
  submittedReportFilterState: ReportFilterState | null;
};

function ReportResultPrice(props: Props) {
  const { submittedReportFilterState } = props;
  const tableContainerRef = useRef<HTMLDivElement>(null); //we can get access to the underlying TableContainer element and react to its scroll events
  const rowVirtualizerInstanceRef = useRef<MRT_RowVirtualizer>(null); //we can get access to the underlying Virtualizer instance and call its scrollToIndex method

  const { data, fetchNextPage, hasNextPage, isFetching, isLoading, isError } =
    useInfiniteQuery({
      queryKey: ["report-price", submittedReportFilterState],
      queryFn: async ({ pageParam = 1 }) => {
        const { result, error, displayMessage } =
          await reportApi.getReportPrice({
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
            page: pageParam,
            limit: 10,
          });

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch report neraca");
        }

        return result;
      },
      enabled: !!submittedReportFilterState,
      getNextPageParam: (lastPage) => {
        return lastPage.page < lastPage.totalPage // Here I'm assuming you have access to the total number of pages
          ? lastPage.page + 1
          : undefined; // If there is not a next page, getNextPageParam will return undefined and the hasNextPage boolean will be set to 'false'
      },
      select: (data) => {
        return data?.pages.flatMap((page) => page.prices) ?? [];
      },
      initialPageParam: 1,
      refetchOnWindowFocus: false,
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
        label: format(selectedDate, "MMM yyyy"),
        value: format(selectedDate, "MMyyyy"),
      };
    });
    return dateRanges;
  }, [
    submittedReportFilterState?.startDate,
    submittedReportFilterState?.endDate,
  ]);

  const columns = useMemo<MRT_ColumnDef<any>[]>(
    () => [
      {
        accessorKey: "title",
        header: "Nama Daerah",
        size: 300,
      },
      ...reportDateRanges.map((reportDateRange) => ({
        accessorKey: reportDateRange.value,
        header: reportDateRange.label,
        size: 150,
        enableColumnPinning: false,
      })),
    ],
    [reportDateRanges]
  );

  //called on scroll and possibly on mount to fetch more data as the user scrolls and reaches bottom of table
  const fetchMoreOnBottomReached = useCallback(
    (containerRefElement?: HTMLDivElement | null) => {
      if (containerRefElement) {
        const { scrollHeight, scrollTop, clientHeight } = containerRefElement;
        //once the user has scrolled within 50px of the bottom of the table, fetch more data if we can
        if (
          scrollHeight - scrollTop - clientHeight < 50 &&
          !isFetching &&
          hasNextPage
        ) {
          fetchNextPage();
        }
      }
    },
    [fetchNextPage, hasNextPage, isFetching]
  );

  const table = useMantineReactTable({
    columns,
    data: data ?? [],
    enablePagination: false,
    enableRowNumbers: true,
    enableRowVirtualization: true, //optional, but recommended if it is likely going to be more than 100 rows
    manualFiltering: true,
    manualSorting: true,
    enableColumnPinning: true,
    enableSorting: false,
    enableColumnActions: false,
    enableTopToolbar: false,
    initialState: {
      columnPinning: { left: ["title"] },
    },
    mantineTableContainerProps: {
      ref: tableContainerRef, //get access to the table container element
      style: { maxHeight: "600px" }, //give the table a max height
      onScroll: (
        event: UIEvent<HTMLDivElement> //add an event listener to the table container element
      ) => fetchMoreOnBottomReached(event.target as HTMLDivElement),
    },
    mantineToolbarAlertBannerProps: {
      color: "red",
      children: "Error loading data",
    },
    state: {
      isLoading,
      showAlertBanner: isError,
      showProgressBars: isFetching,
    },
    rowVirtualizerInstanceRef, //get access to the virtualizer instance
    rowVirtualizerOptions: { overscan: 10 },
  });

  return (
    <Stack>
      <Card withBorder p={0} mt="md">
        <MantineReactTable table={table} />
      </Card>
    </Stack>
  );
}

export default ReportResultPrice;
