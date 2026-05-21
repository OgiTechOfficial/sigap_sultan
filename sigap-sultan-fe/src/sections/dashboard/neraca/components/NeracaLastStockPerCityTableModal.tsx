import {
  NeracaLocationSummaryRequest,
  NeracaMapSummaryRequest,
  neracaApi,
} from "@/api/base/neraca";
import { NeracaFilterState } from "@/sections/dashboard/neraca/components/NeracaFilter";
import { StockTierNeracaType } from "@/types/neraca";
import { Text, Group, Modal, Box, Button, Table, Stack } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import LastStockBadge from "../../../../app/components/LastStockBadge";
import { FormatNumber } from "@/utils/currency";

type Props = {
  opened: boolean;
  close: () => void;
  title: string;
  neracaStatus: StockTierNeracaType;
  neracaState: NeracaFilterState;
};

function NeracaLastStockPerCityTableModal(props: Props) {
  const { title, opened, close, neracaState, neracaStatus } = props;

  const { data: dataStockPerCity, isFetching: isLoadingStockPerCity } =
    useQuery({
      queryKey: [
        "neraca-stock-per-city-table-last-stock",
        neracaState.city,
        neracaState.commodityType,
        neracaState.date,
        neracaStatus,
        neracaState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await neracaApi.getNeracaLastStockCommodityHistory({
            status: neracaStatus,
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
                <Table.Th>Volume</Table.Th>
                <Table.Th>Kondisi Neraca (Stok Akhir)</Table.Th>
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
                    <Text size="md">{FormatNumber(data.stock)}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Stack align="flex-start">
                      <LastStockBadge
                        stockTier={dataStockPerCity.stockTier[data.tier]}
                        tier={data.tier as StockTierNeracaType}
                      />
                    </Stack>
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

export default NeracaLastStockPerCityTableModal;
