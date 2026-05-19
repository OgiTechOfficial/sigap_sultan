import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { environment } from 'src/environments/environment';
import { ILoginPayload } from '../pages/login/login.interface';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private env = environment;

  constructor(
    private http: HttpClient
  ) { }

  login(data: ILoginPayload): Observable<any> {
    // Call API here, but for now let's return it as success
    // return this.http.post(`${this.env}/login`, data);
    return of(null);
  }
}
