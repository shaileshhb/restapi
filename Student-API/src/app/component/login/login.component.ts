import { Component, OnInit  } from '@angular/core';
import { Validators, FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { DEFAULT_INTERRUPTSOURCES, Idle } from '@ng-idle/core';
import { Keepalive } from '@ng-idle/keepalive';
import { CookieService } from 'ngx-cookie-service';
import { ITokenResponses } from 'src/app/ITokenResponse';
import { LoginService } from '../../service/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  registerForm: FormGroup;
  login = "Login";
  tokens: ITokenResponses;

  constructor(
    private formBuilder: FormBuilder, 
    private loginSerive: LoginService,
    private router: Router,
    private modalService: NgbModal,
    private cookieService: CookieService
    ) { 
    this.buildLoginForm()
    this.buildRegisterForm()

  }

  ngOnInit(): void {
  }

  buildLoginForm() {
    this.loginForm = this.formBuilder.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
  }

  // userRegister(registerValue) {
  //   registerValue = true
  //   console.log(registerValue);
  // }

  buildRegisterForm() {
    this.registerForm = this.formBuilder.group({
      email:['', [Validators.required, Validators.pattern("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]],
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
  }

  validateUser() {

    this.loginSerive.userLogin(this.loginForm.value).subscribe(response => {
      console.log(response)
      this.setLoginCookie(response)
      console.log(this.cookieService.get("Token"));
      this.login = 'Logout';
      this.router.navigate(['/students'])

    },
    (err) => {
      alert("Error:" + err.error)
      console.log("Error:" + err.error);
      
    })

  }

  openModal(modalContent: any) {
    this.modalService.open(modalContent, {ariaLabelledBy: "modal-basic-title", backdrop: "static", size: "xl"})
  }

  registerUser() {
    console.log(this.registerForm.value);

    this.loginSerive.register(this.registerForm.value).subscribe(response => {
      console.log(response);
      alert("User Registered Successfully")

      this.setLoginCookie(response)

      this.login = 'Logout';
      this.modalService.dismissAll()
      this.router.navigate(['/students'])
    },
    err => {
      console.log(err.error);
      
      alert("Error:" + err.error)
      // alert(err)
    })
    
  }

  setLoginCookie(value: any) {

    const expiryTime = new Date();
    expiryTime.setMinutes(expiryTime.getMinutes() + 5);

    console.log(expiryTime);
    this.cookieService.set("Token", value, expiryTime)
  }

}
