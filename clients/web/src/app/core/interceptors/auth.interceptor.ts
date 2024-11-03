import {HttpHeaders, HttpInterceptorFn} from '@angular/common/http';
import {inject} from "@angular/core";
import {Store} from "@ngxs/store";
import {AuthState} from "../../state/auth/auth.state";

export const authInterceptor: HttpInterceptorFn = (req, next) => {
    const store = inject(Store)
    const authState = store.selectSnapshot(AuthState.getState);

    if (authState.token === null) {
        return next(req)
    }

    const r = req.clone({
        headers: new HttpHeaders({Authorization: `Bearer ${authState.token}`}),
    })

    return next(r);
};
