import {ChangeDetectionStrategy, Component, inject} from '@angular/core';
import {Router, RouterLink} from "@angular/router";
import {AuthService} from "../../core/services/auth.service";
import {Store} from "@ngxs/store";
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthAction} from "../../state/auth/auth.actions";

@Component({
    selector: 'app-sign-up-page',
    standalone: true,
    imports: [
        RouterLink,
        ReactiveFormsModule
    ],
    templateUrl: './sign-up-page.component.html',
    styleUrl: './sign-up-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignUpPageComponent {
    authService: AuthService = inject(AuthService);
    store = inject(Store)
    router = inject(Router)

    form = new FormGroup({
        login: new FormControl('', [Validators.required]),
        password: new FormControl('', [Validators.required]),
    })

    signUp() {
        if (this.form.invalid) {
            return
        }

        this.authService.signUp(this.form.value.login as string, this.form.value.password as string)
            .subscribe(res => {
                this.store.dispatch(new AuthAction(res));
                this.router.navigateByUrl('/')
            })
    }
}
