import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-master-navbar',
  templateUrl: './master-navbar.component.html',
  styleUrls: ['./master-navbar.component.css']
})
export class MasterNavbarComponent implements OnInit {

  @Input() loggedInValue;
  // @Output() hideRegister: EventEmitter<boolean> = new EventEmitter();

  isUserLoggedIn: boolean;

  constructor(private cookieService: CookieService) { }

  ngOnInit(): void {

    console.log(this.cookieService.get("Token"));

    if (this.cookieService.get("Token") == "") {
      this.isUserLoggedIn = false
    } else {
      this.isUserLoggedIn = true
    }

  }

  loginToggle() {

    console.log(this.loggedInValue, this.isUserLoggedIn);
    
    if (this.loggedInValue == "Logout") {
      console.log("inside logout");
      this.cookieService.delete("Token")
      this.isUserLoggedIn = false
    }
  }

}
