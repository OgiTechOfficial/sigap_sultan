import { OptionMap } from "@/types/option";

export const reportPeriodOptions: OptionMap<string>[] = [
  { label: "Bulanan", value: "monthly" },
];

export const reportTypeOptions: OptionMap<string>[] = [
  { label: "Neraca", value: "neraca" },
  { label: "Harga", value: "price" },
];

export const seenByOptions: OptionMap<string>[] = [
  { label: "Komoditas", value: "komoditas" },
  { label: "Daerah", value: "daerah" },
];