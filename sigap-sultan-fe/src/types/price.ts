import { Asset } from "./asset";

export type PriceTier = {
  title: string;
  priceMin: number;
  priceMinRupiahFormat: string;
  priceMax: number;
  priceMaxRupiahFormat: string;
  color: string;
};

export type PriceLevelTier = {
  title: string;
  color: string;
};

export type PriceLevelMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
  };
  commodity: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
};

export type PriceLevelMapProvince = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodity: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
  tier: string;
};

export type PriceLevelList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  commodity: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  priceTier: Record<string, PriceLevelTier>; // key: PriceTierType
  priceTierCode: PriceTierType[];
  tier: string;
};

export type PriceLevelProvince = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
  tier: string;
};

export type PriceLast5DaysByCommodity = {
  id: number;
  name: string;
  assets: Asset;
  period: {
    startDate: string;
    endDate: string;
  };
  currentPrice: number;
  currentPriceRupiahFormat: string;
  priceDiffLast5Days: number;
  priceDiffLast5DaysFormat: string;
  price: {
    date: string;
    price: number;
    priceRupiahFormat: string;
  }[];
};

export type PriceLast5DaysByCity = {
  id: number;
  name: string;
  assets: Asset;
  period: {
    startDate: string;
    endDate: string;
  };
  currentPrice: number;
  currentPriceRupiahFormat: string;
  priceDiffLast5Days: number;
  priceDiffLast5DaysFormat: string;
  price: {
    date: string;
    price: number;
    priceRupiahFormat: string;
  }[];
};

export type PriceCompareProvinceMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
  };
  commodity: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  tier: PriceTierType;
};

export type PriceCompareProvinceList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareProvince = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
};

export type PriceCompareProvinceCityHistory = {
  date: string;
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareProvinceCommodityHistory = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareNationalCityHistory = {
  date: string;
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareNationalCommodityHistory = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceMTMCityHistory = {
  date: string;
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceMTMCommodityHistory = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceYTYCityHistory = {
  date: string;
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceYTYCommodityHistory = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceYTDCityHistory = {
  date: string;
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceYTDCommodityHistory = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  inflation: number;
  inflationRupiahFormat: string;
};

export type PriceCompareNationalCityMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  tier: PriceTierType;
};

export type PriceCompareNationalProvinceMap = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  tier: PriceTierType;
};

export type PriceCompareNationalList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareNationalProvince = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
};

export type PriceCompareNational = {
  id: number;
  clientId: number;
  country: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
};

export type PriceMTMCityMap = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  tier: PriceTierType;
};

export type PriceMTMProvinceMap = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
  };
  price: number;
  priceRupiahFormat: string;
  tier: PriceTierType;
};

export type PriceMTMList = {
  id: number;
  clientId: number;
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiff: number;
  priceDiffFormat: string;
  priceDiffPercentage: number;
};

export type PriceMTMProvince = {
  id: number;
  clientId: number;
  province: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
  priceDiffPercentage: number;
};

export type PriceMTM = {
  id: number;
  clientId: number;
  country: {
    id: number;
    name: string;
    assets: Asset;
  };
  price: number;
  priceRupiahFormat: string;
};

export type PriceCommodityCity = {
  city: {
    id: number;
    name: string;
    assets: Asset;
  };
  priceDiff: number;
  priceDiffFormat: string;
  price: number;
  priceRupiahFormat: string;
};

export type PriceTierType = "higher" | "same" | "lower";
