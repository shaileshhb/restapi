import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CookieService } from 'ngx-cookie-service';
import { IBookIssue } from 'src/app/IBookIssue';
import { IStudentDTO } from 'src/app/IStudentDTO';
import { BookIssueService } from 'src/app/service/book-issue.service';
import { BooksService } from 'src/app/service/books.service';
import { StudentDTOService } from 'src/app/service/student-dto.service';

@Component({
  selector: 'app-student-crud',
  templateUrl: './student-crud.component.html',
  styleUrls: ['./student-crud.component.css']
})
export class StudentCrudComponent implements OnInit {

  students = [];
  books = [];
  bookIssues: IBookIssue[] = [];
  id: string;
  studentForm: FormGroup;
  studentAPI: IStudentDTO;
  formTitle: string;
  userLoggedIn: boolean = false;
  isViewClicked: boolean = false;
  
  constructor(
    private studentService:StudentDTOService,
    private bookService: BooksService,
    private bookIssueService: BookIssueService,
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal,
    private cookieService: CookieService
    ) { 
      this.formBuild();
  }

  formBuild(){
    this.studentForm = this.formBuilder.group({
      rollNo: ['', [Validators.min(1)]],
      name: ['', [Validators.required, Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: ['', [Validators.min(17)]],
      phone: ['', [Validators.minLength(10), Validators.pattern("^[0-9]*$")]],
      date: [],
      dateTime: [],
      gender: [],
      email: ['', [Validators.required, Validators.email, 
        Validators.pattern("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]]
    });
  }
  
  ngOnInit(): void {
    if (this.cookieService.get("Token") != "") {
      this.userLoggedIn = true
    } else {
      this.userLoggedIn = false
    }
    this.getStudents();
   }

  getStudents():void{
    this.studentService.getStudentDetails().subscribe((data)=>{
      this.students = data;
      for (let i = 0; i < this.students.length; i++) {
        // this.students[i].dateTime = this.students[i].dateTime.replace('T', " ")
        if(this.students[i].isMale != null) {
          this.students[i].isMale = this.students[i].isMale == true ? "Male" : "Female"         
        } else {
          this.students[i].isMale = ""
        }
      }
      console.log(this.students);
      
    },
    (err) => {
      console.log('HTTP Error', err.error)
      alert("Error: " + err.error)
      if (err.status == 401) {
        this.router.navigateByUrl('/login')
        this.modalService.dismissAll() 
      }
    }
    );
  }

  // userRegister(registerValue) {
  //   console.log(registerValue);
  // }

  validate():void{
  
    if(this.studentForm.valid){
      if(this.formTitle == "add"){
        this.addStudent();
      }
      else{
        this.updateStudent();
      }
    }
  }

  setAddAction():void{
    this.formBuild();
    this.formTitle = "Add"
    this.isViewClicked = false
  }

  setUpdateAction(id: string): void {
    this.formBuild()
    this.formTitle = "Update"
    this.isViewClicked = false
    this.prepopulate(id)
  }

  setViewAction(id: string): void {
    this.formBuild()
    this.formTitle = "View"
    this.isViewClicked = true

    console.log(this.isViewClicked);
    
    this.prepopulate(id)
    this.loadBookIssueData(id)
  }

  prepopulate(id:string):void {
      
    this.studentService.getStudentDetail(id).subscribe((response) => {
      
      console.log(response);
      
      this.studentForm.patchValue({
        name: response.name,
        rollNo: response.rollNo,
        age: response.age,
        phone: response.phone,
        date: response.date,
        // dateTime: response[0].dateTime,
        email: response.email,
        gender: response.isMale
      });
    },
    (err) => console.log('HTTP Error', err.error)
    );
  }

  loadBookIssueData(id: string): void {

    this.bookIssueService.getBookIssues().subscribe((response) => {

      for(var i=0; i < response.length; i++) {
        if (id == response[i].studentID) {
          console.log(response[i]);
          this.bookIssues.push(response[i])
        }
      }

    })
  }

  addStudent():void{

    this.studentAPI = {
      id: null, 
      rollNo: (this.studentForm.get('rollNo').value  == "") ? null : this.studentForm.get('rollNo').value, 
      name: this.studentForm.get('name').value, 
      age: (this.studentForm.get('age').value == "") ? null : this.studentForm.get('age').value, 
      email: this.studentForm.get('email').value,
      phone: (this.studentForm.get('phone').value == "") ? null : this.studentForm.get('phone').value,
      date: (this.studentForm.get('date').value == "") ? null : this.studentForm.get('date').value,
      // dateTime: (this.studentForm.get('dateTime').value == "") ? null : this.studentForm.get('dateTime').value,
      isMale: this.studentForm.get('gender').value
    };
    console.log(this.studentAPI);
    

    this.studentService.addNewStudent(this.studentAPI).subscribe(data=>{
      console.log(data);
      
      this.getStudents();
      alert("Student added");
      this.modalService.dismissAll();
    },
    (err) => {
      console.log('HTTP Error', err.error)
      alert("Error: " + err.error)
      
      if (err.status == 401) {
        this.router.navigateByUrl('/login')
        this.modalService.dismissAll() 
      }

    }
    );
  }

  updateStudent():void{

    this.studentService.updateExisitingStudent(this.id, {
      "id":this.id,
      "rollNo": this.studentForm.get('rollNo').value, 
      "name": this.studentForm.get('name').value, 
      "age": this.studentForm.get('age').value, 
      "email": this.studentForm.get('email').value,
      "phone": this.studentForm.get('phone').value,
      "date": this.studentForm.get('date').value,
      // "dateTime": this.studentForm.get('dateTime').value,
      "isMale": this.studentForm.get('gender').value
    }).
    subscribe((data)=>{        
      this.getStudents();
      alert("Student updated");
      this.modalService.dismissAll();
    },
    (err) => {
      console.log('HTTP Error', err.error)
      alert("Error: " + err.error)

      if (err.status == 401) {
        this.router.navigateByUrl('/login')
        this.modalService.dismissAll() 
      }
      
    }
    );
  }

  deleteStudent(id:string):void{
    if(confirm("Are you sure to delete?")) {
      this.studentService.deleteStudent(id).subscribe((data)=>{

        this.getStudents();
        alert("Student deleted");
        this.modalService.dismissAll();

      },
      (err) => {
        console.log('HTTP Error', err.error)
        alert("Error: " + err.error)
          
        if (err.status == 401) {
          this.router.navigateByUrl('/login')
          this.modalService.dismissAll() 
        }
        
      }
      );
    }
  }

  showInventory(): void {

    this.bookService.getBooks().subscribe((response) => {
      this.books = response
      console.log(response);
      
    })

  }
  
  openStudentModalForm(studentModel: any) {
    // if (this.cookieService.get("Token") == "") {
    //   alert("Session has expired. Please login")
    //   this.router.navigateByUrl("/login");
    //   return
    // }
     this.modalService.open(studentModel, {ariaLabelledBy: 'modal-basic-title', backdrop:'static', size:'xl'})
  }
  
}
