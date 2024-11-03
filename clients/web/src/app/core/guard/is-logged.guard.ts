import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";
import {Store} from "@ngxs/store";
import {AuthState} from "../../state/auth/auth.state";

export const isLoggedGuard: CanActivateFn = () => {
    const store = inject(Store)
    const state = store.selectSnapshot(AuthState.getState)
    const router = inject(Router)

    if (state.token !== null) {
        return router.createUrlTree(['/'])
    }

    return true;
};
