import { Card, Stack, Text, Divider, Group } from "@mantine/core";
import Image from "next/image";

import { StockPerCityCardData } from "@/types/stockPerCityMap";
import { StockTier, StockTierNeracaType } from "@/types/neraca";
import LastStockBadge from "./LastStockBadge";

type Props = {
  data: StockPerCityCardData;
  stockTier: Record<string, StockTier>;
  onSelect: () => void;
};

function NeracaLastStockPerCityCard(props: Props) {
  const {
    data: { city, cityImage, stock, tier },
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
            <LastStockBadge
              stockTier={stockTier[tier]}
              tier={tier as StockTierNeracaType}
            />
          </Stack>
          <Image src={cityImage} alt={city} width={70} height={70} />
        </Group>
      </Stack>
    </Card>
  );
}

export default NeracaLastStockPerCityCard;
