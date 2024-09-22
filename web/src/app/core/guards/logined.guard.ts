import { CanActivateFn } from '@angular/router';
import { inject } from '@angular/core';
import { AuthService } from '../services/auth/auth.service';

export const loginedGuard: CanActivateFn = (route, state) => {
  const service = inject(AuthService);

  return service.token === null;
};
