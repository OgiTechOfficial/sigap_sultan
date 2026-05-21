import { Card, Stack, Text, Divider, Group } from "@mantine/core";

import { FormatCurrencyRupiah } from "@/utils/currency";
import { getColorFromPrice } from "@/utils/price-color";
import { PricePerCityCardData } from "@/types/pricePerCityMap";
import ImageContainer from "./ImageContainer";

type Props = {
  data: PricePerCityCardData;
  onSelect: () => void;
};

function PricePerCityCard(props: Props) {
  const {
    data: {
      city,
      cityImage,
      price,
      priceDiff,
      priceType,
      priceTier,
      priceTierColor,
      priceSummary,
    },
    onSelect,
  } = props;

  if (priceType === "PERCENTAGE") {
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
            <Stack gap={0}>
              {priceSummary && (
                <Text size="md">{FormatCurrencyRupiah(priceSummary)}</Text>
              )}
              <Text size="xl" c={getColorFromPrice(price)}>
                {price <= 0 ? price : "+ " + price}%
              </Text>
              {priceDiff && (
                <Text size="md" c={getColorFromPrice(priceDiff)}>
                  {priceDiff <= 0
                    ? FormatCurrencyRupiah(priceDiff)
                    : "+ " + FormatCurrencyRupiah(priceDiff)}
                </Text>
              )}
              {priceTier && (
                <Text size="md" c={priceTierColor}>
                  {priceTier}
                </Text>
              )}
            </Stack>
            <ImageContainer
              src={cityImage}
              alt={city}
              width={70}
              height={70}
              style={{ objectFit: "contain" }}
              onError={(error) => console.log(error)}
            />
          </Group>
        </Stack>
      </Card>
    );
  }

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
          <Stack gap={0}>
            {priceSummary && (
              <Text size="md">{FormatCurrencyRupiah(priceSummary)}</Text>
            )}
            <Text size="xl">{FormatCurrencyRupiah(price)}</Text>
            {priceDiff && (
              <Text size="md" c={getColorFromPrice(priceDiff)}>
                {priceDiff <= 0
                  ? FormatCurrencyRupiah(priceDiff)
                  : "+ " + FormatCurrencyRupiah(priceDiff)}
              </Text>
            )}
            {priceTier && (
              <Text size="md" c={priceTierColor}>
                {priceTier}
              </Text>
            )}
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

export default PricePerCityCard;
