export type City = {
  id: number;
  provinceId: number;
  name: string;
  sequence: number;
  assets_relation_id: number;
};

export type RegionType = "CITY" | "PROVINCE" | "COUNTRY";
