import { commodityApi, CommodityListRequest } from "@/api/base/commodity";
import { OptionMap } from "@/types/option";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";

type CommodityOption = {
  commodities: OptionMap<number>[];
  commodityUnitMap: Record<string, string>;
};

function useCommodityOptions(
  callback?: (commodities: OptionMap<number>[]) => void,
  entryPointKey?: string,
  moduleType?: CommodityListRequest["moduleType"]
) {
  const { data } = useQuery<CommodityOption>({
    queryKey: ["commodity-list", entryPointKey, moduleType],
    queryFn: async () => {
      const { result, error, displayMessage } = await commodityApi.getList({
        name: "",
        moduleType,
      });

      if (error || !result) {
        throw new Error(displayMessage ?? "Failed to fetch commodity list");
      }

      const commodittGroups = result.filter((commodity) => !commodity.parentId);
      const commodities = commodittGroups.map((commodityGroup) => ({
        label: commodityGroup.name,
        value: commodityGroup.id,
        children:
          moduleType !== "neraca"
            ? result
                .filter((commodity) => commodity.parentId === commodityGroup.id)
                .map((commodity) => ({
                  label: commodity.name,
                  value: commodity.id,
                }))
            : [],
      }));

      let commodityUnitMap: Record<string, string> = {};
      result.forEach((commodity) => {
        commodityUnitMap = {
          ...commodityUnitMap,
          [commodity.id]: commodity.unit,
        };
      });

      if (typeof callback === "function") {
        callback(commodities);
      }

      return {
        commodities,
        commodityUnitMap,
      };
    },
  });

  return {
    commodities: data?.commodities || [],
    commodityUnitMap: data?.commodityUnitMap || {},
  };
}

export default useCommodityOptions;
