import {IAuthResponse} from "../../core/models/responce.models";

const scope = '[AUTH]'

export class AuthAction {
    static readonly type = `${scope} sign-in`;

    constructor(readonly payload: IAuthResponse) {
    }
}
