"use client";

import { Card, Grid, Text, Stack, Group, Divider } from "@mantine/core";
import Image from "next/image";
import { NeracaFilterState } from "./components/NeracaFilter";
import { neracaEmpty } from "@/constants/neraca";
import useAvailabilityRegion from "./hooks/useAvailabilityRegion";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { FormatNumber } from "@/utils/currency";
import EmptyPage from "@/app/components/EmptyPage";

type Props = {
  neracaState: NeracaFilterState;
};

function RegionAvailability(props: Props) {
  const { neracaState } = props;
  const { dataLastStockCityMap, isLoading } =
    useAvailabilityRegion(neracaState);
  const { commodityUnitMap } = useCommodityOptions(
    () => {},
    "neraca-filter",
    "neraca"
  );

  return (
    <Grid>
      <Grid.Col span={12}>
        {!isLoading &&
          (dataLastStockCityMap?.stockTierCode || []).length === 0 && (
            <EmptyPage title="Ups! data yang kamu cari tidak ada, Harap ubah pencarian kamu" />
          )}
        <LoadingPageContainer isLoading={isLoading}>
          <Card padding="lg" radius="md" withBorder>
            <Grid mb="md">
              {dataLastStockCityMap?.stockTierCode.map((data) => (
                <Grid.Col span={{ sm: 12, md: 4 }}>
                  <Card padding="md" radius="md" withBorder>
                    <Group justify="space-between" gap="md">
                      <Stack gap={0}>
                        <Text
                          size="xs"
                          c={dataLastStockCityMap.stockTier[data].color}
                        >
                          {dataLastStockCityMap.stockTier[data].title}
                        </Text>
                        <Text size="lg">
                          {dataLastStockCityMap.summary[data]} Komoditas
                        </Text>
                      </Stack>
                      <Card
                        bg={dataLastStockCityMap?.stockTier[data].color}
                        w={40}
                        h={40}
                        p={0}
                        radius={50}
                      />
                    </Group>
                  </Card>
                </Grid.Col>
              ))}
            </Grid>
            <Text size="md" mb="md">
              Komoditas
            </Text>
            <Card p={0} mb="md">
              <Grid>
                {dataLastStockCityMap?.commodityStock.map(
                  (commodity, index) => (
                    <Grid.Col span={{ sm: 12, md: 4 }} key={index}>
                      <Card
                        padding="md"
                        radius="md"
                        withBorder
                        bg={
                          dataLastStockCityMap.stockTier[commodity.tier]
                            ?.backgroundColor
                        }
                      >
                        <Stack gap="sm">
                          <Text size="md" ta="center" fw="bold">
                            {commodity.commodity.name}
                          </Text>
                          <Divider />
                          <Group justify="space-between" gap="md">
                            <Card
                              bg={
                                dataLastStockCityMap.stockTier[commodity.tier]
                                  ?.backgroundColor
                              }
                              w={90}
                              h={90}
                              radius={50}
                              p="xs"
                            >
                              <Image
                                src={commodity.commodity.assets?.assetsUrl}
                                alt={commodity.commodity.name}
                                width={90}
                                height={90}
                              />
                            </Card>
                            {commodity.stock !== null && commodity.stock > 0 ? (
                              <Stack gap={0} align="flex-end">
                                <Text size="lg">
                                  {FormatNumber(commodity.stock)}{" "}
                                  {commodityUnitMap[commodity.commodity.id]}
                                </Text>
                                <Text
                                  size="sm"
                                  c={
                                    dataLastStockCityMap.stockTier[
                                      commodity.tier
                                    ]?.color
                                  }
                                >
                                  {
                                    dataLastStockCityMap.stockTier[
                                      commodity.tier
                                    ]?.title
                                  }
                                </Text>
                              </Stack>
                            ) : (
                              <Card
                                flex={1}
                                p={0}
                                h={24}
                                bg={"transparent"}
                                radius={0}
                                style={{
                                  justifyContent: "flex-end",
                                  alignItems: "flex-end",
                                }}
                              >
                                <Text size="sm">{"N/A"}</Text>
                              </Card>
                            )}
                          </Group>
                        </Stack>
                      </Card>
                    </Grid.Col>
                  )
                )}
              </Grid>
            </Card>
            <Text size="xs">
              Keterangan:  Selisih ketersediaan meningkat/ stabil / menurun
              dibandingkan bulan sebelumnya.
            </Text>
            <Card withBorder p="sm" bg={"#F5FBFF"} mt="md">
              <Group gap="md">
                {dataLastStockCityMap?.stockTierCode.map((data) => (
                  <Group gap="sm">
                    <Card
                      bg={dataLastStockCityMap.stockTier[data].color}
                      w={21}
                      h={21}
                      p={0}
                      radius={50}
                    />
                    <Text size="xs">
                      {dataLastStockCityMap.stockTier[data].title}
                      {dataLastStockCityMap.stockTier[data].start !== null && (
                        <>
                          : {dataLastStockCityMap.stockTier[data].start}% -{" "}
                          {dataLastStockCityMap.stockTier[data].end}%
                        </>
                      )}
                    </Text>
                  </Group>
                ))}
                <Group gap="sm">
                  <Card
                    bg={neracaEmpty.textColor}
                    w={21}
                    h={21}
                    p={0}
                    radius={50}
                  />
                  <Text size="xs">{neracaEmpty.title}</Text>
                </Group>
              </Group>
            </Card>
          </Card>
        </LoadingPageContainer>
      </Grid.Col>
    </Grid>
  );
}

export default RegionAvailability;
