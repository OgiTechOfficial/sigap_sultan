import { Tooltip } from "@mantine/core";
import {
  ComposableMap,
  ComposableMapProps,
  Geographies,
  Geography,
} from "react-simple-maps";
import { assetPrefix } from "@/utils/asset-prefix";
import { StockPerCityData, StockPerCityDataMap } from "@/types/stockPerCityMap";
import { StockTier } from "@/types/neraca";
import { dataMapEmpty } from "@/constants/map/dataMap";

type Props = ComposableMapProps & {
  onSelectCity: (city: StockPerCityData) => void;
  stockPerCityMap: StockPerCityDataMap;
  stockTier: Record<string, StockTier>;
};

function NeracaStockPerCityMap(props: Props) {
  const { onSelectCity, stockPerCityMap, stockTier, ...composableMapProps } =
    props;

  const getColorFromStockMap = (stockPerCity: StockPerCityData) => {
    if (!stockPerCity) return dataMapEmpty.color;
    if (!stockPerCity.tier) return dataMapEmpty.color;

    const selectedStockTier = stockTier[stockPerCity.tier];
    return selectedStockTier?.color ?? dataMapEmpty.color;
  };

  return (
    <ComposableMap {...composableMapProps}>
      <Geographies geography={assetPrefix("/topojson/sulawesi_selatan.json")}>
        {({ geographies }) =>
          geographies.map((geo) => {
            const stockPerCity = stockPerCityMap[geo.properties.kode];
            return (
              <Tooltip.Floating label={geo.properties.kabkot}>
                <Geography
                  key={geo.rsmKey}
                  geography={geo}
                  stroke="#FFF"
                  style={{
                    default: {
                      fill: getColorFromStockMap(stockPerCity),
                    },
                    hover: {
                      fill: getColorFromStockMap(stockPerCity),
                      opacity: 0.8,
                    },
                    pressed: {
                      fill: getColorFromStockMap(stockPerCity),
                    },
                  }}
                  onClick={() => {
                    onSelectCity(stockPerCity);
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

export default NeracaStockPerCityMap;
