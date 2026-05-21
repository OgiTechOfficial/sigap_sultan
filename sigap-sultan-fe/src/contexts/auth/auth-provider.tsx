import type { FC, ReactNode } from "react";
import { useCallback, useEffect, useReducer } from "react";
import PropTypes from "prop-types";

import type { UserAccount } from "@/types/user-account";

import type { State } from "./auth-context";
import { AuthContext, initialState } from "./auth-context";
import {
  ACCESS_TOKEN_STORAGE_KEY,
  REFRESH_TOKEN_STORAGE_KEY,
  USER_STORAGE_KEY,
} from "@/constants/storage";

enum ActionType {
  INITIALIZE = "INITIALIZE",
  SIGN_IN = "SIGN_IN",
  SIGN_UP = "SIGN_UP",
  SIGN_OUT = "SIGN_OUT",
}

type InitializeAction = {
  type: ActionType.INITIALIZE;
  payload: {
    isAuthenticated: boolean;
    user: UserAccount | null;
  };
};

type SignInAction = {
  type: ActionType.SIGN_IN;
  payload: {
    user: UserAccount;
  };
};

type SignUpAction = {
  type: ActionType.SIGN_UP;
  payload: {
    user: UserAccount;
  };
};

type SignOutAction = {
  type: ActionType.SIGN_OUT;
};

type Action = InitializeAction | SignInAction | SignUpAction | SignOutAction;

type Handler = (state: State, action: any) => State;

const handlers: Record<ActionType, Handler> = {
  INITIALIZE: (state: State, action: InitializeAction): State => {
    const { isAuthenticated, user } = action.payload;

    return {
      ...state,
      isAuthenticated,
      isInitialized: true,
      user,
    };
  },
  SIGN_IN: (state: State, action: SignInAction): State => {
    const { user } = action.payload;

    return {
      ...state,
      isAuthenticated: true,
      user,
    };
  },
  SIGN_UP: (state: State, action: SignUpAction): State => {
    const { user } = action.payload;

    return {
      ...state,
      isAuthenticated: true,
      user,
    };
  },
  SIGN_OUT: (state: State): State => ({
    ...state,
    isAuthenticated: false,
    user: null,
  }),
};

const reducer = (state: State, action: Action): State =>
  handlers[action.type] ? handlers[action.type](state, action) : state;

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: FC<AuthProviderProps> = (props) => {
  const { children } = props;
  const [state, dispatch] = useReducer(reducer, initialState);

  const initialize = useCallback(async (): Promise<void> => {
    try {
      const user = window.sessionStorage.getItem(USER_STORAGE_KEY);

      if (user) {
        dispatch({
          type: ActionType.INITIALIZE,
          payload: {
            isAuthenticated: true,
            user: JSON.parse(user),
          },
        });
      } else {
        dispatch({
          type: ActionType.INITIALIZE,
          payload: {
            isAuthenticated: false,
            user: null,
          },
        });
      }
    } catch (err) {
      console.error(err);
      dispatch({
        type: ActionType.INITIALIZE,
        payload: {
          isAuthenticated: false,
          user: null,
        },
      });
    }
  }, [dispatch]);

  useEffect(
    () => {
      initialize();
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const signIn = useCallback(
    async (
      userData: UserAccount,
      accessToken: string,
      refreshToken: string
    ): Promise<void> => {
      sessionStorage.setItem(ACCESS_TOKEN_STORAGE_KEY, accessToken);
      sessionStorage.setItem(REFRESH_TOKEN_STORAGE_KEY, refreshToken);
      sessionStorage.setItem(USER_STORAGE_KEY, JSON.stringify(userData));

      dispatch({
        type: ActionType.SIGN_IN,
        payload: {
          user: userData,
        },
      });
    },
    [dispatch]
  );

  const signOut = useCallback(async (): Promise<void> => {
    sessionStorage.removeItem(ACCESS_TOKEN_STORAGE_KEY);
    sessionStorage.removeItem(REFRESH_TOKEN_STORAGE_KEY);
    sessionStorage.removeItem(USER_STORAGE_KEY);
    dispatch({ type: ActionType.SIGN_OUT });
  }, [dispatch]);

  return (
    <AuthContext.Provider
      value={{
        ...state,
        signIn,
        signOut,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
