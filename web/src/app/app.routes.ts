import {Route} from '@angular/router';
import {provideStates} from "@ngxs/store";
import {ChatState} from "./state/chat/chat.state";
import {isLoggedGuard} from "./core/guard/is-logged.guard";
import {authGuard} from "./core/guard/auth.guard";

export const appRoutes: Route[] = [
    {
        path: 'auth',
        loadComponent: () => import('./layouts/auth-layout/auth-layout.component').then(c => c.AuthLayoutComponent),
        canActivate: [isLoggedGuard],
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
            provideStates([ChatState]),
        ],
        canActivate: [authGuard],
        children: [
            {
                path: '',
                loadComponent: () => import('./pages/chat-page/chat-page.component').then(c => c.ChatPageComponent),
                children: [
                    {
                        path: '',
                        loadComponent: () => import('./pages/start-chat-page/start-chat-page.component')
                            .then(c => c.StartChatPageComponent)
                    },
                    {
                        path: ':conversation',
                        loadComponent: () => import('./pages/chat/chat.component').then(c => c.ChatComponent),
                    }
                ]
            }
        ]
    },
    {
        path: '**',
        pathMatch: 'full',
        redirectTo: '/',
    }
];
