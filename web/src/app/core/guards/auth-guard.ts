import { CanActivateFn, Router } from '@angular/router';
import { inject } from '@angular/core';
import { AuthService } from '../services/auth/auth.service';

export function AuthGuard(): CanActivateFn {
  return () => {
    const authService = inject(AuthService);
    const router = inject(Router);
    
    const a = !!authService.token;

    if (!a) {
      return router.createUrlTree(['/auth']);
    }

    return true;
  };
}
