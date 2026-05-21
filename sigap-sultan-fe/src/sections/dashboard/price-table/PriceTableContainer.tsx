"use client";

import { PriceTableState } from "./components/PriceTableFilter";

import PriceLevel from "./PriceLevel";
import CompareSulsel from "./CompareSulsel";
import CompareNational from "./CompareNational";
import PriceChangeYTY from "./PriceChangeYTY";
import PriceChangeYTD from "./PriceChangeYTD";
import PriceChangeMTM from "./PriceChangeMTM";
import EmptyPage from "@/app/components/EmptyPage";

export type PriceTableType =
  | "level_harga"
  | "compare_province"
  | "compare_national"
  | "mtm"
  | "yty"
  | "yoy";

type Props = {
  priceTableState: PriceTableState;
  priceTableType: PriceTableType | null;
};

function PriceTableContainer(props: Props) {
  switch (props.priceTableType) {
    case "level_harga":
      return <PriceLevel {...props} />;
    case "compare_province":
      return <CompareSulsel {...props} />;
    case "compare_national":
      return <CompareNational {...props} />;
    case "mtm":
      return <PriceChangeMTM {...props} />;
    case "yty":
      return <PriceChangeYTD {...props} />;
    case "yoy":
      return <PriceChangeYTY {...props} />;
    default:
      return (
        <EmptyPage title="Ups! tipe informasi yang kamu cari tidak tersedia, Harap ubah pencarian kamu" />
      );
  }
}

export default PriceTableContainer;
