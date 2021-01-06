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
  // bookIssue: IBookIssue
  books = [];
  bookIssues: IBookIssue[] = [];
  id: string;
  studentForm: FormGroup;
  bookIssueForm: FormGroup;
  studentAPI: IStudentDTO;
  formTitle: string;
  userLoggedIn: boolean = false;
  isViewClicked: boolean = false;
  isBookIssued: boolean = false;
  modalRef: any;
  
  constructor(
    private studentService:StudentDTOService,
    private bookService: BooksService,
    private bookIssueService: BookIssueService,
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal,
    private cookieService: CookieService
    ) { 
      this.studentFormBuild();
  }

  studentFormBuild(){
    this.studentForm = this.formBuilder.group({
      rollNo: ['', [Validators.min(1)]],
      name: ['', [Validators.required, Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: ['', [Validators.min(5)]],
      phone: ['', [Validators.minLength(10), Validators.pattern("^[0-9]*$")]],
      date: [],
      dateTime: [],
      gender: [],
      email: ['', [Validators.required, Validators.email, 
        Validators.pattern("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]]
    });
  }

  bookIssueFormBuilder() {
    this.bookIssueForm = this.formBuilder.group({
      bookID: [{value: '', disabled: true}],
      studentID: [{value: '', disabled: true}],
      issueDate: ['', [Validators.required]]
    })
  } 
  
  ngOnInit(): void {
    if (this.cookieService.get("Token") != "") {
      this.userLoggedIn = true
    } else {
      this.userLoggedIn = false
    }
    this.getStudents();
   }


  // userRegister(registerValue) {
  //   console.log(registerValue);
  // }

  validate():void{
  
    if(this.studentForm.valid){
      if(this.formTitle == "Add"){
        this.addStudent();
      }
      else{
        this.updateStudent();
      }
    }
  }

  setAddAction():void{
    this.studentFormBuild();
    this.formTitle = "Add"
    this.isViewClicked = false
  }

  setUpdateAction(id: string): void {
    this.studentFormBuild()
    this.formTitle = "Update"
    this.isViewClicked = false
    this.prepopulate(id)
    console.log(id);
    
  }

  setViewAction(id: string): void {
    this.studentFormBuild()
    this.isViewClicked = true
    
    // this.prepopulate(id)
    this.loadBookIssueData(id)
  }

  prepopulate(id:string):void {
      
    this.studentService.getStudentDetail(id).subscribe((response) => {
      
      console.log(response);
      
      this.studentAPI = {
        id: id,
        name: response.name,
        rollNo: response.rollNo,
        age: response.age,
        phone: response.phone,
        date: response.date,
        // dateTime: response.dateTime,
        email: response.email,
        isMale: response.isMale 
      }
      
      this.studentForm.patchValue({
        name: response.name,
        rollNo: response.rollNo,
        age: response.age,
        phone: response.phone,
        date: response.date,
        // dateTime: response.dateTime,
        email: response.email,
        gender: response.isMale
      });
      
    },
    (err) => console.log('HTTP Error', err.error)
    );
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
      isMale: this.studentForm.get('gender').value,
    };
    console.log(this.studentAPI);
    

    this.studentService.addNewStudent(this.studentAPI).subscribe(data=>{
      console.log(data);
      
      this.getStudents();
      alert("Student added");
      this.modalRef.close();
    },
    (err) => {
      console.log('HTTP Error', err.error)
      alert("Error: " + err.error)
      
      if (err.status == 401) {
        this.router.navigateByUrl('/login')
        this.modalRef.close() 
      }

    }
    );
  }

  updateStudent():void{

    console.log(this.studentAPI.id);

    this.studentService.updateExisitingStudent(this.studentAPI.id, {
      "id": this.studentAPI.id,
      "rollNo": this.studentForm.get('rollNo').value, 
      "name": this.studentForm.get('name').value, 
      "age": this.studentForm.get('age').value, 
      "email": this.studentForm.get('email').value,
      "phone": this.studentForm.get('phone').value,
      "date": this.studentForm.get('date').value,
      // "dateTime": this.studentForm.get('dateTime').value,
      "isMale": this.studentForm.get('gender').value
    }).
    subscribe((data) => {        
      this.getStudents();
      alert("Student updated");
      this.modalRef.close();
    },
    (err) => {
      console.log('HTTP Error', err.error)
      alert("Error: " + err.error)

      if (err.status == 401) {
        this.router.navigateByUrl('/login')
        this.modalRef.close() 
      }
      
    }
    );
  }

  deleteStudent(id:string):void{
    if(confirm("Are you sure to delete?")) {
      this.studentService.deleteStudent(id).subscribe((data)=>{

        this.getStudents();
        alert("Student deleted");
        this.modalRef.close();

      },
      (err) => {
        console.log('HTTP Error', err.error)
        alert("Error: " + err.error)
          
        if (err.status == 401) {
          this.router.navigateByUrl('/login')
          this.modalRef.close() 
        }
        
      }
      );
    }
  }

  // book issue
  loadBookIssueData(id: string): void {

    this.bookIssueService.getBookIssues(id).subscribe((response) => {
      this.bookIssues = response
    },
    err => {
      alert("ERROR")
      console.log(err.error);
      
    })


    // this.bookIssues = []
    // this.bookIssueService.getBookIssues().subscribe((response) => {

    //   for(var i=0; i < response.length; i++) {
    //     if (id == response[i].studentID) {
    //       console.log(response[i]);
    //       this.bookIssues.push(response[i])
    //     }
    //   }

    // })
  }

  returnBookIssued(bookID: string, studentID: string) {

    if (confirm("Are you sure?")) {
      this.bookIssueService.updateBookIssue(bookID, {
        "studentID": studentID,
        "returnedFlag": true,
        "penalty": 0.0
      }).subscribe(response => {
        console.log(response);
        
        alert("Book Successfully returned")
        this.loadBookIssueData(studentID)
        // this.modalService.dismissAll()
      },
      err => {
        console.log('HTTP Error', err.error)
        // alert("Error: " + err.error)
      })
    }
  }

  // books
  showInventory(studentID: string, bookIssues) {

    console.log(bookIssues);
    
    this.bookService.getBooks().subscribe((response) => {
      localStorage.setItem('studentID', studentID)
      this.books = response
      console.log(response);
    },
    err => {
      console.log("ERROR: ", err.error);
      // alert("ERROR: ", err.error)
    })
  }

  prepopulateBook(id: string) {

    this.bookIssueFormBuilder()
    this.bookIssueForm.patchValue({
      bookID: id,
      studentID: localStorage.getItem('studentID')
    })
  }

  issueBook(bookID: string) {

    this.bookIssueService.addNewBookIssue({
      "bookID": bookID,
      "studentID": localStorage.getItem('studentID'),
      // "issueDate": this.bookIssueForm.get('issueDate').value,
      "returnedFlag": false,
      "penalty": 0.0
    }).subscribe(response => {
      alert("Book successfully issued")
      this.modalRef.close()
      this.getStudents()
    },
    err => {
      alert("Error: " + err.error)
      console.log("Error: " , err.error);
      this.modalRef.close()
      
    })
  }

  openModal(modalContent: any, modalSize?:any) {  

    let size

    if (modalSize == undefined) {
      size = 'xl'
    } else {
      size = modalSize
    }    
    this.modalRef = this.modalService.open(modalContent, {ariaLabelledBy: 'modal-basic-title', backdrop:'static', size: size})
  }

  openModalAfterAuthentication(modalContent: any, modalSize?:any) {

    if (this.cookieService.get("Token") == "") {
      alert("Session has expired. Please login")
      this.router.navigateByUrl("/login");
      return
    }
    
    this.openModal(modalContent, modalSize)
  }
  
}
