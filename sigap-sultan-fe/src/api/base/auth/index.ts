import { baseFetcher } from "../fetcher";
import {
  apiChangePassword,
  apiForgotPassword,
  apiLogin,
  apiResetPassword,
  apiVerifyToken,
} from "../api-paths";

export interface AuthLoginResponse {
  accessible_menu: any | null;
  exp: number;
  id: number;
  iss: string;
  name: string;
  position: string;
  position_id: number;
  token: string;
}

export interface AuthLoginRequest {
  email: string;
  password: string;
}

export interface AuthForgotPasswordRequest {
  email: string;
}

export interface AuthForgotPasswordResponse {
  link: string;
}

export interface AuthResetPasswordResponse {}

export interface AuthResetPasswordRequest {
  newPassword: string;
  newPasswordConfirm: string;
  code: string | null;
}

export interface AuthChangePasswordResponse {}

export interface AuthChangePasswordRequest {
  oldPassword: string;
  newPassword: string;
  newPasswordConfirm: string;
}

export interface AuthVerifyTokenResponse {}

export interface AuthVerifyTokenRequest {
  code: string | null;
}

class AuthApi {
  doLogin(request: AuthLoginRequest) {
    return baseFetcher.apiPost<AuthLoginResponse>({
      path: `${apiLogin}`,
      payload: request,
    });
  }

  doForgotPassword(request: AuthForgotPasswordRequest) {
    return baseFetcher.apiPost<AuthForgotPasswordResponse>({
      path: `${apiForgotPassword}`,
      payload: request,
    });
  }

  doVerifyToken(request: AuthVerifyTokenRequest) {
    return baseFetcher.apiPost<AuthVerifyTokenResponse>({
      path: `${apiVerifyToken}`,
      payload: request,
    });
  }

  doResetPassword(request: AuthResetPasswordRequest) {
    return baseFetcher.apiPost<AuthResetPasswordResponse>({
      path: `${apiResetPassword}`,
      payload: request,
    });
  }

  doChangePassword(request: AuthChangePasswordRequest) {
    return baseFetcher.apiPost<AuthChangePasswordResponse>({
      path: `${apiChangePassword}`,
      payload: request,
    });
  }
}

export const authApi = new AuthApi();
