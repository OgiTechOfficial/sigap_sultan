"use client";

import {
  Stack,
  Text,
  Container,
  Grid,
  PasswordInput,
  TextInput,
  Group,
  Button,
  Anchor,
  rem,
} from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";
import { isNotEmpty, useForm } from "@mantine/form";
import { useRouter } from "next/navigation";
import { IconMail, IconLock } from "@tabler/icons-react";
import { useMutation } from "@tanstack/react-query";
import { authApi, AuthLoginRequest, AuthLoginResponse } from "@/api/base/auth";
import { APIResult } from "@/utils/api/handle-response";
import { notifications } from "@mantine/notifications";
import { useAuth } from "@/hooks/auth/useAuth";

type LoginForm = {
  email: string;
  password: string;
};

export default function Login() {
  const router = useRouter();
  const auth = useAuth();

  const form = useForm<LoginForm>({
    mode: "uncontrolled",
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Format email salah"),
      password: isNotEmpty("Password harus diisi"),
    },
  });

  const { mutateAsync, isPending } = useMutation<
    APIResult<AuthLoginResponse>,
    Error,
    AuthLoginRequest
  >({
    mutationKey: ["auth-login"],
    mutationFn: (payload) => authApi.doLogin(payload),
  });

  const onLogin = async (form: LoginForm) => {
    const { result, error, displayMessage } = await mutateAsync({
      email: form.email,
      password: form.password,
    });

    if (error || !result) {
      return notifications.show({
        message: displayMessage ?? "Failed to fetch login",
        color: "red",
        autoClose: 5000,
      });
    }

    await auth.signIn(result, result.token, result.token);
    router.push("/");
  };

  return (
    <Container fluid p="lg" bg={"#F6F6F6"} style={{ minHeight: "100vh" }}>
      <Stack align="center" py={100}>
        <Stack align="center" py="md" gap="xl">
          <Image
            src={assetPrefix("/logo/logo_bi_sulsel.svg")}
            alt={"Top header logo"}
            width={150}
            height={70}
          />
          <Stack py="md" gap="md" w={{ md: 413 }}>
            <Text fw="bold" c="black" ta="center" style={{ fontSize: 30 }}>
              {"Selamat Datang"}
            </Text>
            <Text size="sm" c="gray" ta="center">
              {
                "Tolong masukan informasi akun mu untuk fitur yang lebih banyak dari biasanya."
              }
            </Text>
            <form onSubmit={form.onSubmit(onLogin)}>
              <Grid pb="sm">
                <Grid.Col span={12}>
                  <TextInput
                    label="Email"
                    placeholder="Masukan Email"
                    leftSection={
                      <IconMail style={{ width: rem(18), height: rem(18) }} />
                    }
                    key={form.key("email")}
                    {...form.getInputProps("email")}
                    value={form.getInputProps("email").value}
                    onChange={(e) => {
                      form.getInputProps("email").onChange(e.target.value);
                    }}
                  />
                </Grid.Col>
                <Grid.Col span={12}>
                  <PasswordInput
                    label="Password"
                    placeholder="Masukan Password"
                    key={form.key("password")}
                    leftSection={
                      <IconLock style={{ width: rem(18), height: rem(18) }} />
                    }
                    {...form.getInputProps("password")}
                    value={form.getInputProps("password").value}
                    onChange={(e) => {
                      form.getInputProps("password").onChange(e.target.value);
                    }}
                  />
                </Grid.Col>
                <Grid.Col span={12}>
                  <Group justify="flex-end">
                    <Anchor href="/forgot-password">Lupa Password?</Anchor>
                  </Group>
                </Grid.Col>
                <Grid.Col span={12}>
                  <Button
                    variant="filled"
                    fullWidth
                    type="submit"
                    loading={isPending}
                  >
                    Masuk
                  </Button>
                </Grid.Col>
              </Grid>
            </form>
          </Stack>
        </Stack>
      </Stack>
    </Container>
  );
}
