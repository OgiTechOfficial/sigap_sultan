export const FormatCurrencyRupiah = (value: number): string => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0, // Avoid showing cents
  })
    .format(value)
    .replace(/IDR/g, "Rp")
    .replace(/\u00A0/g, " ");
};

export const FormatNumber = (
  value: number,
  minimumFractionDigits?: number
): string => {
  return new Intl.NumberFormat("id-ID", {
    minimumFractionDigits,
  })
    .format(value)
    .replace(/\u00A0/g, "");
};
