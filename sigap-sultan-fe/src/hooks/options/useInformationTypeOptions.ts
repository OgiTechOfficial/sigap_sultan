import { informationTypeApi } from "@/api/base/information-type";
import { OptionMap } from "@/types/option";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";

function useInformationTypeOptions(
  callback?: (commodities: OptionMap<string>[]) => void,
  entryPointKey?: string
) {
  const { data: informationType = [] } = useQuery<OptionMap<string>[]>({
    queryKey: ["information-type-list", entryPointKey],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await informationTypeApi.getList({ page: 1, limit: 100 });

      if (error || !result) {
        throw new Error(
          displayMessage ?? "Failed to fetch information type list"
        );
      }

      const informationType = result.map((informationType) => ({
        label: informationType.name,
        value: informationType.id.toString(),
        children: informationType.detailJenisInformasi.map((child) => ({
          label: child.name,
          value: child.code,
        })),
      }));

      if (typeof callback === "function") {
        callback(informationType);
      }

      return informationType;
    },
  });

  return {
    informationType,
  };
}

export default useInformationTypeOptions;
