import { ChangeDetectionStrategy, Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import {
  TuiButton,
  TuiLink,
  TuiTextfieldComponent,
  TuiTextfieldDirective,
  TuiTextfieldOptionsDirective
} from '@taiga-ui/core';
import { AuthService } from '../../core/services/auth/auth.service';

@Component({
  selector: 'app-sign-up-page',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    TuiButton,
    TuiLink,
    TuiTextfieldComponent,
    TuiTextfieldDirective,
    TuiTextfieldOptionsDirective
  ],
  templateUrl: './sign-up-page.component.html',
  styleUrl: './sign-up-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignUpPageComponent {
  form: FormGroup = new FormGroup({
    login: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });

  constructor(private router: Router, private authService: AuthService) {
  }

  toSignInPage(e: MouseEvent) {
    e.preventDefault();

    this.router.navigateByUrl('auth/sign-in')
      .catch(err => console.log(err));
  }

  signUp() {
    this.authService.signUp(this.form.value)
      .subscribe(res => {
        this.authService.token = res.token;
      });
  }
}
