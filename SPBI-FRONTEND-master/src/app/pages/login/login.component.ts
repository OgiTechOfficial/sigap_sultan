import { Component, OnInit } from '@angular/core';
import { UntypedFormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required]],
  });

  constructor(
    private fb: UntypedFormBuilder,
    private authService: AuthService,
    private router: Router
  ) {}

  ngOnInit(): void {}

  login() {
    try {
      if (this.loginForm.valid) {
        const formValue = this.loginForm.value;
        this.authService.login(formValue).subscribe(_ => {
          this.router.navigate(['/main/dashboard']);
        });
      }
    } catch (error) {
      console.error(error);
    }
  }

  get ctrl() {
    return this.loginForm.controls;
  }
}
