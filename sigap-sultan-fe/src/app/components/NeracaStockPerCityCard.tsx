import { Card, Stack, Text, Divider, Group } from "@mantine/core";
import Image from "next/image";

import { StockPerCityCardData } from "@/types/stockPerCityMap";
import { StockTier, StockTierType } from "@/types/neraca";
import NeracaStockBadge from "./NeracaStockBadge";
import ImageContainer from "./ImageContainer";

type Props = {
  data: StockPerCityCardData;
  stockTier: Record<string, StockTier>;
  onSelect: () => void;
};

function NeracaStockPerCityCard(props: Props) {
  const {
    data: { city, cityImage, stock, tier, stockDiff },
    stockTier,
    onSelect,
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
        <Text size="md" ta="center" fw="bold">
          {city}
        </Text>
        <Divider />
        <Group justify="space-between" gap="md">
          <Stack gap={0} align="flex-start">
            <Text size="xl">{stock} Ton</Text>
            <Group gap="md">
              {stockDiff && (
                <Text c={stockTier[tier]?.color}>{stockDiff}%</Text>
              )}
              <NeracaStockBadge
                stockTier={stockTier[tier]}
                tier={tier as StockTierType}
              />
            </Group>
          </Stack>
          <ImageContainer
            src={cityImage}
            alt={city}
            width={70}
            height={70}
            style={{ objectFit: "contain" }}
          />
        </Group>
      </Stack>
    </Card>
  );
}

export default NeracaStockPerCityCard;
