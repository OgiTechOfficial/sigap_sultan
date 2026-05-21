"use client";

import { useState } from "react";
import { Card, Grid, Select, Text, Stack, Group } from "@mantine/core";
import { OptionMap } from "@/types/option";
import PricePerCityMap from "@/app/components/PricePerCityMap";
import { useDisclosure } from "@mantine/hooks";
import CompareNationalPerCityChartModal from "./components/CompareNationalPerCityChartModal";
import CompareNationalPerCityTableModal from "./components/CompareNationalPerCityTableModal";
import { PricePerCityData } from "@/types/pricePerCityMap";
import { PriceTableState } from "./components/PriceTableFilter";
import useCompareNational from "@/sections/dashboard/price-table/hooks/useCompareNational";
import { PriceTierType } from "@/types/price";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import PricePerCityCardCompact from "@/app/components/PricePerCityCardCompact";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";
import { RegionType } from "@/types/city";
import EmptyPage from "@/app/components/EmptyPage";

type SelectedPricePerCity = {
  city: PricePerCityData;
  regionType: RegionType;
};

type Props = {
  priceTableState: PriceTableState;
};

function CompareNational(props: Props) {
  const { priceTableState } = props;
  const [sortBy, setSortBy] = useState<string | null>("");
  const [mapModalOpened, { open: openMapModal, close: closeMapModal }] =
    useDisclosure(false);
  const [tableModalOpened, { open: openTableModal, close: closeTableModal }] =
    useDisclosure(false);
  const [selectedCity, setSelectedCity] = useState<SelectedPricePerCity | null>(
    null
  );
  const [selectedStatus, setSelectedStatus] = useState<PriceTierType | null>(
    null
  );
  const {
    isLoadingPriceLevelList,
    isLoadingPriceLevelMap,
    dataPriceLevelMap,
    dataPriceLevelList,
  } = useCompareNational(priceTableState, sortBy);
  const { commodityUnitMap } = useCommodityOptions();

  const sortOptions: OptionMap<string>[] = [
    { label: "Default", value: "" },
    { label: "Urutkan Kecil ke Besar", value: "asc" },
    { label: "Urutkan Besar ke Kecil", value: "desc" },
  ];

  const onSelectCity = (city: PricePerCityData) => {
    setSelectedCity({ city, regionType: "CITY" });
    openMapModal();
  };

  const onSelectProvince = (city: PricePerCityData) => {
    setSelectedCity({ city, regionType: "PROVINCE" });
    openMapModal();
  };

  const onSelectCountry = (city: PricePerCityData) => {
    setSelectedCity({ city, regionType: "COUNTRY" });
    openMapModal();
  };

  const onSelectStatus = (status: PriceTierType) => {
    setSelectedStatus(status);
    openTableModal();
  };

  return (
    <>
      <Text size="xl" mb="md">
        {priceTableState.priceInfoType?.label} -{" "}
        {priceTableState.priceInfoSubType?.label} -{" "}
        {priceTableState.commodityType?.label}
        {priceTableState.commodityType
          ? ` / ${commodityUnitMap[priceTableState.commodityType?.value]}`
          : ""}
      </Text>
      <Grid>
        <Grid.Col span={{ sm: 12, md: 3 }}>
          <Group w={"100%"} gap="md" align="flex-start" mb="md">
            {dataPriceLevelList?.pricePerNational && (
              <PricePerCityCardCompact
                data={dataPriceLevelList.pricePerNational}
                onSelect={() =>
                  onSelectCountry(dataPriceLevelList.pricePerNational!)
                }
              />
            )}
            {dataPriceLevelList?.pricePerProvince && (
              <PricePerCityCardCompact
                data={dataPriceLevelList.pricePerProvince}
                onSelect={() =>
                  onSelectProvince(dataPriceLevelList.pricePerProvince!)
                }
              />
            )}
          </Group>
          <LoadingPageContainer isLoading={isLoadingPriceLevelMap}>
            <PricePerCityMap
              onSelectCity={onSelectCity}
              pricePerCityMap={dataPriceLevelMap?.priceLevel ?? {}}
              priceTiers={
                dataPriceLevelMap
                  ? dataPriceLevelMap.priceTier.map(
                      (price) => dataPriceLevelMap.priceTierMap[price]
                    )
                  : []
              }
              projectionConfig={{
                center: [480.5, -4.15],
                scale: 17000,
              }}
              height={1600}
            />
            {selectedCity && (
              <CompareNationalPerCityChartModal
                opened={mapModalOpened}
                close={closeMapModal}
                title="Selisih dibandingkan Nasional (Rp)"
                selectedCity={selectedCity.city}
                priceTableState={priceTableState}
                regionType={selectedCity.regionType}
              />
            )}
          </LoadingPageContainer>
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: 9 }}>
          <Grid mb="md">
            {dataPriceLevelMap?.priceTier.map((price) => (
              <Grid.Col span={{ sm: 12, md: 3 }} key={price}>
                <Card
                  padding="md"
                  radius="md"
                  withBorder
                  style={{ cursor: "pointer" }}
                  onClick={() => onSelectStatus(price)}
                >
                  <Group justify="space-between" gap="md">
                    <Card
                      bg={dataPriceLevelMap?.priceTierMap[price].color}
                      w={38}
                      h={38}
                      radius={50}
                    />
                    <Stack gap={0}>
                      <Text
                        size="xs"
                        c={dataPriceLevelMap?.priceTierMap[price].color}
                      >
                        {dataPriceLevelMap?.priceTierMap[price].title}
                      </Text>
                      <Text size="lg">
                        {dataPriceLevelMap?.summary[price]} Kab / Kota
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
            <CompareNationalPerCityTableModal
              opened={tableModalOpened}
              close={closeTableModal}
              title={`Kab / Kota - Harga Beras ${dataPriceLevelMap?.priceTierMap[selectedStatus].title}`}
              headerLabel={"Perbedaan harga dengan Nasional"}
              priceTableState={priceTableState}
              priceTierType={selectedStatus}
            />
          )}
          {!isLoadingPriceLevelList &&
            (dataPriceLevelList?.pricePerCityPaginated || []).length === 0 && (
              <EmptyPage title="Ups! data yang kamu cari tidak ada, Harap ubah pencarian kamu" />
            )}
          <LoadingPageContainer isLoading={isLoadingPriceLevelList}>
            <Grid>
              {dataPriceLevelList?.pricePerCityPaginated.map(
                (pricePage, index) => (
                  <Grid.Col span={{ sm: 12, md: 3 }}>
                    <Grid>
                      {pricePage.map((price) => (
                        <Grid.Col span={12}>
                          <PricePerCityCardCompact
                            key={index}
                            data={price}
                            onSelect={() => onSelectCity(price)}
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
    </>
  );
}

export default CompareNational;
