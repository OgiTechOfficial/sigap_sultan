"use client";

import {
  Group,
  Button,
  Divider,
  Box,
  Burger,
  BackgroundImage,
  UnstyledButton,
  Avatar,
  Text,
  Menu,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import classes from "./Navbar.module.css";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { assetPrefix } from "@/utils/asset-prefix";
import Link from "next/link";
import { useAuth } from "@/hooks/auth/useAuth";
import { forwardRef } from "react";
import { IconChevronRight } from "@tabler/icons-react";

function Navbar() {
  const [drawerOpened, { toggle: toggleDrawer }] = useDisclosure(false);
  const router = useRouter();
  const auth = useAuth();

  const handleLogout = async () => {
    await auth.signOut();

    router.push("/login");
  };

  return (
    <Box style={{ position: "fixed", zIndex: 10, width: "100%" }}>
      <BackgroundImage
        src={assetPrefix("/background-image/bg_header_navigation.svg")}
      >
        <header className={classes.header}>
          <Group justify="space-between" h="100%">
            <Box visibleFrom="sm">
              <Image
                src={assetPrefix("/header_logo.svg")}
                alt={"Top header logo"}
                width={400}
                height={58}
              />
            </Box>
            <Box hiddenFrom="sm">
              <Image
                src={assetPrefix("/header_logo_mobile.svg")}
                alt={"Top header logo"}
                width={270}
                height={58}
              />
            </Box>
            <Group h="100%" gap={0} visibleFrom="sm">
              <Link href="/dashboard/price-table" className={classes.link}>
                Tabel Harga
              </Link>
              <Link href="/dashboard/neraca" className={classes.link}>
                Tabel Neraca
              </Link>
              <Link href="/dashboard/report" className={classes.link}>
                Unduh Laporan
              </Link>
            </Group>

            {!auth.isAuthenticated && (
              <Group visibleFrom="sm">
                <Button variant="default" onClick={() => router.push("/login")}>
                  Log in
                </Button>
              </Group>
            )}
            {auth.isAuthenticated && (
              <Group visibleFrom="sm">
                <Menu withArrow width={200}>
                  <Menu.Target>
                    <UserButton
                      name={auth.user?.name || ""}
                      email={auth.user?.position || ""}
                    />
                  </Menu.Target>
                  <Menu.Dropdown>
                    <Menu.Item
                      onClick={() => router.push("/dashboard/profile")}
                    >
                      Profile
                    </Menu.Item>
                    <Menu.Item onClick={handleLogout}>Logout</Menu.Item>
                  </Menu.Dropdown>
                </Menu>
              </Group>
            )}

            <Burger
              opened={drawerOpened}
              onClick={toggleDrawer}
              hiddenFrom="sm"
              color="#fff"
            />
          </Group>
        </header>
      </BackgroundImage>

      {drawerOpened && (
        <Box
          p="md"
          hiddenFrom="sm"
          pos="absolute"
          top={60}
          w={"100%"}
          bg="#fff"
          style={{
            zIndex: 10,
            borderBottomLeftRadius: 12,
            borderBottomRightRadius: 12,
            boxShadow:
              "rgba(0, 0, 0, 0.02) 0px 2px 3px 0px, rgba(27, 31, 35, 0.15) 0px 2px 2px 1px",
          }}
        >
          <Button
            variant="subtle"
            fullWidth
            onClick={() => router.push("/dashboard/price-table")}
          >
            Tabel Harga
          </Button>
          <Button
            variant="subtle"
            fullWidth
            onClick={() => router.push("/dashboard/neraca")}
          >
            Tabel Neraca
          </Button>
          <Button
            variant="subtle"
            fullWidth
            onClick={() => router.push("/dashboard/report")}
          >
            Unduh Laporan ga dipake
          </Button>

          <Divider my="sm" />

          {!auth.isAuthenticated && (
            <Button fullWidth onClick={() => router.push("/login")}>
              Log in
            </Button>
          )}

          {auth.isAuthenticated && (
            <>
              <Button
                variant="subtle"
                fullWidth
                onClick={() => router.push("/dashboard/profile")}
              >
                Profile
              </Button>
              <Button variant="subtle" fullWidth onClick={handleLogout}>
                Logout
              </Button>
            </>
          )}
        </Box>
      )}
    </Box>
  );
}

interface UserButtonProps extends React.ComponentPropsWithoutRef<"button"> {
  name: string;
  email: string;
  icon?: React.ReactNode;
}

const UserButton = forwardRef<HTMLButtonElement, UserButtonProps>(
  ({ name, email, icon, ...others }: UserButtonProps, ref) => (
    <UnstyledButton
      ref={ref}
      style={{
        color: "var(--mantine-color-text)",
        borderRadius: "var(--mantine-radius-sm)",
      }}
      {...others}
    >
      <Group>
        <div style={{ flex: 1 }}>
          <Text size="sm" c="white" fw={500}>
            {name}
          </Text>

          <Text c="white" size="xs">
            {email}
          </Text>
        </div>

        {icon || <IconChevronRight color="white" size={16} />}
      </Group>
    </UnstyledButton>
  )
);

export default Navbar;
