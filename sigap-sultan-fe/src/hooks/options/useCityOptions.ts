import { cityApi } from "@/api/base/city";
import { OptionMap } from "@/types/option";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";

function useCityOptions(
  callback?: (commodities: OptionMap<string>[]) => void,
  entryPointKey?: string,
  moduleType?: "neraca"
) {
  const { data: cities = [] } = useQuery<OptionMap<string>[]>({
    queryKey: ["city-list", entryPointKey, moduleType],
    queryFn: async () => {
      const { result, error, displayMessage } = await cityApi.getList({
        provinceId: 73,
      });

      if (error || !result) {
        throw new Error(displayMessage ?? "Failed to fetch city list");
      }

      let cities = result.map((city) => ({
        label: city.name,
        value: String(city.id),
      }));
      if (moduleType === "neraca") {
        cities = [
          {
            label: "Sulawesi Selatan",
            value: "73",
          },
          ...cities,
        ];
      }

      if (typeof callback === "function") {
        callback(cities);
      }

      return cities;
    },
  });

  return {
    cities,
  };
}

export default useCityOptions;
