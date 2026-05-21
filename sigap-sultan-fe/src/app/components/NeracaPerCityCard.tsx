import { Card, Stack, Text, Divider, Group } from "@mantine/core";
import Image from "next/image";

import { StockPerCityCardData } from "@/types/stockPerCityMap";
import { getColorFromPrice } from "@/utils/price-color";
import NeracaStockBadge from "./NeracaStockBadge";
import { StockTier, StockTierType } from "@/types/neraca";

type Props = {
  data: StockPerCityCardData;
  onSelect: () => void;
  stockTier: Record<string, StockTier>;
};

function NeracaPerCityCard(props: Props) {
  const {
    data: { city, cityImage, stock, stockDiff, tier },
    onSelect,
    stockTier,
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
            <Group gap="xs">
              {stockDiff && (
                <Text size="sm" c={getColorFromPrice(stockDiff ?? 0)}>
                  {stockDiff}%
                </Text>
              )}
              <NeracaStockBadge
                stockTier={stockTier[tier]}
                tier={tier as StockTierType}
              />
            </Group>
          </Stack>
          <Image src={cityImage} alt={city} width={70} height={70} />
        </Group>
      </Stack>
    </Card>
  );
}

export default NeracaPerCityCard;
