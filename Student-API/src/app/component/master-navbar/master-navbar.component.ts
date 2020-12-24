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

  isUserLoggedIn: boolean;

  constructor(private cookieService: CookieService) { }

  ngOnInit(): void {

    if (this.loggedInValue == "Login") {
      this.isUserLoggedIn = false
      this.cookieService.delete("Token")
    } else {
      this.isUserLoggedIn = true

    }
  }

}
