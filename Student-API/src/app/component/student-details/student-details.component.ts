import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

import { StudentDTOService } from '../../service/student-dto.service';

@Component({
  selector: 'app-student-details',
  templateUrl: './student-details.component.html',
  styleUrls: ['./student-details.component.css']
})
export class StudentDetailsComponent implements OnInit {

  studentDetails = [];
  login = "Login";

  constructor(private cookieService: CookieService) { }

  ngOnInit(): void {

    console.log(this.cookieService.get("Token"));
    if (this.cookieService.get("Token") == "") {
      this.login = "Login"
    } else {
      this.login = "Logout"
    }

  }


}
