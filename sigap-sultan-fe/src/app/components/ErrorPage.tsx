import React from "react";
import { Stack, Text } from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

type Props = {
  title: string;
  description: string;
};

function ErrorPage(props: Props) {
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
          src={assetPrefix("/error-server.svg")}
          alt={"empty-page"}
          width={528}
          height={391}
        />
        <Stack align="center" py="md" gap="md" w={{ md: 528 }}>
          <Text size="xl" c="blue" ta="center">
            {props.title}
          </Text>
          <Text size="sm" c="gray" ta="center">
            {props.description}
          </Text>
        </Stack>
      </Stack>
    </Stack>
  );
}

export default ErrorPage;
