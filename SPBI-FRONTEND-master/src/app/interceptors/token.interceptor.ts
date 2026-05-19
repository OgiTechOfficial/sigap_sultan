import {HttpInterceptorFn} from '@angular/common/http';
import {environment} from 'src/environments/environment';

export const tokenInterceptor: HttpInterceptorFn = (req, next) => {
  const token = localStorage.getItem('token');
  let url = req.url;
  
  if (!url.startsWith('http://') && !url.startsWith('https://') && !url.startsWith('assets/')) {
    url = `${environment.apiUrl}/${url.replace(/^\//, '')}`;
  }
  
  const headers: { [header: string]: string } = {};
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  req = req.clone({
    url,
    setHeaders: headers
  });
  
  return next(req);
};
