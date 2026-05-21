import { StockTierNeracaType, StockTierType } from "./neraca";

export type StockPerCityData = {
  cityId: number;
  city: string;
  stock: number;
  tier: StockTierType | StockTierNeracaType;
};

export type StockPerCityDataMap = {
  [cityName: string]: StockPerCityData;
};

export type StockPerCityCardData = {
  cityId: number;
  cityNumber?: number;
  city: string;
  cityImage: string;
  stock: number;
  stockDiff?: number;
  tier: StockTierType | StockTierNeracaType;
  stockType: "PERCENTAGE" | "ABSOLUTE";
};
