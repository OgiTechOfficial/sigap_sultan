import {
  NeracaMapSummaryRequest,
  neracaAvailabilityApi,
} from "@/api/base/neraca/neracaAvailability";
import { NeracaFilterState } from "@/sections/dashboard/neraca/components/NeracaFilter";
import { StockTier, StockTierType } from "@/types/neraca";
import {
  Text,
  Group,
  Modal,
  Box,
  Button,
  Table,
  Card,
  Stack,
} from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import { FormatNumber } from "@/utils/currency";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";

type Props = {
  opened: boolean;
  close: () => void;
  title: string;
  stockTier: StockTier | null;
  neracaStatus: StockTierType;
  neracaState: NeracaFilterState;
};

function NeracaAvailabilityPerCityTableModal(props: Props) {
  const { title, stockTier, opened, close, neracaState, neracaStatus } = props;
  const { commodityUnitMap } = useCommodityOptions(
    () => {},
    "neraca-filter",
    "neraca"
  );
  const unitSuffix = neracaState.commodityType?.value
    ? commodityUnitMap[neracaState.commodityType?.value]
    : "";

  const { data: dataStockPerCity, isFetching: isLoadingStockPerCity } =
    useQuery({
      queryKey: [
        "neraca-per-city-table-availability",
        neracaState.city,
        neracaState.commodityType,
        neracaState.date,
        neracaStatus,
        neracaState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await neracaAvailabilityApi.getNeracaAvailabilityCommodityHistory({
            status: neracaStatus as NeracaMapSummaryRequest["status"],
            commodityId: neracaState.commodityType?.value,
            selectedDate: neracaState.date
              ? format(neracaState.date, "yyyy-MM-dd")
              : "",
          } as NeracaMapSummaryRequest);

        if (error || !result) {
          throw new Error(displayMessage ?? "Failed to fetch price list");
        }

        return result;
      },
    });

  return (
    <Modal
      size={1178}
      opened={opened}
      onClose={close}
      withCloseButton={false}
      centered
    >
      <Group justify="space-between" gap="md" mb="xl">
        <Text size="lg">{title}</Text>
        <Text size="md">
          {format(neracaState.date ?? new Date(), "dd MMM yyyy")}
        </Text>
      </Group>
      <Box px="md" mb="md">
        <Table.ScrollContainer minWidth={500}>
          <Table>
            <Table.Thead bg={"#F9FAFB"}>
              <Table.Tr>
                <Table.Th>Nomor</Table.Th>
                <Table.Th>Daerah</Table.Th>
                <Table.Th>Volume ketersediaan</Table.Th>
                <Table.Th>% Selisih ketersediaan</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {dataStockPerCity?.cityStock?.map((data, index) => (
                <Table.Tr key={index}>
                  <Table.Td>
                    <Text size="md">{index + 1}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md">{data.city.name}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md">
                      {FormatNumber(data.ketersediaan)} {unitSuffix}
                    </Text>
                  </Table.Td>
                  <Table.Td>
                    {stockTier && (
                      <Stack align="flex-start">
                        <Card bg={stockTier.backgroundColor} px={10} py={5}>
                          <Text size="sm" c={stockTier.color}>
                            {data.ketersediaanDiffPercentage}%
                          </Text>
                        </Card>
                      </Stack>
                    )}
                  </Table.Td>
                </Table.Tr>
              ))}
            </Table.Tbody>
          </Table>
        </Table.ScrollContainer>
      </Box>
      <Button fullWidth onClick={close}>
        Close
      </Button>
    </Modal>
  );
}

export default NeracaAvailabilityPerCityTableModal;
