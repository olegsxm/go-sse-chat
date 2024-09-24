import { HttpInterceptorFn } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { inject } from '@angular/core';
import { AuthService } from '../services/auth/auth.service';

export const apiInterceptor: HttpInterceptorFn = (req, next) => {

  const authService = inject(AuthService);
  req = req.clone({
    url: `${environment.api}${req.url}`
  });


  if (authService.token) {
    req = req.clone({
      headers: req.headers.append('Authorization', 'Bearer ' + authService.token)
    });
  }

  return next(req);
};
