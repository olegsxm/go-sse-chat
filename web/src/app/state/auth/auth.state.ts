import {Injectable} from '@angular/core';
import {Action, Selector, State, StateContext} from '@ngxs/store';
import {IUser} from "../../core/models/user.model";
import {AuthAction} from "./auth.actions";

const ls_token_key = 'ls_token_key';
const ls_user_key = 'ls_user_key';

export interface IAuthState {
    token: string | null;
    user: IUser | null;
}

@State<IAuthState>({
    name: 'auth',
    defaults: {
        token: localStorage.getItem(ls_token_key),
        user: (() => {
            const u = localStorage.getItem(ls_user_key);
            return u ? JSON.parse(u) : null;
        })()
    }
})
@Injectable()
export class AuthState {

    @Selector()
    static getState(state: IAuthState) {
        return state;
    }

    @Selector()
    static getUser(state: IAuthState) {
        return state.user;
    }

    @Action(AuthAction)
    auth(ctx: StateContext<IAuthState>, {payload}: AuthAction) {
        localStorage.setItem(ls_token_key, payload.token);
        localStorage.setItem(ls_user_key, JSON.stringify(payload.user));
        ctx.setState(payload)
    }
}
