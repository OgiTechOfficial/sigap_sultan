export type ReportNeraca = {
  informationTypes: string[]; // ketersediaan, kebutuhan, neraca
  page: number;
  totalPage: number;
  totalData: number;
  data: {
    title: string;
    /* ex: 
    "ketesediaan": {
      "032024": 2996,
      "042024": 2062
    },
    "kebutuhan": {
      "032024": 3428,
      "042024": 2110
    },
    "neraca": {
      "032024": 2777,
      "042024": 2029
    } */
    stocks: Record<string, Record<string, number>>;
  }[];
};

export type ReportPrice = {
  page: number;
  totalPage: number;
  totalData: number;
  /* ex:
  {
    "title": "Makassar",
    "032024": "Rp 10.000",
    "042024": "Rp 15.000"
  },
  {
    "title": "Palopo",
    "032024": "Rp 10.000",
    "042024": "Rp 15.000"
  } */
  prices: Record<string, string> &
    {
      title: string;
    }[];
};
