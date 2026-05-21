import React, { useState } from "react";
import {
  CloseButton,
  Combobox,
  Input,
  InputBase,
  InputBaseProps,
  ScrollArea,
  useCombobox,
} from "@mantine/core";
import { OptionMap } from "@/types/option";

type Props<T> = Omit<InputBaseProps, "value" | "onChange"> & {
  placeholder?: string;
  searchable?: boolean;
  clearable?: boolean;
  data: OptionMap<T>[];
  value: OptionMap<T> | null;
  onChange: (value: OptionMap<T> | null) => void;
};

function SelectDropdown<T>(props: Props<T>) {
  const {
    placeholder,
    searchable,
    clearable,
    data,
    value: _value,
    onChange,
    ...restProps
  } = props;
  const [search, setSearch] = useState("");
  const combobox = useCombobox({
    onDropdownClose: () => {
      combobox.resetSelectedOption();
      combobox.focusTarget();
      setSearch("");
    },

    onDropdownOpen: () => {
      combobox.focusSearchInput();
    },
  });

  const [value, setValue] = useState<string | null>(null);

  const [prevValue, setPrevValue] = useState<OptionMap<T> | null>(_value);
  if (_value !== prevValue) {
    setPrevValue(_value);
    setValue(_value?.label ?? null);
  }

  const options = data
    .filter((item) =>
      item.label.toLowerCase().includes(search.toLowerCase().trim())
    )
    .map((item) => (
      <Combobox.Option value={item.value as string} key={item.value as string}>
        {item.label}
      </Combobox.Option>
    ));

  return (
    <Combobox
      store={combobox}
      withinPortal={false}
      onOptionSubmit={(val) => {
        const selectedOption = data.find((item) => item.value === val);
        onChange(selectedOption ?? null);
        setValue(selectedOption?.label ?? null);
        combobox.closeDropdown();
      }}
    >
      <Combobox.Target>
        <InputBase
          {...restProps}
          component="button"
          pointer
          rightSection={
            clearable && value !== null ? (
              <CloseButton
                size="sm"
                onMouseDown={(event) => event.preventDefault()}
                onClick={() => setValue(null)}
                aria-label="Clear value"
              />
            ) : (
              <Combobox.Chevron />
            )
          }
          onClick={() => combobox.toggleDropdown()}
          rightSectionPointerEvents={
            clearable && value === null ? "none" : "all"
          }
        >
          {value || (
            <Input.Placeholder>{placeholder ?? "Pick value"}</Input.Placeholder>
          )}
        </InputBase>
      </Combobox.Target>

      <Combobox.Dropdown>
        {searchable && (
          <Combobox.Search
            value={search}
            onChange={(event) => setSearch(event.currentTarget.value)}
            placeholder={"Search..."}
          />
        )}
        <Combobox.Options>
          <ScrollArea.Autosize type="scroll" mah={200}>
            {options.length > 0 ? (
              options
            ) : (
              <Combobox.Empty>Nothing found</Combobox.Empty>
            )}
          </ScrollArea.Autosize>
        </Combobox.Options>
      </Combobox.Dropdown>
    </Combobox>
  );
}

export default SelectDropdown;
