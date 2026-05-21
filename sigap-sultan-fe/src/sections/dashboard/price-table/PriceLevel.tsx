"use client";

import { useState } from "react";
import { Card, Grid, Select, Text, Group, ScrollArea } from "@mantine/core";
import { OptionMap } from "@/types/option";
import PricePerCityMap from "@/app/components/PricePerCityMap";
import { useDisclosure } from "@mantine/hooks";
import PricePerCityChartModal from "@/sections/dashboard/price-table/components/PricePerCityChartModal";
import { PricePerCityData } from "@/types/pricePerCityMap";
import { PriceTableState } from "./components/PriceTableFilter";
import usePriceLevel from "@/sections/dashboard/price-table/hooks/usePriceLevel";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { lightOrDark } from "@/utils/get-light-or-dark";
import { dataMapEmpty } from "@/constants/map/dataMap";
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

function PriceLevel(props: Props) {
  const { priceTableState } = props;
  const [sortBy, setSortBy] = useState<string | null>("");
  const [mapModalOpened, { open: openMapModal, close: closeMapModal }] =
    useDisclosure(false);
  const [selectedCity, setSelectedCity] = useState<SelectedPricePerCity | null>(
    null
  );
  const {
    isLoadingPriceLevelMap,
    isLoadingPriceLevelList,
    dataPriceLevelMap,
    dataPriceLevelList,
  } = usePriceLevel(priceTableState, sortBy);
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
          {dataPriceLevelList?.pricePerProvince && (
            <PricePerCityCardCompact
              data={dataPriceLevelList.pricePerProvince}
              onSelect={() =>
                onSelectProvince(dataPriceLevelList.pricePerProvince!)
              }
            />
          )}
          <br />
          <ScrollArea>
            <Card withBorder p="sm" bg={"#F5FBFF"} w={740} mb="sm">
              <Text size="xs" mb="xs">{`Percentile Harga (dalam Rp)`}</Text>
              <Group gap={0}>
                {dataPriceLevelMap?.priceTier.map((priceTier) => (
                  <Card
                    flex={1}
                    p={0}
                    bg={priceTier.color}
                    h={24}
                    radius={0}
                    style={{
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <Text
                      size="xs"
                      c={
                        lightOrDark(priceTier.color) === "light"
                          ? "#000"
                          : "#FFF"
                      }
                    >
                      {priceTier.title}
                    </Text>
                  </Card>
                ))}
                <Card
                  flex={1}
                  p={0}
                  bg={dataMapEmpty.color}
                  h={24}
                  radius={0}
                  style={{
                    justifyContent: "center",
                    alignItems: "center",
                  }}
                >
                  <Text
                    size="xs"
                    c={
                      lightOrDark(dataMapEmpty.textColor) === "light"
                        ? "#000"
                        : "#FFF"
                    }
                  >
                    {dataMapEmpty.title}
                  </Text>
                </Card>
              </Group>
              <Group gap={0} mb="xs">
                {dataPriceLevelMap?.priceTier.map((priceTier) => (
                  <Card
                    flex={1}
                    p={0}
                    h={24}
                    bg={"transparent"}
                    radius={0}
                    style={{
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <Text size={"9"}>
                      {priceTier.priceMinRupiahFormat} -{" "}
                      {priceTier.priceMaxRupiahFormat}
                    </Text>
                  </Card>
                ))}
                <Card
                  flex={1}
                  p={0}
                  h={24}
                  bg={"transparent"}
                  radius={0}
                  style={{
                    justifyContent: "center",
                    alignItems: "center",
                  }}
                >
                  <Text size={"9"}>{"N/A"}</Text>
                </Card>
              </Group>
            </Card>
          </ScrollArea>
          <LoadingPageContainer isLoading={isLoadingPriceLevelMap}>
            <PricePerCityMap
              onSelectCity={onSelectCity}
              pricePerCityMap={dataPriceLevelMap?.priceLevel ?? {}}
              priceTiers={dataPriceLevelMap?.priceTier ?? []}
              projectionConfig={{
                center: [480.5, -4.15],
                scale: 17000,
              }}
              height={1600}
            />
            {selectedCity && (
              <PricePerCityChartModal
                opened={mapModalOpened}
                close={closeMapModal}
                title="Harga (Rp)"
                selectedCity={selectedCity.city}
                priceTableState={priceTableState}
                regionType={selectedCity.regionType}
              />
            )}
          </LoadingPageContainer>
        </Grid.Col>
        <Grid.Col span={{ sm: 12, md: 9 }}>
          <Grid mb="xs">
            <Grid.Col span="auto">
              <Select
                label="Urutkan"
                data={sortOptions}
                value={sortBy}
                onChange={(value) => setSortBy(value)}
              />
            </Grid.Col>
          </Grid>
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

export default PriceLevel;
