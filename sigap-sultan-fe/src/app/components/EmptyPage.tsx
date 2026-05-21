import React from "react";
import { Stack, Text } from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

type Props = {
  title: string;
};

function EmptyPage(props: Props) {
  return (
    <Stack align="center">
      <Stack align="center" py="md" gap="md" w={{ md: 528 }}>
        <Image
          src={assetPrefix("/empty.svg")}
          alt={"empty-page"}
          width={528}
          height={391}
        />
        <Text size="xl" c="blue" ta="center">
          {props.title}
        </Text>
      </Stack>
    </Stack>
  );
}

export default EmptyPage;
