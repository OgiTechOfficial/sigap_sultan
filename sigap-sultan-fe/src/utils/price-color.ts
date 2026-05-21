export const getColorFromPrice = (price: number) => {
  if (typeof price !== "number") return;
  if (price < 0) {
    return "#23A65F";
  } else if (price > 0) {
    return "#DC4531";
  }
  return "#667085";
};
