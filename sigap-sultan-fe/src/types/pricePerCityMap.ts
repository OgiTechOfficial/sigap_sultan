export type PricePerCityData = {
  cityId: number;
  city: string;
  price: number;
  color?: string;
};

export type PricePerCityDataMap = {
  [cityName: string]: PricePerCityData;
};

export type PricePerCityCardData = {
  cityId: number;
  cityNumber?: number;
  city: string;
  cityImage: string;
  price: number;
  priceDiff?: number;
  priceTier?: string;
  priceTierColor?: string;
  priceSummary?: number;
  priceType: "PERCENTAGE" | "ABSOLUTE";
};
