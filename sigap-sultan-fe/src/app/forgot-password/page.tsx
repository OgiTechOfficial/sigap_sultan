"use client";

import {
  Stack,
  Text,
  Container,
  Grid,
  TextInput,
  Button,
  rem,
} from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";
import { useForm } from "@mantine/form";
import { useRouter } from "next/navigation";
import { IconMail } from "@tabler/icons-react";
import { APIResult } from "@/utils/api/handle-response";
import { useMutation } from "@tanstack/react-query";
import {
  authApi,
  AuthForgotPasswordRequest,
  AuthForgotPasswordResponse,
} from "@/api/base/auth";
import { notifications } from "@mantine/notifications";

type ForgotPasswordForm = {
  email: string;
};

export default function ForgotPassword() {
  const router = useRouter();

  const form = useForm<ForgotPasswordForm>({
    mode: "uncontrolled",
    initialValues: {
      email: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Format email salah"),
    },
  });

  const onForgotPassword = async (form: ForgotPasswordForm) => {
    const { result, error, displayMessage } = await mutateAsync({
      email: form.email,
    });

    if (error || !result) {
      return notifications.show({
        message: displayMessage ?? "Failed to fetch forgot password",
        color: "red",
        autoClose: 5000,
      });
    }

    router.push(result.link);
  };

  const { mutateAsync, isPending } = useMutation<
    APIResult<AuthForgotPasswordResponse>,
    Error,
    AuthForgotPasswordRequest
  >({
    mutationKey: ["auth-forgot-password"],
    mutationFn: (payload) => authApi.doForgotPassword(payload),
  });

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
              {"Lupa Password"}
            </Text>
            <Text size="sm" c="gray" ta="center">
              {
                "Masukan email kamu & kami akan mengirimkan link untuk mengubah password kamu."
              }
            </Text>
            <form onSubmit={form.onSubmit(onForgotPassword)}>
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
                  <Button
                    variant="filled"
                    fullWidth
                    type="submit"
                    loading={isPending}
                  >
                    Kirim Link
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
