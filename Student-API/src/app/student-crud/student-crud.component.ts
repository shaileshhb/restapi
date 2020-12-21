import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { IStudentDTO } from 'src/app/IStudentDTO';
import { StudentDTOService } from 'src/app/service/student-dto.service';

@Component({
  selector: 'app-student-crud',
  templateUrl: './student-crud.component.html',
  styleUrls: ['./student-crud.component.css']
})
export class StudentCrudComponent implements OnInit {

  students:IStudentDTO[] = [];
  id: string;
  studentForm: FormGroup;
  studentAPI: IStudentDTO;
  addOrUpdateAction: string;
  
  constructor(
    private studentService:StudentDTOService, 
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal
    ) { 
      this.formBuild();
  }

  formBuild(){
    this.studentForm = this.formBuilder.group({
      rollNo: [],
      name: ['', [Validators.required, Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: [],
      date: [],
      dateTime: [],
      gender: [],
      email: ['', [Validators.required, Validators.email, 
        Validators.pattern("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]]
    });
  }
  
  ngOnInit(): void {
    this.getStudents();
   }

  getStudents():void{
    this.studentService.getStudentDetails().subscribe((data)=>{
      this.students = data;
    },
    (err) => console.log('HTTP Error', err)
    );
  }

  validate():void{
  
    if(this.studentForm.valid){
      if(this.addOrUpdateAction == "add"){
        this.addStudent();
      }
      else{
        this.updateStudent();
      }
    }
  }

  setAddAction():void{
    this.formBuild();
    this.addOrUpdateAction = "add";
  }

    prepopulate(id:string):void{
      this.formBuild();
      this.addOrUpdateAction = "update";
      this.id = id;
      this.studentService.getStudentDetails(id).subscribe((response)=>{

        this.studentForm.patchValue({
          name: response[0].name,
          rollNo: response[0].rollNo,
          age: response[0].age,
          date: response[0].date,
          dateTime: response[0].dateTime,
          email: response[0].email,
          gender: response[0].isMale
        });
      },
      (err) => console.log('HTTP Error', err)
      );
    }

    addStudent():void{
      
      let dateWithTime = this.studentForm.get('dateTime').value
      let date = dateWithTime.slice(0, 10)
      let time = dateWithTime.slice(11)
      console.log(date, time);
      
      
      this.studentAPI = {
        id: null, 
        rollNo: this.studentForm.get('rollNo').value, 
        name: this.studentForm.get('name').value, 
        age: this.studentForm.get('age').value, 
        email: this.studentForm.get('email').value, 
        date: this.studentForm.get('date').value,
        dateTime: date.concat(" ", time),
        isMale: this.studentForm.get('gender').value

      };
      this.studentService.addNewStudent(this.studentAPI).subscribe(data=>{
        this.getStudents();
        alert("Student added");
        this.modalService.dismissAll();
      },
      (err) => console.log('HTTP Error', err)
      );
    }

    updateStudent():void{

      this.studentService.updateExisitingStudent(this.id, {
        // "id":this.id,
        "rollNo": this.studentForm.get('rollNo').value, 
        "name": this.studentForm.get('name').value, 
        "age": this.studentForm.get('age').value, 
        "email": this.studentForm.get('email').value, 
        "date": this.studentForm.get('date').value,
        "dateTime": this.studentForm.get('dateTime').value,
        "isMale": this.studentForm.get('gender').value
      }).
      subscribe((data)=>{
        this.getStudents();
        alert("Student updated");
        this.modalService.dismissAll();
      },
      (err) => console.log('HTTP Error', err)
      );
    }

    deleteStudent(id:string):void{
      if(confirm("Are you sure to delete?")) {
        this.studentService.deleteStudent(id).subscribe((data)=>{
          this.getStudents();
          alert("Student deleted");
          this.modalService.dismissAll();

        },
        (err) => console.log('HTTP Error', err)
        );
      }
    }

    openStudentModalForm(studentModel: any) {
      this.modalService.open(studentModel, {ariaLabelledBy: 'modal-basic-title', backdrop:'static', size:'xl'})

    }

}
