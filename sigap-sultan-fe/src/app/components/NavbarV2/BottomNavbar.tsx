"use client";

import { assetPrefix } from "@/utils/asset-prefix";
import {
  Box,
  BackgroundImage,
  Text,
  Stack,
  Group,
  Divider,
  Grid,
} from "@mantine/core";
import Image from "next/image";

function BottomNavbar() {
  return (
    <Box>
      <BackgroundImage
        src={assetPrefix("/background-image/bg_bottom_navigation.svg")}
      >
        <footer>
          <Grid>
            <Grid.Col span={{ sm: 12, md: 7 }}>
              <Stack align="center" gap={0} pt="xl">
                <Box>
                  <Image
                    src={assetPrefix("/logo/logo_bi_sulsel.svg")}
                    alt={"Top header logo"}
                    width={100}
                    height={50}
                  />
                </Box>
                <Box mt="md">
                  <Image
                    src={assetPrefix("/logo/logo_sigap_sultan_hr.svg")}
                    alt="BI dan Sulsel"
                    width={200}
                    height={50}
                  />
                </Box>
                <Stack align="center" gap={0} mt="md" px={50}>
                  <Text size="sm" fw={700} c="white" ta="center">
                    INFORMASI HARGA DAN PASOKAN PANGAN
                  </Text>
                  <Text size="sm" fw={700} c="#F4CC22" ta="center">
                    SULAWESI SELATAN
                  </Text>
                </Stack>
                <Stack align="center" gap={0} mt="xl" px={50}>
                  <Text size="md" fw={700} c="white" ta="center">
                    Sekretariat Tim Pengendalian Inflasi Provinsi Sulawesi
                    Selatan.
                  </Text>
                  <Text size="md" fw={400} c="white" ta="center">
                    Jl. Jendral Urip Sumoharjo No.269, Panaikang, Kec.
                    Panakkukang, Kota Makassar, Sulawesi Selatan
                  </Text>
                </Stack>
                <Group mt="xl" w={"100%"} px={50}>
                  <Divider w={"100%"} />
                </Group>
                <Group justify="center" gap={0} w={"100%"} p={50}>
                  <Text size="md" fw={700} c="white">
                    neracapangan@sulselprov.go.id
                  </Text>
                </Group>
              </Stack>
            </Grid.Col>
            <Grid.Col span={{ sm: 12, md: 5 }}>
              <Stack align="center" gap={0} pt={100}>
                <Text size="xl" pt="xl" c="white">
                  Sumber Data:
                </Text>
                <Group
                  mt="xl"
                  w={"100%"}
                  align="center"
                  gap="xl"
                  justify="center"
                >
                  <Image
                    src={assetPrefix(
                      "/logo/logo_dinas_ketahanan_pangan_sulsel.png"
                    )}
                    alt={"Top header logo"}
                    width={150}
                    height={100}
                    style={{ objectFit: "contain" }}
                  />
                  <Image
                    src={assetPrefix("/logo/logo_dtph_bun_sulsel.png")}
                    alt={"Top header logo"}
                    width={100}
                    height={100}
                    style={{ objectFit: "contain" }}
                  />
                  <Image
                    src={assetPrefix("/logo/logo_bi_white.png")}
                    alt={"Top header logo"}
                    width={100}
                    height={100}
                    style={{ objectFit: "contain" }}
                  />
                </Group>
              </Stack>
            </Grid.Col>
          </Grid>
        </footer>
      </BackgroundImage>
    </Box>
  );
}

export default BottomNavbar;
