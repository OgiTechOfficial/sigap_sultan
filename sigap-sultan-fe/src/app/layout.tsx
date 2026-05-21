import type { Metadata, Viewport } from "next";
import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";
import "@mantine/dates/styles.css";
import "@mantine/charts/styles.css";
import "mantine-react-table/styles.css";
import { Layout } from "./layouts/root";
import { assetPrefix } from "@/utils/asset-prefix";

export const metadata: Metadata = {
  title: "Bank Indonesia - Sigap Sultan",
  description: "Sistem Informasi Harga dan Pasokan Pangan Sulawesi Selatan",
  generator: "Next.js",
  manifest: assetPrefix("/manifest.json"),
  keywords: [
    "bank indonesia",
    "bi",
    "sigap",
    "sigap sultan",
    "bank indonesia sigap sultan",
  ],
  authors: [{ name: "Bank Infonesia" }],
  icons: [
    { rel: "apple-touch-icon", url: "images/icons/icon-128x128.png" },
    { rel: "icon", url: "images/icons/icon-128x128.png" },
  ],
};

export const viewport: Viewport = {
  themeColor: [{ media: "(prefers-color-scheme: dark)", color: "#fff" }],
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="font-sans bg-white">
        <Layout>{children}</Layout>
      </body>
    </html>
  );
}
