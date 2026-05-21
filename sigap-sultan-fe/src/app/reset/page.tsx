"use client";

import {
  Stack,
  Text,
  Container,
  Grid,
  Button,
  rem,
  PasswordInput,
} from "@mantine/core";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";
import { useForm } from "@mantine/form";
import { useRouter, useSearchParams } from "next/navigation";
import { IconLock } from "@tabler/icons-react";
import { APIResult } from "@/utils/api/handle-response";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  authApi,
  AuthResetPasswordRequest,
  AuthResetPasswordResponse,
} from "@/api/base/auth";
import { notifications } from "@mantine/notifications";
import LoadingPage from "../components/LoadingPage";
import ExpiredPage from "../components/ExpiredPage";

type ResetPasswordForm = {
  newPassword: string;
  newPasswordConfirm: string;
};

export default function ResetPassword() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const form = useForm<ResetPasswordForm>({
    mode: "uncontrolled",
    initialValues: {
      newPassword: "",
      newPasswordConfirm: "",
    },

    validate: {
      newPassword: (value, values) =>
        values.newPasswordConfirm === value ? null : "Password tidak sama",
      newPasswordConfirm: (value, values) =>
        values.newPassword === value ? null : "Password tidak sama",
    },
  });

  const onResetPassword = async (form: ResetPasswordForm) => {
    const { result, error, displayMessage } = await mutateAsync({
      newPassword: form.newPassword,
      newPasswordConfirm: form.newPasswordConfirm,
      code: searchParams.get("code"),
    });

    if (error || !result) {
      return notifications.show({
        message: displayMessage ?? "Failed to fetch reset password",
        color: "red",
        autoClose: 5000,
      });
    }

    router.push("/login");
  };

  const { mutateAsync, isPending } = useMutation<
    APIResult<AuthResetPasswordResponse>,
    Error,
    AuthResetPasswordRequest
  >({
    mutationKey: ["auth-reset-password"],
    mutationFn: (payload) => authApi.doResetPassword(payload),
  });

  const { data: dataVerifyToken, isFetching: isLoadingVerifyToken } = useQuery({
    queryKey: ["verify-token", searchParams.get("code")],
    queryFn: async () => {
      const { result, error, displayMessage } = await authApi.doVerifyToken({
        code: searchParams.get("code"),
      });

      if (error || !result) {
        throw new Error(displayMessage ?? "Failed to verify token");
      }

      return result;
    },
  });

  if (isLoadingVerifyToken) return <LoadingPage />;

  if (!dataVerifyToken)
    return (
      <ExpiredPage
        title="Link ubah password sudah expired, silahkan request kembali melalui button dibawah ini:"
        onPress={() => router.push("/forgot-password")}
      />
    );

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
              {"Buat Password Baru"}
            </Text>
            <Text size="sm" c="gray" ta="center">
              {"Masukan password baru kamu untuk login selanjutnya"}
            </Text>
            <form onSubmit={form.onSubmit(onResetPassword)}>
              <Grid pb="sm">
                <Grid.Col span={12}>
                  <PasswordInput
                    label="Passwor Baru"
                    placeholder="Masukan Password"
                    key={form.key("password")}
                    leftSection={
                      <IconLock style={{ width: rem(18), height: rem(18) }} />
                    }
                    {...form.getInputProps("newPassword")}
                    value={form.getInputProps("newPassword").value}
                    onChange={(e) => {
                      form
                        .getInputProps("newPassword")
                        .onChange(e.target.value);
                    }}
                  />
                </Grid.Col>
                <Grid.Col span={12}>
                  <PasswordInput
                    label="Konfirmasi Password Baru"
                    placeholder="Masukan Password"
                    key={form.key("password")}
                    leftSection={
                      <IconLock style={{ width: rem(18), height: rem(18) }} />
                    }
                    {...form.getInputProps("newPasswordConfirm")}
                    value={form.getInputProps("newPasswordConfirm").value}
                    onChange={(e) => {
                      form
                        .getInputProps("newPasswordConfirm")
                        .onChange(e.target.value);
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
                    Ubah Password
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
