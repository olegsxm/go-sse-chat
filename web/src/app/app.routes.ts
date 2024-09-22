import { Route } from '@angular/router';
import { AuthGuard } from './core/guards/auth-guard';
import { loginedGuard } from './core/guards/logined.guard';

export const appRoutes: Route[] = [
  {
    path: 'auth',
    canActivate: [loginedGuard],
    loadComponent: () =>
      import('./layouts/auth-layout/auth-layout.component').then(
        (m) => m.AuthLayoutComponent
      ),
    children: [
      {
        path: 'sign-in',
        loadComponent: () => import('./pages/sign-in-page/sign-in-page.component').then(c => c.SignInPageComponent)
      },
      {
        path: 'sign-up',
        loadComponent: () => import('./pages/sign-up-page/sign-up-page.component').then(c => c.SignUpPageComponent)
      },
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'sign-in'
      }
    ]
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
          )
      }
    ]
  },
  {
    path: '**',
    pathMatch: 'full',
    redirectTo: '/'
  }
];
