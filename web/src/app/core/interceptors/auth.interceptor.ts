import {HttpInterceptorFn} from '@angular/common/http';
import {inject} from "@angular/core";
import {AuthState} from "../../state/auth/auth.state";

export const authInterceptor: HttpInterceptorFn = (req, next) => {
    const state = inject(AuthState)
    

    return next(req);
};
