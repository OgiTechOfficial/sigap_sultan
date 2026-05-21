"use client";

import { useState } from "react";
import { Card, Grid, Select, Text, Stack, Group } from "@mantine/core";
import { OptionMap } from "@/types/option";
import Image from "next/image";
import { NeracaFilterState } from "./components/NeracaFilter";
import NeracaStockPerCityMap from "@/app/components/NeracaStockPerCityMap";
import { useDisclosure } from "@mantine/hooks";
import NeracaLastStockPerCityChartModal from "@/sections/dashboard/neraca/components/NeracaLastStockPerCityChartModal";
import NeracaLastStockPerCityTableModal from "@/sections/dashboard/neraca/components/NeracaLastStockPerCityTableModal";
import { neracaEmpty } from "@/constants/neraca";
import { assetPrefix } from "@/utils/asset-prefix";
import useLastStockCommodity from "./hooks/useLastStockCommodity";
import { StockTierNeracaType } from "@/types/neraca";
import { StockPerCityData } from "@/types/stockPerCityMap";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import NeracaLastStockPerCityCardCompact from "@/app/components/NeracaLastStockPerCityCardCompact";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import EmptyPage from "@/app/components/EmptyPage";

type Props = {
  neracaState: NeracaFilterState;
};

function ComodityLastStock(props: Props) {
  const { neracaState } = props;
  const [sortBy, setSortBy] = useState<string | null>("");
  const [mapModalOpened, { open: openMapModal, close: closeMapModal }] =
    useDisclosure(false);
  const [tableModalOpened, { open: openTableModal, close: closeTableModal }] =
    useDisclosure(false);
  const [selectedCity, setSelectedCity] = useState<StockPerCityData | null>(
    null
  );
  const [selectedStatus, setSelectedStatus] =
    useState<StockTierNeracaType | null>(null);
  const {
    isLoadingLastStockCommodityList,
    isLoadingLastStockCommodityMap,
    dataLastStockCommodityMap,
    dataLastStockCommodityList,
  } = useLastStockCommodity(neracaState, sortBy);
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

  const onSelectStatus = (status: StockTierNeracaType) => {
    setSelectedStatus(status);
    openTableModal();
  };

  const renderIcon = (status: StockTierNeracaType) => {
    switch (status) {
      case "aman":
        return (
          <Card bg={"#D1FADF"} w={50} h={50} radius={50} p="xs">
            <Image
              src={assetPrefix("/icon/icon_safe.svg")}
              alt={"safe icon"}
              width={30}
              height={30}
            />
          </Card>
        );
      case "waspada":
        return (
          <Card bg={"#FFFAEB"} w={50} h={50} radius={50} p="xs">
            <Image
              src={assetPrefix("/icon/icon_alert.svg")}
              alt={"alert icon"}
              width={30}
              height={30}
            />
          </Card>
        );
      case "rentan":
        return (
          <Card bg={"#FEEFC6"} w={50} h={50} radius={50} p="xs">
            <Image
              src={assetPrefix("/icon/icon_vulnerable.svg")}
              alt={"vulnerable icon"}
              width={30}
              height={30}
            />
          </Card>
        );
      case "defisit":
        return (
          <Card bg={"#FEF3F2"} w={50} h={50} radius={50} p="xs">
            <Image
              src={assetPrefix("/icon/icon_deficit.svg")}
              alt={"deficit icon"}
              width={30}
              height={30}
            />
          </Card>
        );
      default:
        return null;
    }
  };

  return (
    <Grid>
      <Grid.Col span={{ sm: 12, md: 3 }}>
        <Text size="xs">Keterangan Kondisi Neraca (Stok akhir):</Text>
        <Text size="xs">
          Penentuan Kondisi Neraca (Stok akhir) dihitung dari persentase volume
          neraca (selisih ketersediaandengan kebutuhan) dibagi dengan kebutuhan
          selama satu bulan dengan klasifikasi threshold sebagaimana dibawah.
        </Text>
        <Card withBorder p="sm" bg={"#F5FBFF"} mt="md" mb="md">
          <Group gap="md">
            {dataLastStockCommodityMap?.stockTierCode.map((stock) => (
              <Group gap="sm" key={stock}>
                <Card
                  bg={dataLastStockCommodityMap.stockTier[stock].color}
                  w={21}
                  h={21}
                  p={0}
                  radius={50}
                />
                <Group gap={0}>
                  <Text size="xs">
                    {dataLastStockCommodityMap.stockTier[stock].title}
                  </Text>
                  {dataLastStockCommodityMap.stockTier[stock].start !==
                    null && (
                    <Text size="xs">
                      : {dataLastStockCommodityMap.stockTier[stock].start}% -{" "}
                      {dataLastStockCommodityMap.stockTier[stock].end}%
                    </Text>
                  )}
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
        <LoadingPageContainer isLoading={isLoadingLastStockCommodityMap}>
          <NeracaStockPerCityMap
            onSelectCity={onSelectCity}
            stockPerCityMap={dataLastStockCommodityMap?.cityStock ?? {}}
            stockTier={dataLastStockCommodityMap?.stockTier ?? {}}
            projectionConfig={{
              center: [480.5, -4.15],
              scale: 17000,
            }}
            height={1600}
          />
          {selectedCity && (
            <NeracaLastStockPerCityChartModal
              opened={mapModalOpened}
              close={closeMapModal}
              title={`Neraca (${unitSuffix})`}
              selectedCity={selectedCity}
              neracaState={neracaState}
              unitSuffix={unitSuffix}
            />
          )}
        </LoadingPageContainer>
      </Grid.Col>
      <Grid.Col span={{ sm: 12, md: 9 }}>
        <Grid mb="md">
          {dataLastStockCommodityMap?.stockTierCode.map((stock) => (
            <Grid.Col span="auto" key={stock}>
              <Card
                padding="md"
                radius="md"
                withBorder
                style={{ cursor: "pointer" }}
                onClick={() => onSelectStatus(stock)}
              >
                <Group justify="space-between" gap="md">
                  {renderIcon(stock)}
                  <Stack gap={0}>
                    <Text
                      size="xs"
                      c={dataLastStockCommodityMap.stockTier[stock].color}
                    >
                      {dataLastStockCommodityMap.stockTier[stock].title}
                    </Text>
                    <Text size="lg">
                      {dataLastStockCommodityMap.summary[stock]} Daerah
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
          <NeracaLastStockPerCityTableModal
            opened={tableModalOpened}
            close={closeTableModal}
            title={`Neraca (Stock Akhir) - ${dataLastStockCommodityMap?.stockTier[selectedStatus].title}`}
            neracaStatus={selectedStatus}
            neracaState={neracaState}
          />
        )}
        {!isLoadingLastStockCommodityList &&
          (dataLastStockCommodityList?.stockPerCityPaginated || []).length ===
            0 && (
            <EmptyPage title="Ups! data yang kamu cari tidak ada, Harap ubah pencarian kamu" />
          )}
        <LoadingPageContainer isLoading={isLoadingLastStockCommodityList}>
          <Grid>
            {dataLastStockCommodityList?.stockPerCityPaginated.map(
              (stockPage, pageIndex) => (
                <Grid.Col span={{ sm: 12, md: 3 }} key={`page-${pageIndex}`}>
                  <Grid>
                    {stockPage.map((stock, stockIndex) => (
                      <Grid.Col span={12} key={`stock-${stock.city.id}-${stockIndex}`}>
                        <NeracaLastStockPerCityCardCompact
                          data={stock}
                          onSelect={() => onSelectCity(stock)}
                          stockTier={dataLastStockCommodityList.stockTier}
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

export default ComodityLastStock;
