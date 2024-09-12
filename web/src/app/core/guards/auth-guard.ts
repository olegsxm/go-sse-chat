import { CanActivateFn, Router } from '@angular/router';
import { inject } from '@angular/core';

export function AuthGuard(): CanActivateFn {
  console.log('auth guard');

  return () => {
    const a = true;

    if (!a) {
      const router = inject(Router);
      return router.createUrlTree(['/auth']);
    }

    return true;
  };
}
