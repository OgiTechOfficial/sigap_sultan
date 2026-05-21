import { createContext } from "react";

import type { UserAccount } from "@/types/user-account";

export interface State {
  isInitialized: boolean;
  isAuthenticated: boolean;
  user: UserAccount | null;
}

export const initialState: State = {
  isAuthenticated: false,
  isInitialized: false,
  user: null,
};

export interface AuthContextType extends State {
  signIn: (
    userData: UserAccount,
    accessToken: string,
    refreshToken: string
  ) => Promise<void>;
  signOut: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextType>({
  ...initialState,
  signIn: () => Promise.resolve(),
  signOut: () => Promise.resolve(),
});
