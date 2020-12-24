import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Validators, FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { User } from 'src/app/user';
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

  constructor(
    private formBuilder: FormBuilder, 
    private loginSerive: LoginService,
    private router: Router,
    private modalService: NgbModal
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

  userRegister(registerValue) {
    console.log(registerValue);
    
  }

  buildRegisterForm() {
    this.registerForm = this.formBuilder.group({
      email:['', [Validators.required, Validators.pattern("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]],
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
  }

  validateUser() {
    console.log(this.loginForm.value);

    this.loginSerive.userLogin(this.loginForm.value).subscribe(response => {

      console.log(response);

      if (response == "Success") {
        this.login = 'Logout';
        this.router.navigateByUrl('/students')
      } else {
        alert('Invalid username or password')
      }
      
    },
    (err) => {
      alert(err)
    })
    
  }

  openModal(modalContent: any) {
    this.modalService.open(modalContent, {ariaLabelledBy: "modal-basic-title", backdrop: "static", size: "xl"})
  }

  registerUser() {
    console.log(this.registerForm.value);

    this.loginSerive.register(this.registerForm.value).subscribe(response => {
      console.log(response);
      alert("Sccuessful")
      this.modalService.dismissAll()
      this.router.navigateByUrl('/students')
    },
    err => {
      alert(err)
    })
    
  }

}
