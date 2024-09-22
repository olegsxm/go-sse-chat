import { ChangeDetectionStrategy, Component } from '@angular/core';
import { TuiButton, TuiLink, TuiTextfield } from '@taiga-ui/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';

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

  constructor(private router: Router) {
  }

  toSignUpPage(e: MouseEvent) {
    e.preventDefault();

    this.router.navigateByUrl('auth/sign-up')
      .catch(err => console.log(err));
  }
}
