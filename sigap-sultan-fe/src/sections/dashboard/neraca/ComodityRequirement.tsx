"use client";

import { useState } from "react";
import { Card, Grid, Select, Text, Stack, Group } from "@mantine/core";
import { OptionMap } from "@/types/option";
import { NeracaFilterState } from "./components/NeracaFilter";
import { useDisclosure } from "@mantine/hooks";
import NeracaRequirementPerCityChartModal from "@/sections/dashboard/neraca/components/NeracaRequirementPerCityChartModal";
import NeracaRequirementPerCityTableModal from "@/sections/dashboard/neraca/components/NeracaRequirementPerCityTableModal";
import { neracaEmpty } from "@/constants/neraca";
import useRequirementStockCommodity from "./hooks/useRequirementStockCommodity";
import { StockPerCityData } from "@/types/stockPerCityMap";
import { StockTierType } from "@/types/neraca";
import NeracaStockPerCityMap from "@/app/components/NeracaStockPerCityMap";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import NeracaStockPerCityCardCompact from "@/app/components/NeracaStockPerCityCardCompact";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import EmptyPage from "@/app/components/EmptyPage";

type Props = {
  neracaState: NeracaFilterState;
};

function ComodityRequirement(props: Props) {
  const { neracaState } = props;
  const [sortBy, setSortBy] = useState<string | null>("");
  const [mapModalOpened, { open: openMapModal, close: closeMapModal }] =
    useDisclosure(false);
  const [tableModalOpened, { open: openTableModal, close: closeTableModal }] =
    useDisclosure(false);
  const [selectedCity, setSelectedCity] = useState<StockPerCityData | null>(
    null
  );
  const [selectedStatus, setSelectedStatus] = useState<StockTierType | null>(
    null
  );
  const {
    isLoadingRequirementStockCommodityList,
    isLoadingRequirementStockCommodityMap,
    dataRequirementStockCommodityMap,
    dataRequirementStockCommodityList,
  } = useRequirementStockCommodity(neracaState, sortBy);
  const { commodityUnitMap } = useCommodityOptions(
    () => {},
    "neraca-filter",
    "neraca"
  );
  const unitSuffix = neracaState.commodityType?.value
    ? commodityUnitMap[neracaState.commodityType?.value]
    : "";

  const sortOptions: OptionMap<string>[] = [
    { label: "Default", value: "" },
    { label: "Urutkan Kecil ke Besar", value: "asc" },
    { label: "Urutkan Besar ke Kecil", value: "desc" },
  ];

  const onSelectCity = (city: StockPerCityData) => {
    setSelectedCity(city);
    openMapModal();
  };

  const onSelectStatus = (status: StockTierType) => {
    setSelectedStatus(status);
    openTableModal();
  };

  return (
    <Grid>
      <Grid.Col span={{ sm: 12, md: 3 }}>
        <Text size="xs">
          Keterangan:  Selisih ketersediaan meningkat/ stabil / menurun
          dibandingkan bulan sebelumnya.
        </Text>
        <Card withBorder p="sm" bg={"#F5FBFF"} mt="md" mb="md">
          <Group gap="md">
            {dataRequirementStockCommodityMap?.stockTierCode.map((stock) => (
              <Group gap="sm">
                <Card
                  bg={dataRequirementStockCommodityMap.stockTier[stock].color}
                  w={21}
                  h={21}
                  p={0}
                  radius={50}
                />
                <Group gap={0}>
                  <Text size="xs">
                    {dataRequirementStockCommodityMap.stockTier[stock].title}
                  </Text>
                </Group>
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
        <LoadingPageContainer isLoading={isLoadingRequirementStockCommodityMap}>
          <NeracaStockPerCityMap
            onSelectCity={onSelectCity}
            stockPerCityMap={dataRequirementStockCommodityMap?.cityStock ?? {}}
            stockTier={dataRequirementStockCommodityMap?.stockTier ?? {}}
            projectionConfig={{
              center: [480.5, -4.15],
              scale: 17000,
            }}
            height={1600}
          />
          {selectedCity && (
            <NeracaRequirementPerCityChartModal
              opened={mapModalOpened}
              close={closeMapModal}
              title={`Kebutuhan (${unitSuffix})`}
              selectedCity={selectedCity}
              neracaState={neracaState}
              unitSuffix={unitSuffix}
            />
          )}
        </LoadingPageContainer>
      </Grid.Col>
      <Grid.Col span={{ sm: 12, md: 9 }}>
        <Grid mb="md">
          {dataRequirementStockCommodityMap?.stockTierCode.map((stock) => (
            <Grid.Col span="auto" key={stock}>
              <Card
                padding="md"
                radius="md"
                withBorder
                style={{ cursor: "pointer" }}
                onClick={() => onSelectStatus(stock)}
              >
                <Group justify="space-between" gap="md">
                  <Card
                    bg={dataRequirementStockCommodityMap.stockTier[stock].color}
                    w={38}
                    h={38}
                    radius={50}
                    p={0}
                  />
                  <Stack gap={0}>
                    <Text
                      size="xs"
                      c={
                        dataRequirementStockCommodityMap.stockTier[stock].color
                      }
                    >
                      {dataRequirementStockCommodityMap.stockTier[stock].title}
                    </Text>
                    <Text size="lg">
                      {dataRequirementStockCommodityMap.summary[stock]} Daerah
                    </Text>
                  </Stack>
                </Group>
              </Card>
            </Grid.Col>
          ))}
          <Grid.Col span="auto">
            <Select
              label="Urutkan"
              data={sortOptions}
              value={sortBy}
              onChange={(value) => setSortBy(value)}
            />
          </Grid.Col>
        </Grid>
        {selectedStatus && (
          <NeracaRequirementPerCityTableModal
            opened={tableModalOpened}
            close={closeTableModal}
            title={`Kebutuhan - ${dataRequirementStockCommodityMap?.stockTier[selectedStatus].title}`}
            stockTier={
              dataRequirementStockCommodityMap?.stockTier[selectedStatus] ??
              null
            }
            neracaStatus={selectedStatus}
            neracaState={neracaState}
          />
        )}
        {!isLoadingRequirementStockCommodityList &&
          (dataRequirementStockCommodityList?.stockPerCityPaginated || [])
            .length === 0 && (
            <EmptyPage title="Ups! data yang kamu cari tidak ada, Harap ubah pencarian kamu" />
          )}
        <LoadingPageContainer
          isLoading={isLoadingRequirementStockCommodityList}
        >
          <Grid>
            {dataRequirementStockCommodityList?.stockPerCityPaginated.map(
              (stockPage, index) => (
                <Grid.Col span={{ sm: 12, md: 3 }}>
                  <Grid>
                    {stockPage.map((stock) => (
                      <Grid.Col span={12}>
                        <NeracaStockPerCityCardCompact
                          key={index}
                          data={stock}
                          onSelect={() => onSelectCity(stock)}
                          stockTier={
                            dataRequirementStockCommodityList.stockTier
                          }
                          unitSuffix={unitSuffix}
                        />
                      </Grid.Col>
                    ))}
                  </Grid>
                </Grid.Col>
              )
            )}
          </Grid>
        </LoadingPageContainer>
      </Grid.Col>
    </Grid>
  );
}

export default ComodityRequirement;
