"use client";

import { PropsWithChildren, ErrorInfo } from "react";
import {
  ErrorBoundary as BaseErrorBoundary,
  FallbackProps,
} from "react-error-boundary";
import ErrorPage from "../ErrorPage";

const logError = (error: Error, _: ErrorInfo) => {
  if (process.env.NODE_ENV === "development") {
    console.error(error);
  }
};

function ErrorBoundary(props: PropsWithChildren<{}>) {
  return (
    <BaseErrorBoundary
      FallbackComponent={(_: FallbackProps) => (
        <ErrorPage
          title="Server kami sedang padat merayap Harap, Harap Kembali lagi nanti!"
          description="Hi, Halaman ini tidak dapat ditemukan. harap kembali ke halaman sebelumnya!"
        />
      )}
      onError={logError}
    >
      {props.children}
    </BaseErrorBoundary>
  );
}

export default ErrorBoundary;
