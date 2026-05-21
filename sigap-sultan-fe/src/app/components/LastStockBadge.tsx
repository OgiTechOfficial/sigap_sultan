import { StockTier, StockTierNeracaType } from "@/types/neraca";
import { Card, Text } from "@mantine/core";

type Props = {
  stockTier: StockTier;
  tier: StockTierNeracaType;
  dense?: boolean;
};

function LastStockBadge({ stockTier, tier, dense = false }: Props) {
  if (!stockTier) return null;

  const px = dense ? 5 : 10;
  const py = dense ? 2 : 5;

  switch (tier) {
    case "aman":
      return (
        <Card bg={"#D1FADF"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    case "waspada":
      return (
        <Card bg={"#FFFAEB"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    case "rentan":
      return (
        <Card bg={"#FEEFC6"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    case "defisit":
      return (
        <Card bg={"#FEF3F2"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    default:
      return null;
  }
}

export default LastStockBadge;
