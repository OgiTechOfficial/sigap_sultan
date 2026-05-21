"use client";

import {
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { MantineProvider } from "@mantine/core";
import { Notifications } from "@mantine/notifications";
import { AuthConsumer, AuthProvider } from "@/contexts/auth";

// Suppress recharts defaultProps console warnings
if (typeof window !== "undefined") {
  const originalError = console.error;
  console.error = (...args: any[]) => {
    if (
      args[0] &&
      typeof args[0] === "string" &&
      args[0].includes("Support for defaultProps will be removed")
    ) {
      return;
    }
    originalError(...args);
  };
}

// Create a client
const queryCache = new QueryCache();
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
    },
  },
  queryCache,
});

export const Layout = (props: React.PropsWithChildren<{}>) => {
  const { children } = props;

  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider>
        <AuthProvider>
          <AuthConsumer>
            {(auth) => {
              if (!auth.isInitialized) return null;
              return <>{children}</>;
            }}
          </AuthConsumer>
        </AuthProvider>
        <Notifications />
      </MantineProvider>
    </QueryClientProvider>
  );
};
