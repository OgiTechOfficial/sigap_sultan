import Link from "next/link";
import Image from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

const TopNavbar = () => {
  return (
    <nav className="bg-top-navbar bg-cover flex items-center pt-4 pb-4">
      <div className="w-11/12 mx-auto flex items-center justify-between">
        <div className="shrink-0">
          <Image
            src={"/header_logo.svg"}
            alt={"Top header logo"}
            width={350}
            height={88}
            layout="responsive"
          />
        </div>
        <div className="flex-1 flex justify-center items-center gap-8">
          <Link href="/dashboard/price-table">
            <span className="text-sm text-white font-semibold hover:text-gray-200 cursor-pointer">
              Tabel Harga
            </span>
          </Link>
          <Link href="/dashboard/neraca">
            <span className="text-sm text-white font-semibold hover:text-gray-200 cursor-pointer">
              Tabel Neraca
            </span>
          </Link>
          <Link href="/dashboard/report">
            <span className="text-sm text-white font-semibold hover:text-gray-200 cursor-pointer">
              Unduh Laporan c
            </span>
          </Link>
        </div>
        <div className="items-center justify-end">
          <Link
            href={"/login"}
            className="text-custom-blue bg-white px-4 py-2 rounded border-2"
          >
            Login
          </Link>
        </div>
      </div>
    </nav>
  );
};

const BottomNavbar = () => {
  return (
    <div className="inset-x-0 bottom-0 bg-bottom-navbar bg-cover text-white text-sm pt-14 pb-10">
      <div className="w-11/12 mx-auto py-2 flex flex-col justify-between items-center">
        <div className="items-center">
          <Image
            src={assetPrefix("/logo/logo_bi_sulsel.svg")}
            alt="BI dan Sulsel"
            height="100"
            width="100"
          />
        </div>
        <div className="items-center mt-4">
          <Image
            src={assetPrefix("/logo/logo_sigap_sultan_hr.svg")}
            alt="BI dan Sulsel"
            height="200"
            width="200"
          />
        </div>
        <div className="text-center mt-4">
          <p className="font-bold">SISTEM INFORMASI HARGA DAN PASOKAN PANGAN</p>
          <p className="font-bold text-custom-yellow">SULAWESI SELATAN</p>
        </div>
        <div className="text-center mt-8">
          <p className="font-bold text-base">
            Sekretariat Tim Pengendalian Inflasi Provinsi Sulawesi Selatan.
          </p>
        </div>
        <div className="text-center">
          <span className="font-medium text-base">
            Jl. Jendral Urip Sumoharjo No.269, Panaikang, Kec. Panakkukang, Kota
            Makassar, Sulawesi Selatan |{" "}
            <span className="underline">sigapsultan@gmail.com</span>
          </span>
        </div>
        <div className="w-full border-b border-white mt-8 mb-6"></div>
      </div>
      <div className="w-11/12 mx-auto flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0">
        <div className="text-center md:text-left">
          <p className="font-thin text-sm md:text-base">
            © 2024 Sulsel. All rights reserved.
          </p>
        </div>
        <div className="text-center">
          <p className="font-bold text-sm md:text-base">
            sigapsultan@gmail.com
          </p>
        </div>
        <div className="text-center md:text-right">
          <p className="font-thin text-sm md:text-base">Powered by Sentech</p>
        </div>
      </div>
    </div>
  );
};

export { TopNavbar, BottomNavbar };
