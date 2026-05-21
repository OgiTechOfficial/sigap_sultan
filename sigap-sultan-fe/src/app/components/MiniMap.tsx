"use client";

import { Card } from "@mantine/core";
import { ComposableMap, Geographies, Geography } from "react-simple-maps";
import { assetPrefix } from "@/utils/asset-prefix";
import { useMediaQuery } from "@mantine/hooks";

type Props = {
  color?: string;
};

function MiniMap(props: Props) {
  const matches = useMediaQuery("(min-width: 56.25em)");

  return (
    <Card
      withBorder
      p="sm"
      bg={"#F5FBFF"}
      w={matches ? 180 : 90}
      h={matches ? 333 : 160}
      mt={matches ? -333 : -150}
    >
      <ComposableMap
        projectionConfig={{
          center: [480.2, -4],
          scale: 5800,
        }}
        width={160}
      >
        <Geographies geography={assetPrefix("/topojson/sulawesi_selatan.json")}>
          {({ geographies }) =>
            geographies.map((geo) => {
              return (
                <Geography
                  key={geo.rsmKey}
                  geography={geo}
                  fill={props.color ?? "#989898"}
                  stroke="#FFF"
                  style={{
                    default: { outline: "none" },
                    hover: { outline: "none" },
                    pressed: { outline: "none" },
                  }}
                />
              );
            })
          }
        </Geographies>
      </ComposableMap>
    </Card>
  );
}

export default MiniMap;
