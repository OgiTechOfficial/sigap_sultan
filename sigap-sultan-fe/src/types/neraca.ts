import { Asset } from "./asset";

export type StockTierNeracaType = "defisit" | "rentan" | "waspada" | "aman";

export type StockTierType = "meningkat" | "menurun" | "stabil";

export type StockTierUnit = "percentage" | "absolute";

export type StockUnit = "ton" | "kg";

export type StockTier = {
  title: string;
  start: number | null;
  end: number | null;
  unit: "percentage" | "absolute";
  color: string;
  backgroundColor: string;
};

export type CityLastStockCommodityMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  tier: StockTierNeracaType;
};

export type ProvinceLastStockCommodityList = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  tier: StockTierNeracaType;
};

export type CityLastStockCommodityList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  tier: StockTierNeracaType;
};

export type CityLastStockLocationSummary = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  tier: StockTierType;
};

export type CityAvailabilityStockLocationSummary = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  ketersediaan: number;
  ketersediaanDiffPercentage: number;
};

export type CityRequirementStockLocationSummary = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  kebutuhan: number;
  kebutuhanDiffPercentage: number;
};

export type CityAvailabilityStockCommodityMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  stockDiff: number;
  tier: StockTierType;
};

export type ProvinceAvailabilityStockCommodityList = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  stockDiff: number;
  tier: StockTierType;
};

export type CityAvailabilityStockCommodityList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  stockDiff: number;
  tier: StockTierType;
};

export type CommodityAvailabilityStockCityMap = {
  id: number;
  clientId: number;
  commodity: {
    id: number;
    name: string;
    assets: Asset;
  };
  stock: number;
  tier: StockTierType;
};

export type StockByPeriod = {
  period: string; // e.g. Januari 2024
  neraca: number;
  ketersediaan: number;
  kebutuhan: number;
  produksi: number;
  tier: StockTierType;
};
