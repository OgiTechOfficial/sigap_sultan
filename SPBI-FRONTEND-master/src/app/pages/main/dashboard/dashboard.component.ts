import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

interface RegionData {
  id: number;
  name: string;
  logo: string;
  priceChangePercent: number;
  priceChangeRupiah: number;
  stockTonnage: number;
  stockStatus: 'Aman' | 'Waspada' | 'Defisit' | 'Rentan';
  color: string;
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  activeTab: 'harga' | 'neraca' | 'laporan' = 'harga';

  // Filters (Tabel Harga)
  komoditasSelected: string = 'Beras';
  jenisInfoHarga: string = 'Perubahan Harga (%)';
  detailInfoHarga: string = 'Bulan ke Bulan (MtM)';
  tanggalSelected: string = 'March 11, 2025';
  showDatePicker: boolean = false;
  sortSelected: string = 'Default';

  // Filters (Tabel Neraca)
  lihatBerdasarkan: string = 'Komoditas';
  jenisInfoNeraca: string = 'Neraca (Stok Akhir)';
  bulanSelected: string = 'October 2025';
  showMonthPicker: boolean = false;
  daerahSelected: string = 'Semua Daerah';

  // Dropdown lists
  commodities: string[] = ['Beras', 'Bawang Merah', 'Cabai Rawit', 'Minyak Goreng', 'Gula Pasir'];
  infoHargaTypes: string[] = ['Perubahan Harga (%)', 'Harga Konsumen (Eceran)'];
  detailHargaTypes: string[] = ['Bulan ke Bulan (MtM)', 'Tahun ke Tahun (YoY)'];
  infoNeracaTypes: string[] = ['Neraca (Stok Akhir)', 'Ketersediaan vs Kebutuhan'];
  daerahList: string[] = ['Semua Daerah', 'Kota Makassar', 'Kabupaten Gowa', 'Kabupaten Maros'];

  // Calendar dates
  months: string[] = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  years: number[] = [2024, 2025, 2026];
  selectedYear: number = 2025;

  // Region data array (24 districts of South Sulawesi)
  regions: RegionData[] = [
    {
      id: 1,
      name: 'Kota Makassar',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/f/f0/Lambang_Kota_Makassar.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3180.56,
      stockTonnage: 100047,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 2,
      name: 'Kabupaten Gowa',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/e/ec/Logo_Kabupaten_Gowa.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3110.12,
      stockTonnage: 120560,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 3,
      name: 'Kabupaten Maros',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/6/64/Lambang_Kabupaten_Maros.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3020.45,
      stockTonnage: 89450,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 4,
      name: 'Kota Parepare',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/b/b2/Logo_Kota_Parepare.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3202.22,
      stockTonnage: 52545,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 5,
      name: 'Kabupaten Wajo',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/d/de/Lambang_Kabupaten_Wajo.png',
      priceChangePercent: -100,
      priceChangeRupiah: -2944.44,
      stockTonnage: 386872,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 6,
      name: 'Kabupaten Bulukumba',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/d/df/Logo_Bulukumba.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3066.67,
      stockTonnage: 52545,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 7,
      name: 'Kabupaten Bantaeng',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/0/07/Lambang_Kabupaten_Bantaeng.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3120.00,
      stockTonnage: 31279,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 8,
      name: 'Kabupaten Kepulauan Selayar',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/3/36/Lambang_Kabupaten_Kepulauan_Selayar.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3127.83,
      stockTonnage: 13174,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 9,
      name: 'Kabupaten Luwu',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/e/ea/Lambang_Kabupaten_Luwu.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3150.00,
      stockTonnage: 37613,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 10,
      name: 'Kabupaten Pinrang',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/d/df/Logo_Pinrang.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3400.94,
      stockTonnage: 234489,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 11,
      name: 'Kabupaten Sinjai',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/1/13/Lambang_Kabupaten_Sinjai.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3111.11,
      stockTonnage: 287705,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 12,
      name: 'Kabupaten Soppeng',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/8/87/Lambang_Kabupaten_Soppeng.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3050.00,
      stockTonnage: 65123,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 13,
      name: 'Kabupaten Sidenreng Rappang',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/3/32/Logo_Sidenreng_Rappang.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3250.00,
      stockTonnage: 96927,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 14,
      name: 'Kabupaten Bone',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/4/4e/Logo_Kabupaten_Bone.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3300.00,
      stockTonnage: 412500,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 15,
      name: 'Kabupaten Jeneponto',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/c/cd/Logo_Jeneponto.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3180.00,
      stockTonnage: 210085,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 16,
      name: 'Kabupaten Takalar',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/d/db/Lambang_Kabupaten_Takalar.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3140.00,
      stockTonnage: 75300,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 17,
      name: 'Kabupaten Barru',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/5/5a/Lambang_Kabupaten_Barru.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3100.00,
      stockTonnage: 68900,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 18,
      name: 'Kabupaten Pangkajene Kepulauan',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/3/3a/Lambang_Kabupaten_Pangkajene_dan_Kepulauan.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3200.00,
      stockTonnage: 110400,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 19,
      name: 'Kabupaten Enrekang',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/8/8f/Lambang_Kabupaten_Enrekang.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3280.00,
      stockTonnage: 54120,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 20,
      name: 'Kabupaten Tana Toraja',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/d/d4/Lambang_Kabupaten_Tana_Toraja.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3500.00,
      stockTonnage: 43200,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 21,
      name: 'Kabupaten Toraja Utara',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/7/75/Lambang_Kabupaten_Toraja_Utara.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3480.00,
      stockTonnage: 39500,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 22,
      name: 'Kabupaten Luwu Utara',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/e/ea/Lambang_Kabupaten_Luwu_Utara.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3120.00,
      stockTonnage: 182400,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 23,
      name: 'Kabupaten Luwu Timur',
      logo: 'https://upload.wikimedia.org/wikipedia/commons/e/e0/Lambang_Kabupaten_Luwu_Timur.png',
      priceChangePercent: -100,
      priceChangeRupiah: -3080.00,
      stockTonnage: 154700,
      stockStatus: 'Aman',
      color: '#15803d'
    },
    {
      id: 24,
      name: 'Kabupaten Bone Bolango', // Placeholder for 24th
      logo: 'https://upload.wikimedia.org/wikipedia/commons/e/e4/Coat_of_arms_of_South_Sulawesi.svg',
      priceChangePercent: -100,
      priceChangeRupiah: -3200.00,
      stockTonnage: 45000,
      stockStatus: 'Aman',
      color: '#15803d'
    }
  ];

  constructor(private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
    // Listen to query params for tab switching
    this.route.queryParams.subscribe(params => {
      const tab = params['tab'];
      if (tab === 'harga' || tab === 'neraca' || tab === 'laporan') {
        this.activeTab = tab;
      } else {
        this.activeTab = 'harga';
      }
    });
  }

  setTab(tab: 'harga' | 'neraca' | 'laporan') {
    this.activeTab = tab;
    this.router.navigate([], {
      relativeTo: this.route,
      queryParams: { tab: tab },
      queryParamsHandling: 'merge'
    });
  }

  // Formatting helpers
  formatRupiah(val: number): string {
    const formatted = Math.abs(val).toLocaleString('id-ID', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    return val < 0 ? `-Rp ${formatted}` : `Rp ${formatted}`;
  }

  formatTonnage(val: number): string {
    return val.toLocaleString('id-ID') + ' ton';
  }

  // Date and month pickers toggle
  toggleDatePicker() {
    this.showDatePicker = !this.showDatePicker;
    this.showMonthPicker = false;
  }

  toggleMonthPicker() {
    this.showMonthPicker = !this.showMonthPicker;
    this.showDatePicker = false;
  }

  selectDate(day: number) {
    this.tanggalSelected = `March ${day}, 2025`;
    this.showDatePicker = false;
  }

  selectMonth(month: string) {
    this.bulanSelected = `${month} ${this.selectedYear}`;
    this.showMonthPicker = false;
  }

  applyFilters() {
    // Mock action to apply filters
    console.log('Filters applied');
  }

  clearKomoditas() {
    this.komoditasSelected = '';
  }
}
