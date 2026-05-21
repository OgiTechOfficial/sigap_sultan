
export interface OptionMap<T> {
  label: string;
  value: T;
  children?: OptionMap<T>[];
}