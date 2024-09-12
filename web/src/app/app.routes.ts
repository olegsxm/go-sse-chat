import { Route } from '@angular/router';
import { AuthGuard } from './core/guards/auth-guard';

export const appRoutes: Route[] = [
  {
    path: 'auth',
    loadComponent: () =>
      import('./layouts/auth-layout/auth-layout.component').then(
        (m) => m.AuthLayoutComponent
      ),
    children: [],
  },
  {
    path: '',
    loadComponent: () =>
      import('./layouts/chat-layout/chat-layout.component').then(
        (c) => c.ChatLayoutComponent
      ),
    canActivate: [AuthGuard()],
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/dialog-page/dialog-page.component').then(
            (c) => c.DialogPageComponent
          ),
      },
    ],
  },
  {
    path: '**',
    pathMatch: 'full',
    redirectTo: '/',
  },
];
