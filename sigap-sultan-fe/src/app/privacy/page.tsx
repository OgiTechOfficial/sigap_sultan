"use client";

import { useState } from "react";
import { Card, Grid, Container } from "@mantine/core";
import ReportFilter, {
  ReportFilterState,
} from "@/sections/dashboard/report/components/ReportFilter";
import ReportSummary from "@/sections/dashboard/report/components/ReportSummary";
import ReportResultNeraca from "@/sections/dashboard/report/components/ReportResultNeraca";
import ReportResultPrice from "@/sections/dashboard/report/components/ReportResultPrice";

export type PriceState = {
  priceCommodity: any[];
  priceDiff: any;
};

export default function ReportPage() {
  const [submittedReportFilterState, setSubmittedReportFilterState] =
    useState<ReportFilterState | null>(null);

  const handleSearch = (form: ReportFilterState) => {
    setSubmittedReportFilterState({
      ...form,
      requestTimestamp: new Date().getTime(),
    });
  };

  return (
      <Container fluid p="lg" bg={"#F9FAFB"}>
        <header>
          <h1>Kebijakan Privasi</h1>
        </header>
        <main>
          <p>Aplikasi <strong>TPID Sulsel - Sigap Sultan</strong> bersifat publik dan tidak mengumpulkan, menyimpan, atau
            memproses data pribadi pengguna.</p>

          <h2>1. Informasi yang Kami Kumpulkan</h2>
          <p>Kami tidak mengumpulkan informasi pribadi dalam bentuk apa pun. Pengguna dapat mengakses seluruh fitur dan informasi dalam aplikasi ini tanpa perlu mendaftar, login, atau memberikan data apa pun.</p>

          <h2>2. Cookie dan Pelacakan</h2>
          <p>Kami tidak menggunakan cookie, alat pelacakan pihak ketiga, atau metode analitik lain yang dapat mengidentifikasi pengguna secara individu.</p>

          <h2>3. Keamanan</h2>
          <p>Karena tidak ada data pribadi yang dikumpulkan, tidak ada informasi pengguna yang perlu diamankan. Namun, kami tetap menjaga integritas dan ketersediaan informasi yang ditampilkan dalam aplikasi ini.</p>

          <h2>4. Perubahan Kebijakan</h2>
          <p>Kebijakan privasi ini dapat diperbarui dari waktu ke waktu. Segala perubahan akan dipublikasikan di halaman ini.</p>

          <h2>5. Kontak</h2>
          <p>Jika Anda memiliki pertanyaan tentang kebijakan privasi ini, silakan hubungi kami melalui email:
            <a href="mailto:neracapangan.tpidsulsel@gmail.com">neracapangan.tpidsulsel@gmail.com</a>
          </p>
        </main>
      </Container>
  );
}
