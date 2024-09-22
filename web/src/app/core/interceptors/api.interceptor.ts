import { HttpInterceptorFn } from '@angular/common/http';
import { environment } from '../../../environments/environment';

export const apiInterceptor: HttpInterceptorFn = (req, next) => {

  req = req.clone({
    url: `${environment.api}${req.url}`
  });

  return next(req);
};
