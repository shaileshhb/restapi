import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-master-navbar',
  templateUrl: './master-navbar.component.html',
  styleUrls: ['./master-navbar.component.css']
})
export class MasterNavbarComponent implements OnInit {

  @Input() loggedInValue;
  @Output() hideRegister: EventEmitter<boolean> = new EventEmitter();

  isUserLoggedIn: boolean;

  constructor() { }

  ngOnInit(): void {

    if (this.loggedInValue == "Login") {
      this.isUserLoggedIn = false
      console.log(this.isUserLoggedIn);
    } else {
      this.isUserLoggedIn = true
      console.log(this.isUserLoggedIn);
      
    }
  }

}
