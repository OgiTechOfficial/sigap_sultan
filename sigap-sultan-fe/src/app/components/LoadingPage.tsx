import React from "react";
import { Group, Stack, Text } from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

type Props = {
  title?: string;
  height?: number;
};

function LoadingPage(props: Props) {
  const { title, height = 528 } = props;

  return (
    <Group h={height} align="center" justify="center">
      <Stack align="center" justify="center">
        <Image
          src={assetPrefix("/logo/logo_bi.svg")}
          alt={"Top header logo"}
          width={100}
          height={50}
        />
        <Text size="xl" c="blue" ta="center">
          {title ?? "Sedang Proses..."}
        </Text>
      </Stack>
    </Group>
  );
}

type LoadingPageContainerProps = Props & {
  isLoading: boolean;
};

export function LoadingPageContainer(
  props: React.PropsWithChildren<LoadingPageContainerProps>
) {
  const { isLoading, ...restProps } = props;

  if (isLoading) {
    return <LoadingPage {...restProps} />;
  }

  return props.children;
}

export default LoadingPage;
