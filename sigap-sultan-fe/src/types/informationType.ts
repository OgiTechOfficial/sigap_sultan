
export type InformationTypeDetail = {
  id: number;
  parentId: number | null;
  name: string;
  code: string;
}

export type InformationType = {
  id: number;
  clientId: string;
  parentId: number | null;
  name: string;
  detailJenisInformasi: InformationTypeDetail[];
}