import { ChangeDetectionStrategy, Component } from '@angular/core';
import { TuiButton, TuiLink, TuiTextfield } from '@taiga-ui/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../core/services/auth/auth.service';

@Component({
  selector: 'app-sign-in-page',
  standalone: true,
  imports: [
    TuiTextfield,
    ReactiveFormsModule,
    TuiButton,
    TuiLink
  ],
  templateUrl: './sign-in-page.component.html',
  styleUrl: './sign-in-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignInPageComponent {
  form: FormGroup = new FormGroup({
    login: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });

  constructor(private router: Router, private authService: AuthService) {
  }

  toSignUpPage(e: MouseEvent) {
    e.preventDefault();

    this.router.navigateByUrl('auth/sign-up')
      .catch(err => console.log(err));
  }

  signIn() {
    this.authService.signIn(this.form.value)
      .subscribe(res => {
        this.authService.token = res.token;
        this.router.navigateByUrl('/');
      });
  }
}
