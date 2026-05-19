import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {

  constructor(private http: HttpClient) { }

  /**
   * Mengambil semua daftar komoditas pangan
   */
  getCommodities(): Observable<any> {
    return this.http.get('commodities');
  }

  /**
   * Mengambil tingkat harga sebaran komoditas di seluruh kota/kabupaten
   */
  getPrices(commodityId: string, date: string): Observable<any> {
    return this.http.get(`price/level-harga/list?commodityId=${commodityId}&selectedDate=${date}`);
  }

  /**
   * Mengambil stok neraca pangan akhir di seluruh kota/kabupaten
   */
  getNeraca(commodityId: string, date: string): Observable<any> {
    return this.http.get(`neraca/stock-akhir-by-commodity/list?commodityId=${commodityId}&selectedDate=${date}`);
  }

  /**
   * Mengambil tren pergerakan historis komoditas dalam 5 hari terakhir
   */
  getHistoricalTrend(commodityId: string): Observable<any> {
    return this.http.get(`price/last-5-days-by-commodity?commodityId=${commodityId}`);
  }
}
