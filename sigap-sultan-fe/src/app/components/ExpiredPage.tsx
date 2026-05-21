import React from "react";
import { Button, Stack, Text } from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

type Props = {
  title: string;
  onPress: () => void;
};

function ExpiredPage(props: Props) {
  return (
    <Stack align="center" py={100}>
      <Stack align="center" py="md" gap="xl">
        <Image
          src={assetPrefix("/logo/logo_bi_sulsel.svg")}
          alt={"Top header logo"}
          width={150}
          height={70}
        />
        <Image
          src={assetPrefix("/expired.svg")}
          alt={"empty-page"}
          width={528}
          height={391}
        />
        <Stack align="center" py="md" gap="md" w={{ md: 528 }}>
          <Text size="xl" c="blue" ta="center">
            {props.title}
          </Text>
          <Button variant="filled" type="submit" onClick={props.onPress}>
            Kembali
          </Button>
        </Stack>
      </Stack>
    </Stack>
  );
}

export default ExpiredPage;
