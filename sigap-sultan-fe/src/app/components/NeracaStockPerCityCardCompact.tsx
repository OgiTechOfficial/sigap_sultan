import { Card, Stack, Text, Group } from "@mantine/core";

import { StockPerCityCardData } from "@/types/stockPerCityMap";
import { StockTier, StockTierType } from "@/types/neraca";
import NeracaStockBadge from "./NeracaStockBadge";
import ImageContainer from "./ImageContainer";
import { FormatNumber } from "@/utils/currency";

type Props = {
  data: StockPerCityCardData;
  stockTier: Record<string, StockTier>;
  onSelect: () => void;
  unitSuffix: string;
};

function NeracaStockPerCityCardCompact(props: Props) {
  const {
    data: { city, cityNumber, cityImage, stock, tier, stockDiff },
    stockTier,
    onSelect,
    unitSuffix,
  } = props;

  return (
    <Card
      padding="md"
      radius="md"
      withBorder
      style={{ cursor: "pointer" }}
      onClick={onSelect}
    >
      <Stack gap="sm">
        <Text size="md" fw="bold" lineClamp={2} h={49.59}>
          {cityNumber ? `${cityNumber}.` : ""} {city}
        </Text>
        <Group justify="space-between" gap="xs">
          <ImageContainer
            src={cityImage}
            alt={city}
            width={36}
            height={36}
            style={{ objectFit: "contain" }}
          />
          <Stack gap={0} align="flex-start" h={49.09}>
            <Text size="md">
              {stock !== null
                ? `${FormatNumber(stock)} ${unitSuffix}`
                : "N/A"}
            </Text>
            {stockTier[tier] && (
              <Group gap="xs" align="flex-start">
                {stockDiff && (
                  <Text c={stockTier[tier].color}>
                    {FormatNumber(stockDiff)}%
                  </Text>
                )}
                <NeracaStockBadge
                  stockTier={stockTier[tier]}
                  tier={tier as StockTierType}
                  dense
                />
              </Group>
            )}
          </Stack>
        </Group>
      </Stack>
    </Card>
  );
}

export default NeracaStockPerCityCardCompact;
