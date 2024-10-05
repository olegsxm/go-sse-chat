import {HttpInterceptorFn, HttpResponse} from '@angular/common/http';
import {environment} from "../../../environments/environment";
import {of} from "rxjs";

export const apiInterceptor: HttpInterceptorFn = (req, next) => {
    console.log("API interceptor called");

    if (req.url.startsWith("/auth")) {
        const {login} = req.body as { login: string }
        return of(new HttpResponse({status: 200, body: {token: '4343434223 token', user: {id: 1, login: login}}}));
    }

    return next(req.clone({
        url: `${environment.api}${req.url}`,
    }));
};
