import {Route} from '@angular/router';
import {provideStates} from "@ngxs/store";
import {ChatState} from "./state/chat/chat.state";

export const appRoutes: Route[] = [
    {
        path: 'auth',
        loadComponent: () => import('./layouts/auth-layout/auth-layout.component').then(c => c.AuthLayoutComponent),
        children: [
            {
                path: 'sign-in',
                loadComponent: () => import('./pages/sign-in-page/sign-in-page.component').then(c => c.SignInPageComponent)
            },
            {
                path: 'sign-up',
                loadComponent: () => import('./pages/sign-up-page/sign-up-page.component').then(c => c.SignUpPageComponent)
            }
        ]
    },
    {
        path: '',
        loadComponent: () => import('./layouts/default-layout/default-layout.component').then(c => c.DefaultLayoutComponent),
        providers: [
            provideStates([ChatState])
        ],
        children: [
            {
                path: '',
                loadComponent: () => import('./pages/chat-page/chat-page.component').then(c => c.ChatPageComponent)
            }
        ]
    },
    {
        path: '**',
        pathMatch: 'full',
        redirectTo: '/',
    }
];
