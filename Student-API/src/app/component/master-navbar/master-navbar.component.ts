import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-master-navbar',
  templateUrl: './master-navbar.component.html',
  styleUrls: ['./master-navbar.component.css']
})
export class MasterNavbarComponent implements OnInit {

  @Input() loggedInValue;
  @Output() hideRegister: EventEmitter<boolean> = new EventEmitter();

  // isUserLoggedIn: boolean = false;

  constructor(private cookieService: CookieService) { }

  ngOnInit(): void {
  }

  loginToggle() {

    console.log(this.loggedInValue);
    

    if (this.loggedInValue == "Logout") {
      console.log("inside logout");
      this.cookieService.delete("Token")
      // this.isUserLoggedIn = false
    }
  }

}
