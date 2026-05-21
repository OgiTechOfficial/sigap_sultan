export type Commodity = {
  id: number;
  clientId: string;
  parentId: number | null;
  code: number | null;
  name: string;
  unit: string;
};
