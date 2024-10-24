import {ApplicationConfig, provideZoneChangeDetection} from '@angular/core';
import {provideRouter, withComponentInputBinding} from '@angular/router';
import {appRoutes} from './app.routes';
import {provideStore} from '@ngxs/store';
import {provideHttpClient, withFetch, withInterceptors} from "@angular/common/http";
import {apiInterceptor} from "./core/interceptors/api.interceptor";
import {authInterceptor} from "./core/interceptors/auth.interceptor";
import {AuthState} from "./state/auth/auth.state";
import {withNgxsReduxDevtoolsPlugin} from "@ngxs/devtools-plugin";

export const appConfig: ApplicationConfig = {
    providers: [
        provideZoneChangeDetection({eventCoalescing: true}),
        provideHttpClient(withFetch(), withInterceptors([apiInterceptor, authInterceptor])),
        provideRouter(appRoutes, withComponentInputBinding()),
        provideStore(
            [AuthState],
            // withNgxsLoggerPlugin(),
            withNgxsReduxDevtoolsPlugin()
        ),
    ],
};
