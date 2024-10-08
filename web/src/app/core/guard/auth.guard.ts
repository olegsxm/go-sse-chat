import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";
import {AuthState} from "../../state/auth/auth.state";
import {Store} from "@ngxs/store";

export const authGuard: CanActivateFn = () => {
    const store = inject(Store)
    const state = store.selectSnapshot(AuthState.getState)
    const router = inject(Router)

    if (state.token !== null) return true;

    return router.createUrlTree(['auth/sign-in'])
};
