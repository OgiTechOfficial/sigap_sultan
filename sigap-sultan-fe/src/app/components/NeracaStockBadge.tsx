import { StockTier, StockTierType } from "@/types/neraca";
import { Card, Text } from "@mantine/core";

type Props = {
  stockTier: StockTier;
  tier: StockTierType;
  dense?: boolean;
};

function NeracaStockBadge({ stockTier, tier, dense = false }: Props) {
  if (!stockTier) return null;

  const px = dense ? 5 : 10;
  const py = dense ? 2 : 5;

  switch (tier) {
    case "menurun":
      return (
        <Card bg={"#FFFAEB"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    case "stabil":
      return (
        <Card bg={"#ECFDF3"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    case "meningkat":
      return (
        <Card bg={"#D1FADF"} px={px} py={py}>
          <Text size="sm" c={stockTier.color}>
            {stockTier.title}
          </Text>
        </Card>
      );
    default:
      return null;
  }
}

export default NeracaStockBadge;
