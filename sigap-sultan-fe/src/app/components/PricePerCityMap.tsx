import { Tooltip } from "@mantine/core";
import {
  ComposableMap,
  ComposableMapProps,
  Geographies,
  Geography,
} from "react-simple-maps";
import { assetPrefix } from "@/utils/asset-prefix";
import { PricePerCityData, PricePerCityDataMap } from "@/types/pricePerCityMap";
import { PriceTier } from "@/types/price";
import { dataMapEmpty } from "@/constants/map/dataMap";

type Props = ComposableMapProps & {
  onSelectCity: (city: PricePerCityData) => void;
  pricePerCityMap: PricePerCityDataMap;
  priceTiers: PriceTier[];
};

function PricePerCityMap(props: Props) {
  const { onSelectCity, pricePerCityMap, priceTiers, ...composableMapProps } =
    props;

  const getColorFromPriceMap = (pricePerCity: PricePerCityData) => {
    if (!pricePerCity) return dataMapEmpty.color;
    if (pricePerCity.color) return pricePerCity.color;
    if (!pricePerCity.price) return dataMapEmpty.color;

    const selectedPriceTier = priceTiers.find(
      (priceTier) =>
        priceTier.priceMin <= pricePerCity.price &&
        pricePerCity.price <= priceTier.priceMax
    );
    return selectedPriceTier?.color ?? dataMapEmpty.color;
  };

  return (
    <ComposableMap {...composableMapProps}>
      <Geographies geography={assetPrefix("/topojson/sulawesi_selatan.json")}>
        {({ geographies }) =>
          geographies.map((geo) => {
            const pricePerCity = pricePerCityMap[geo.properties.kode];
            return (
              <Tooltip.Floating label={geo.properties.kabkot}>
                <Geography
                  key={geo.rsmKey}
                  geography={geo}
                  stroke="#FFF"
                  style={{
                    default: {
                      fill: getColorFromPriceMap(pricePerCity),
                      outline: "none",
                    },
                    hover: {
                      fill: getColorFromPriceMap(pricePerCity),
                      outline: "none",
                      opacity: 0.8,
                    },
                    pressed: {
                      fill: getColorFromPriceMap(pricePerCity),
                      outline: "none",
                    },
                  }}
                  onClick={() => {
                    onSelectCity(pricePerCity);
                  }}
                  cursor="pointer"
                />
              </Tooltip.Floating>
            );
          })
        }
      </Geographies>
    </ComposableMap>
  );
}

export default PricePerCityMap;
