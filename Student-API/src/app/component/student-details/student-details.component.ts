import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { StudentDTOService } from '../../service/student-dto.service';

@Component({
  selector: 'app-student-details',
  templateUrl: './student-details.component.html',
  styleUrls: ['./student-details.component.css']
})
export class StudentDetailsComponent implements OnInit {

  studentDetails = [];
  login = "Login";

  constructor(private studentDto: StudentDTOService, private router: Router ) { }

  ngOnInit(): void {

    this.studentDto.getStudentDetails()
    .subscribe(res => {
      this.studentDetails = res;
      for (let i = 0; i < this.studentDetails.length; i++) {
        this.studentDetails[i].isMale =  this.studentDetails[i].isMale == true? "Male": "Female";
      }
      console.log(this.studentDetails);
    }, 
    e => {
      console.log(e);
      
    })
  }

  deleteStudent = function(studentID) {

    if(confirm("Are you sure you want to delete student?")) {
      console.log(studentID);
      this.studentDto.deleteStuden(studentID)
      .subscribe(res=> {
        this.ngOnInit();
      })
    }
    
  }

  saveStudentID = function(studentID) {
    
    if(confirm('Are you sure you want update this student details?')) {
      localStorage.setItem('studentID', studentID);
      this.router.navigateByUrl('/updateStudent');
    }
  }

}
