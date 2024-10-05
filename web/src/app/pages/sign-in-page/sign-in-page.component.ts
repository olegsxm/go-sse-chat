import {ChangeDetectionStrategy, Component, inject} from '@angular/core';
import {Router, RouterLink} from "@angular/router";
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthService} from "../../core/services/auth.service";
import {Store} from "@ngxs/store";
import {AuthAction} from "../../state/auth/auth.actions";

@Component({
    selector: 'app-sign-in-page',
    standalone: true,
    imports: [
        RouterLink,
        ReactiveFormsModule
    ],
    templateUrl: './sign-in-page.component.html',
    styleUrl: './sign-in-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignInPageComponent {
    authService: AuthService = inject(AuthService);
    store = inject(Store)
    router = inject(Router)

    form = new FormGroup({
        login: new FormControl('', [Validators.required]),
        password: new FormControl('', [Validators.required]),
    })

    signIn() {
        if (this.form.invalid) {
            return
        }

        this.authService.signIn(this.form.value.login as string, this.form.value.password as string)
            .subscribe(res => {
                this.store.dispatch(new AuthAction(res));
                this.router.navigateByUrl('/')
            })
    }
}
