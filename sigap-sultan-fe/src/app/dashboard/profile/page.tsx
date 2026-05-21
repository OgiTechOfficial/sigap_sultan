"use client";

import {
  authApi,
  AuthChangePasswordRequest,
  AuthChangePasswordResponse,
} from "@/api/base/auth";
import { positionApi } from "@/api/base/position";
import {
  profileApi,
  UpdateProfileRequest,
  UpdateProfileResponse,
} from "@/api/base/profile";
import { OptionMap } from "@/types/option";
import { APIResult } from "@/utils/api/handle-response";
import {
  Button,
  Card,
  ComboboxData,
  Container,
  Grid,
  Group,
  Modal,
  PasswordInput,
  Select,
  Stack,
  Text,
  TextInput,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { notifications } from "@mantine/notifications";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useState } from "react";

type ProfileForm = {
  username: string;
  email: string;
  nama_depan: string;
  nama_belakang: string;
  organisasi: string;
  bidang_unit_kerja: string;
  jabatan: OptionMap<string> | null;
};

const initialProfileForm: ProfileForm = {
  username: "",
  email: "",
  nama_depan: "",
  nama_belakang: "",
  organisasi: "",
  bidang_unit_kerja: "",
  jabatan: null,
};

type ChangePasswordForm = {
  password: string;
  newPassword: string;
  newPasswordConfirm: string;
};

const initialChangePasswordForm: ChangePasswordForm = {
  password: "",
  newPassword: "",
  newPasswordConfirm: "",
};

export default function Profile() {
  const [changePasswordOpened, setChangePasswordOpened] = useState(false);
  const profileForm = useForm<ProfileForm>({
    mode: "controlled",
    initialValues: initialProfileForm,
    validate: {
      username: (value) => (!value ? "Username harus diisi" : null),
      email: (value) => (!value ? "Email harus diisi" : null),
      nama_depan: (value) => (!value ? "Nama depan harus diisi" : null),
      nama_belakang: (value) => (!value ? "Nama belakang harus diisi" : null),
      organisasi: (value) => (!value ? "Organisasi harus diisi" : null),
      bidang_unit_kerja: (value) => (!value ? "Unit kerja harus diisi" : null),
      jabatan: (value) => (!value ? "Jabatan harus diisi" : null),
    },
  });
  const changePasswordForm = useForm<ChangePasswordForm>({
    mode: "controlled",
    initialValues: initialChangePasswordForm,
    validate: {
      password: (value) => (!value ? "Password saat ini harus diisi" : null),
      newPassword: (value, values) =>
        values.newPasswordConfirm === value ? null : "Password tidak sama",
      newPasswordConfirm: (value, values) =>
        values.newPassword === value ? null : "Password tidak sama",
    },
  });

  const { data: profileData } = useQuery({
    queryKey: ["profile"],
    queryFn: async () => profileApi.getDetail(),
    select: (data) => data.result,
  });

  const { data: positionData = [] } = useQuery({
    queryKey: ["position"],
    queryFn: async () => positionApi.getList(),
    select: (data) => data.result || [],
  });

  const {
    mutateAsync: mutateAsyncUpdateProfile,
    isPending: isPendingUpdateProfile,
  } = useMutation<
    APIResult<UpdateProfileResponse>,
    Error,
    UpdateProfileRequest
  >({
    mutationKey: ["update-profile"],
    mutationFn: (payload) => profileApi.update(payload),
  });

  const {
    mutateAsync: mutateAsyncChangePassword,
    isPending: isPendingChangePassword,
  } = useMutation<
    APIResult<AuthChangePasswordResponse>,
    Error,
    AuthChangePasswordRequest
  >({
    mutationKey: ["change-password"],
    mutationFn: (payload) => authApi.doChangePassword(payload),
  });

  const [prevProfileData, setPrevProfileData] = useState(profileData);
  if (JSON.stringify(prevProfileData) !== JSON.stringify(profileData)) {
    setPrevProfileData(profileData);
    profileForm.setValues({
      username: profileData?.username || "",
      email: profileData?.email || "",
      nama_depan: profileData?.nama_depan || "",
      nama_belakang: profileData?.nama_belakang || "",
      organisasi: profileData?.organisasi || "",
      bidang_unit_kerja: profileData?.bidang_unit_kerja || "",
      jabatan: profileData?.jabatan_id
        ? {
            value: String(profileData.jabatan_id),
            label: profileData.jabatan,
          }
        : null,
    });
  }

  const handleProfileSubmit = async (form: ProfileForm) => {
    const { result, error, displayMessage } = await mutateAsyncUpdateProfile({
      id: profileData?.id || null,
      username: form.username,
      email: form.email,
      nama_depan: form.nama_depan,
      nama_belakang: form.nama_belakang,
      organisasi: form.organisasi,
      bidang_unit_kerja: form.bidang_unit_kerja,
      jabatan: form.jabatan?.label || "",
      jabatan_id: form.jabatan?.value ? Number(form.jabatan.value) : null,
    });

    if (error || !result) {
      return notifications.show({
        message: displayMessage ?? "Failed to update profile",
        color: "red",
        autoClose: 5000,
      });
    }

    notifications.show({
      message: "Profile updated",
      color: "green",
      autoClose: 5000,
    });
  };

  const handleChangePasswordSubmit = async (form: ChangePasswordForm) => {
    const { result, error, displayMessage } = await mutateAsyncChangePassword({
      oldPassword: form.password,
      newPassword: form.newPassword,
      newPasswordConfirm: form.newPasswordConfirm,
    });

    if (error || !result) {
      return notifications.show({
        message: displayMessage ?? "Failed to change password",
        color: "red",
        autoClose: 5000,
      });
    }

    notifications.show({
      message: "Password changed",
      color: "green",
      autoClose: 5000,
    });
    setChangePasswordOpened(false);
  };

  return (
    <>
      <Container fluid p="lg" bg={"#F9FAFB"}>
        <form onSubmit={profileForm.onSubmit(handleProfileSubmit)}>
          <Group gap="md" justify="space-between">
            <Text size="lg" fw="bold">
              Profile
            </Text>
            <Group gap="md">
              <Button
                variant="outline"
                onClick={() => setChangePasswordOpened(true)}
              >
                Ubah Password
              </Button>
              <Button type="submit" loading={isPendingUpdateProfile}>
                Edit Profile
              </Button>
            </Group>
          </Group>
          <Stack gap={0} mt="md">
            <Text size="md" fw="bold">
              Informasi personal
            </Text>
            <Text size="sm">Detail informasi personal tentang akun mu.</Text>
          </Stack>
          <Card padding="md" radius="md" withBorder mt="md">
            <Grid pb="sm">
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="NIP/Username"
                  placeholder="Masukan NIP/Username"
                  key={profileForm.key("username")}
                  {...profileForm.getInputProps("username")}
                  value={profileForm.getInputProps("username").value}
                  onChange={(e) => {
                    profileForm
                      .getInputProps("username")
                      .onChange(e.target.value);
                  }}
                />
              </Grid.Col>
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="Email"
                  placeholder="Masukan Email"
                  key={profileForm.key("email")}
                  {...profileForm.getInputProps("email")}
                  value={profileForm.getInputProps("email").value}
                  onChange={(e) => {
                    profileForm.getInputProps("email").onChange(e.target.value);
                  }}
                />
              </Grid.Col>
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="Nama Depan"
                  placeholder="Masukan Nama Depan"
                  key={profileForm.key("nama_depan")}
                  {...profileForm.getInputProps("nama_depan")}
                  value={profileForm.getInputProps("nama_depan").value}
                  onChange={(e) => {
                    profileForm
                      .getInputProps("nama_depan")
                      .onChange(e.target.value);
                  }}
                />
              </Grid.Col>
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="Nama Belakang (Opsional)"
                  placeholder="Masukan Nama Belakang"
                  key={profileForm.key("nama_belakang")}
                  {...profileForm.getInputProps("nama_belakang")}
                  value={profileForm.getInputProps("nama_belakang").value}
                  onChange={(e) => {
                    profileForm
                      .getInputProps("nama_belakang")
                      .onChange(e.target.value);
                  }}
                />
              </Grid.Col>
            </Grid>
          </Card>
          <Stack gap={0} mt="md">
            <Text size="md" fw="bold">
              Informasi Region
            </Text>
            <Text size="sm">Detail informasi personal tentang akun mu.</Text>
          </Stack>
          <Card padding="md" radius="md" withBorder mt="md">
            <Grid pb="sm">
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="OPD/Organisasi"
                  placeholder="Masukan OPD/Organisasi"
                  key={profileForm.key("organisasi")}
                  {...profileForm.getInputProps("organisasi")}
                  value={profileForm.getInputProps("organisasi").value}
                  onChange={(e) => {
                    profileForm
                      .getInputProps("organisasi")
                      .onChange(e.target.value);
                  }}
                />
              </Grid.Col>
              <Grid.Col span={{ sm: 12, md: 6 }}>
                <TextInput
                  label="Bidang Unit Kerja"
                  placeholder="Masukan Bidang Unit Kerja"
                  key={profileForm.key("bidang_unit_kerja")}
                  {...profileForm.getInputProps("bidang_unit_kerja")}
                  value={profileForm.getInputProps("bidang_unit_kerja").value}
                  onChange={(e) => {
                    profileForm
                      .getInputProps("bidang_unit_kerja")
                      .onChange(e.target.value);
                  }}
                />
              </Grid.Col>
              <Grid.Col span={{ sm: 12, md: 12 }}>
                <Select
                  label="Jabatan"
                  placeholder="Masukan Jabatan"
                  data={positionData.map((position) => ({
                    value: String(position.id),
                    label: position.name,
                  }))}
                  key={profileForm.key("jabatan")}
                  {...profileForm.getInputProps("jabatan")}
                  value={profileForm.getInputProps("jabatan").value}
                  onChange={(value) => {
                    profileForm.getInputProps("jabatan").onChange(value);
                  }}
                />
              </Grid.Col>
            </Grid>
          </Card>
        </form>
      </Container>
      <Modal
        size={641}
        opened={changePasswordOpened}
        onClose={() => setChangePasswordOpened(false)}
        withCloseButton={false}
        centered
      >
        <Stack gap={0} align="center">
          <Text size="md" fw="bold">
            Ubah Password
          </Text>
          <Text size="sm">Masukan password dan ubah passwordmu</Text>
        </Stack>
        <form
          onSubmit={changePasswordForm.onSubmit(handleChangePasswordSubmit)}
        >
          <Grid pb="sm">
            <Grid.Col span={{ sm: 12, md: 12 }}>
              <PasswordInput
                label="Password Saat Ini"
                placeholder="Masukan password saat ini"
                key={changePasswordForm.key("password")}
                {...changePasswordForm.getInputProps("password")}
                value={changePasswordForm.getInputProps("password").value}
                onChange={(e) => {
                  changePasswordForm
                    .getInputProps("password")
                    .onChange(e.target.value);
                }}
              />
            </Grid.Col>
            <Grid.Col span={{ sm: 12, md: 12 }}>
              <PasswordInput
                label="Password Baru"
                placeholder="Masukan password baru"
                key={changePasswordForm.key("newPassword")}
                {...changePasswordForm.getInputProps("newPassword")}
                value={changePasswordForm.getInputProps("newPassword").value}
                onChange={(e) => {
                  changePasswordForm
                    .getInputProps("newPassword")
                    .onChange(e.target.value);
                }}
              />
            </Grid.Col>
            <Grid.Col span={{ sm: 12, md: 12 }}>
              <PasswordInput
                label="Konfirmasi Password Baru"
                placeholder="Masukan konfirmasi password baru"
                key={changePasswordForm.key("newPasswordConfirm")}
                {...changePasswordForm.getInputProps("newPasswordConfirm")}
                value={
                  changePasswordForm.getInputProps("newPasswordConfirm").value
                }
                onChange={(e) => {
                  changePasswordForm
                    .getInputProps("newPasswordConfirm")
                    .onChange(e.target.value);
                }}
              />
            </Grid.Col>
          </Grid>
          <Stack gap="sm" mt="md">
            <Button
              variant="outline"
              fullWidth
              onClick={() => setChangePasswordOpened(false)}
            >
              Batal
            </Button>
            <Button type="submit" fullWidth loading={isPendingChangePassword}>
              Change Password
            </Button>
          </Stack>
        </form>
      </Modal>
    </>
  );
}
