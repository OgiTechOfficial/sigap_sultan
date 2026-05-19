export interface INav {
  name: string;             // Nama menu
  path: string;             // Path URL, tanpa domain. Contoh: /main/welcome
  privileges: string[];     // List privilege yang berhak akses menu
  active: boolean;          // Untuk set class active di menu navbar
}
