import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { DEFAULT_INTERRUPTSOURCES, Idle } from '@ng-idle/core';
import { Keepalive } from '@ng-idle/keepalive';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-master-navbar',
  templateUrl: './master-navbar.component.html',
  styleUrls: ['./master-navbar.component.css']
})
export class MasterNavbarComponent implements OnInit {

  loggedInValue: string = "Login";

  constructor(
    private cookieService: CookieService,
    private idle: Idle,
    private keepalive: Keepalive,
    private router: Router
    ) { }

  ngOnInit(): void {
    if (this.cookieService.get("Token") == "") {
      this.loggedInValue = "Login"
    } else {
      this.loggedInValue = "Logout"
    }
  }

  loginToggle() {

    console.log(this.loggedInValue);
    
    if (this.cookieService.get("Token") != "") {
      this.cookieService.delete("Token")
      this.loggedInValue = "Login"
    } else {
      this.loggedInValue = "Logout"
    }
  }

  setUserIdleState() {

    this.idle.setIdle(5)
    this.idle.setTimeout(5)
    this.idle.setInterrupts(DEFAULT_INTERRUPTSOURCES)

    // this.idle.onIdleEnd.subscribe(() => this.idleState = 'No longer idle.');
    this.idle.onTimeout.subscribe(() => {
      // this.idleState = 'Timed out!';
      // this.timedOut = true;
      alert("Session Timeout. Please login")
      this.cookieService.delete("Token")
      this.router.navigateByUrl('/login')
    })
    // this.idle.onIdleStart.subscribe(() => this.idleState = 'You\'ve gone idle!');
    // this.idle.onTimeoutWarning.subscribe((countdown) => this.idleState = 'You will time out in ' + countdown + ' seconds!');
    this.keepalive.interval(15)
    // this.keepalive.onPing.subscribe(() => this.lastPing = new Date());

    this.reset()
  }
  reset() {
    this.idle.watch();
    // this.idleState = 'Started.';
    // this.timedOut = false;
  }
}
