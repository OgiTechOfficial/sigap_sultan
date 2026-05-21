"use client";

import { useState } from "react";
import { Card, Grid, Container } from "@mantine/core";
import ReportFilter, {
  ReportFilterState,
} from "@/sections/dashboard/report/components/ReportFilter";
import ReportSummary from "@/sections/dashboard/report/components/ReportSummary";
import ReportResultNeraca from "@/sections/dashboard/report/components/ReportResultNeraca";
import ReportResultPrice from "@/sections/dashboard/report/components/ReportResultPrice";

export type PriceState = {
  priceCommodity: any[];
  priceDiff: any;
};

export default function ReportPage() {
  const [submittedReportFilterState, setSubmittedReportFilterState] =
    useState<ReportFilterState | null>(null);

  const handleSearch = (form: ReportFilterState) => {
    setSubmittedReportFilterState({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
    <Container fluid p="lg" bg={"#F9FAFB"}>
      <Grid pb="sm">
        <Grid.Col span={{ sm: 12, md: 4 }}>
          <ReportFilter
            submittedReportFilterState={submittedReportFilterState}
            handleSearch={handleSearch}
          />
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: 8 }}>
          <Card padding="lg" radius="md">
            <ReportSummary
              submittedReportFilterState={submittedReportFilterState}
            />
            {submittedReportFilterState?.reportType === "neraca" && (
              <ReportResultNeraca
                submittedReportFilterState={submittedReportFilterState}
              />
            )}
            {submittedReportFilterState?.reportType === "price" && (
              <ReportResultPrice
                submittedReportFilterState={submittedReportFilterState}
              />
            )}
          </Card>
        </Grid.Col>
      </Grid>
    </Container>
  );
}
