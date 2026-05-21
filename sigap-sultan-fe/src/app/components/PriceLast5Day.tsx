"use client";

import { Card, Text, Box } from "@mantine/core";
import { FormatCurrencyRupiah } from "@/utils/currency";
import { AreaChart } from "@mantine/charts";
import { PriceLast5DaysByCity, PriceLast5DaysByCommodity } from "@/types/price";

type Props = {
  last5DaysData: PriceLast5DaysByCommodity | PriceLast5DaysByCity;
};

function PriceLast5Days(props: Props) {
  const { last5DaysData } = props;

  const getColorFromPrice = (price: number) => {
    if (price < 0) {
      return "#23A65F";
    } else if (price > 0) {
      return "#DC4531";
    }
    return "#667085";
  };

  return (
    <Card padding="md" radius="md" withBorder>
      <Box mb="md">
        <Text size="sm">{last5DaysData.name}</Text>
        <Text size="md" fw="bold">
          {FormatCurrencyRupiah(last5DaysData.currentPrice)}
        </Text>
        <Text size="sm" c={getColorFromPrice(last5DaysData.priceDiffLast5Days)}>
          {last5DaysData.priceDiffLast5Days < 0
            ? FormatCurrencyRupiah(last5DaysData.priceDiffLast5Days)
            : "+ " + FormatCurrencyRupiah(last5DaysData.priceDiffLast5Days)}
        </Text>
      </Box>
      <AreaChart
        h={100}
        data={last5DaysData.price}
        dataKey="date"
        series={[
          {
            name: "price",
            color: getColorFromPrice(last5DaysData.priceDiffLast5Days),
            label: "Harga",
          },
        ]}
        curveType="natural"
        tickLine="none"
        gridAxis="none"
        withYAxis={false}
        valueFormatter={(value) => FormatCurrencyRupiah(value)}
      />
    </Card>
  );
}

export default PriceLast5Days;
