import React from "react";
import Navbar from "@/app/components/NavbarV2/Navbar";
import BottomNavbar from "@/app/components/NavbarV2/BottomNavbar";
import { Layout } from "../layouts/root";
import ErrorBoundary from "../components/ErrorBoundary/ErrorBoundary";
import { Box } from "@mantine/core";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <Layout>
      <Navbar />
      <ErrorBoundary>
        <Box pt={60}>{children}</Box>
      </ErrorBoundary>
      <BottomNavbar />
    </Layout>
  );
}
