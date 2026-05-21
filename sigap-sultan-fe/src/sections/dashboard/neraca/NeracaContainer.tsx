"use client";

import { NeracaFilterState } from "./components/NeracaFilter";

import ComodityLastStock from "./ComodityLastStock";
import ComodityAvailability from "./ComodityAvailability";
import ComodityRequirement from "./ComodityRequirement";
import RegionLastStock from "./RegionLastStock";
import RegionAvailability from "./RegionAvailability";
import RegionRequirement from "./RegionRequirement";
import EmptyPage from "@/app/components/EmptyPage";

export type NeracaType =
  | "komoditas_stok_akhir"
  | "komoditas_ketersediaan"
  | "komoditas_kebutuhan"
  | "daerah_stok_akhir"
  | "daerah_ketersediaan"
  | "daerah_kebutuhan";

type Props = {
  neracaState: NeracaFilterState;
  neracaType: NeracaType | null;
};

function NeracaContainer(props: Props) {
  switch (props.neracaType) {
    case "komoditas_stok_akhir":
      return <ComodityLastStock {...props} />;
    case "komoditas_ketersediaan":
      return <ComodityAvailability {...props} />;
    case "komoditas_kebutuhan":
      return <ComodityRequirement {...props} />;
    case "daerah_stok_akhir":
      return <RegionLastStock {...props} />;
    case "daerah_ketersediaan":
      return <RegionAvailability {...props} />;
    case "daerah_kebutuhan":
      return <RegionRequirement {...props} />;
    default:
      return (
        <EmptyPage title="Ups! tipe informasi yang kamu cari tidak tersedia, Harap ubah pencarian kamu" />
      );
  }
}

export default NeracaContainer;
