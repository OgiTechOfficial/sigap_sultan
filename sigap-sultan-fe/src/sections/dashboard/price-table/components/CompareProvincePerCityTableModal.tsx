import {
  priceCompareProvinceApi,
  PriceCompareProvinceCommodityHistoryRequest,
} from "@/api/base/price/priceCompareProvince";
import LoadingPage from "@/app/components/LoadingPage";
import { PriceTableState } from "@/sections/dashboard/price-table/components/PriceTableFilter";
import { PriceTierType } from "@/types/price";
import { FormatCurrencyRupiah } from "@/utils/currency";
import { getColorFromPrice } from "@/utils/price-color";
import { Text, Group, Modal, Box, Button, Table } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";

type Props = {
  opened: boolean;
  close: () => void;
  title: string;
  headerLabel: string;
  priceTableState: PriceTableState;
  priceTierType: PriceTierType;
};

function CompareProvincePerCityTableModal(props: Props) {
  const { title, opened, close, headerLabel, priceTableState, priceTierType } =
    props;

  const { data: dataPricePerCity, isFetching: isLoadingPricePerCity } =
    useQuery({
      queryKey: [
        "price-per-city-table-compare-province",
        priceTableState.commodityType?.value,
        priceTierType,
        priceTableState.date,
        priceTableState.requestTimestamp,
      ],
      queryFn: async () => {
        const { result, error, displayMessage } =
          await priceCompareProvinceApi.getPriceCompareProvinceCommodityHistory(
            {
              commodityId: priceTableState.commodityType?.value,
              status: priceTierType,
              selectedDate: priceTableState.date
                ? format(priceTableState.date, "yyyy-MM-dd")
                : "",
            } as PriceCompareProvinceCommodityHistoryRequest
          );

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
          {format(priceTableState.date ?? new Date(), "dd MMM yyyy")}
        </Text>
      </Group>
      <Box px="md" mb="md">
        <Table.ScrollContainer minWidth={500}>
          <Table>
            <Table.Thead bg={"#F9FAFB"}>
              <Table.Tr>
                <Table.Th>Nomor</Table.Th>
                <Table.Th>Daerah</Table.Th>
                <Table.Th>Harga</Table.Th>
                <Table.Th>{headerLabel}</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {isLoadingPricePerCity && (
                <Table.Tr>
                  <Table.Td colSpan={4}>
                    <LoadingPage />
                  </Table.Td>
                </Table.Tr>
              )}
              {dataPricePerCity?.priceDiff?.map((data, index) => (
                <Table.Tr key={index}>
                  <Table.Td>
                    <Text size="md">{index + 1}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md">{data.city.name}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md" c={getColorFromPrice(data.priceDiff)}>
                      {FormatCurrencyRupiah(data.price)}
                    </Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md" c={getColorFromPrice(data.priceDiff)}>
                      {FormatCurrencyRupiah(data.priceDiff)}
                    </Text>
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

export default CompareProvincePerCityTableModal;
