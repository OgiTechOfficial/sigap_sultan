import { Card, Stack, Text, Group } from "@mantine/core";

import { FormatCurrencyRupiah } from "@/utils/currency";
import { getColorFromPrice } from "@/utils/price-color";
import { PricePerCityCardData } from "@/types/pricePerCityMap";
import ImageContainer from "./ImageContainer";
import { useMemo } from "react";

type Props = {
  data: PricePerCityCardData;
  onSelect: () => void;
};

function PricePerCityCardCompact(props: Props) {
  const {
    data: { city, cityNumber, cityImage, priceType },
    onSelect,
  } = props;

  return (
    <Card
      padding="md"
      radius="md"
      withBorder
      style={{ cursor: "pointer" }}
      onClick={onSelect}
      flex={1}
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
          {priceType === "PERCENTAGE" ? (
            <PricePercentage {...props.data} />
          ) : (
            <PriceAbsolute {...props.data} />
          )}
        </Group>
      </Stack>
    </Card>
  );
}

function PriceAbsolute({
  priceSummary,
  price,
  priceDiff,
  priceTier,
  priceTierColor,
}: PricePerCityCardData) {
  return (
    <Stack gap={0}>
      {priceSummary && (
        <Text size="md">{FormatCurrencyRupiah(priceSummary)}</Text>
      )}
      <Text size="md">{FormatCurrencyRupiah(price)}</Text>
      {priceDiff ? (
        <Text size="sm" c={getColorFromPrice(priceDiff)}>
          {priceDiff <= 0
            ? FormatCurrencyRupiah(priceDiff)
            : "+ " + FormatCurrencyRupiah(priceDiff)}
        </Text>
      ) : (
        <Stack h={20.8} />
      )}
      {priceTier && (
        <Text size="sm" c={priceTierColor}>
          {priceTier}
        </Text>
      )}
    </Stack>
  );
}

function PricePercentage({
  priceSummary,
  price,
  priceDiff,
  priceTier,
  priceTierColor,
}: PricePerCityCardData) {
  const renderPrice = useMemo(() => {
    if (price === null) return "N/A";
    if (price <= 0) return price + "%";
    return "+ " + price + "%";
  }, [price]);

  return (
    <Stack gap={0}>
      {priceSummary && (
        <Text size="md">{FormatCurrencyRupiah(priceSummary)}</Text>
      )}
      <Text size="md" c={getColorFromPrice(price)}>
        {renderPrice}
      </Text>
      {priceDiff ? (
        <Text size="sm" c={getColorFromPrice(priceDiff)}>
          {priceDiff <= 0
            ? FormatCurrencyRupiah(priceDiff)
            : "+ " + FormatCurrencyRupiah(priceDiff)}
        </Text>
      ) : (
        <Stack h={20.8} />
      )}
      {priceTier && (
        <Text size="sm" c={priceTierColor}>
          {priceTier}
        </Text>
      )}
    </Stack>
  );
}

export default PricePerCityCardCompact;
