"use client";

import { Card, Stack, Table, Text } from "@mantine/core";
import { NeracaDetailFilterState } from "./NeracaDetailFilter";
import { useQuery } from "@tanstack/react-query";
import { neracaApi } from "@/api/base/neraca";
import { notifications } from "@mantine/notifications";
import { useState } from "react";
import { usePagination } from "@mantine/hooks";
import { LoadingPageContainer } from "@/app/components/LoadingPage";
import { FormatNumber } from "@/utils/currency";
import useCommodityOptions from "@/hooks/options/useCommodityOptions";

type Props = {
  neracaDetailFilterSubmitted: NeracaDetailFilterState;
};

function NeracaSummaryByPeriode(props: Props) {
  const { neracaDetailFilterSubmitted } = props;
  const [page, onChange] = useState(1);
  const pagination = usePagination({ total: 10, page, onChange });
  const { commodityUnitMap } = useCommodityOptions(
    () => {},
    "neraca-filter",
    "neraca"
  );
  const unitSuffix = neracaDetailFilterSubmitted.commodityType?.value
    ? commodityUnitMap[neracaDetailFilterSubmitted.commodityType?.value]
    : "";

  const {
    data: dataNeracaLastStockTable,
    isFetching: isLoadingNeracaLastStockTable,
  } = useQuery({
    queryKey: [
      "neraca-last-stock-city-and-commodity",
      neracaDetailFilterSubmitted.city,
      neracaDetailFilterSubmitted.commodityType,
      pagination.active,
      neracaDetailFilterSubmitted.requestTimestamp,
    ],
    queryFn: async () => {
      const { result, error, displayMessage } =
        await neracaApi.getNeracaLastStockByCityAndCommodity(
          {
            cityId: neracaDetailFilterSubmitted.city
              ? Number(neracaDetailFilterSubmitted.city.value)
              : 0,
            commodityId: neracaDetailFilterSubmitted.commodityType
              ? neracaDetailFilterSubmitted.commodityType.value
              : 0,
          },
          {
            page: pagination.active,
            limit: 10,
          }
        );

      if (error || !result) {
        throw new Error(
          displayMessage ?? "Failed to fetch last stock city and commodity"
        );
      }

      return result;
    },
    enabled:
      !!neracaDetailFilterSubmitted.city &&
      !!neracaDetailFilterSubmitted.commodityType,
  });

  return (
    <LoadingPageContainer isLoading={isLoadingNeracaLastStockTable}>
      <Card padding="md" radius="md" withBorder>
        <Table.ScrollContainer minWidth={500}>
          <Table>
            <Table.Thead bg={"#F9FAFB"}>
              <Table.Tr>
                <Table.Th>Periode</Table.Th>
                <Table.Th>Neraca</Table.Th>
                <Table.Th>Ketersediaan</Table.Th>
                {/*<Table.Th>Produksi</Table.Th>*/}
                <Table.Th>Kebutuhan</Table.Th>
                <Table.Th>Status</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {dataNeracaLastStockTable?.stock?.map((data, index) => (
                <Table.Tr key={index}>
                  <Table.Td>
                    <Text size="md">{data.period}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md">{`${FormatNumber(
                      data.neraca
                    )} ${unitSuffix}`}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="md">{`${FormatNumber(
                      data.ketersediaan
                    )} ${unitSuffix}`}</Text>
                  </Table.Td>
                  {/*<Table.Td>*/}
                  {/*  <Text size="md">{`${FormatNumber(*/}
                  {/*    data.produksi*/}
                  {/*  )} ${unitSuffix}`}</Text>*/}
                  {/*</Table.Td>*/}
                  <Table.Td>
                    <Text size="md">{`${FormatNumber(
                      data.kebutuhan
                    )} ${unitSuffix}`}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Stack align="flex-start">
                      {dataNeracaLastStockTable?.stockTier[data.tier] && (
                        <Card
                          bg={`#${
                            dataNeracaLastStockTable?.stockTier[data.tier]
                              .backgroundColor
                          }`}
                          px={10}
                          py={5}
                        >
                          <Text
                            size="sm"
                            c={
                              dataNeracaLastStockTable?.stockTier[data.tier]
                                .color
                            }
                          >
                            {
                              dataNeracaLastStockTable?.stockTier[data.tier]
                                .title
                            }
                          </Text>
                        </Card>
                      )}
                    </Stack>
                  </Table.Td>
                </Table.Tr>
              ))}
            </Table.Tbody>
          </Table>
        </Table.ScrollContainer>
      </Card>
    </LoadingPageContainer>
  );
}

export default NeracaSummaryByPeriode;
