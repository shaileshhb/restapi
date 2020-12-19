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
      rollNo: ['', Validators.required],
      name: ['', [Validators.required,  Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: ['', Validators.required],
      date: ['', Validators.required],
      gender: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]]
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

  addStudent():void{
    this.studentAPI = {
      id:null, 
      rollNo:this.studentForm.get('rollNo').value, 
      name:this.studentForm.get('name').value, 
      age:this.studentForm.get('age').value, 
      email:this.studentForm.get('email').value, 
      isMale:this.studentForm.get('gender').value, 
      date:this.studentForm.get('date').value
    };
    this.studentService.addNewStudent(this.studentAPI).subscribe(data=>{
      this.getStudents();
      alert("Student added");
      this.modalService.dismissAll();

    },
    (err) => console.log('HTTP Error', err)
    );
  }

  dobChange():void{
      let dobDate:Date = new Date(this.studentForm.controls['date'].value);
      let diff = (new Date().getTime() - dobDate.getTime());
      let ageTotal = Math.trunc(diff/ (1000 * 3600 * 24 *365));
      this.studentForm.patchValue({
        age: ageTotal,
      });
  }

    setAddAction():void{
      
      this.formBuild();
      this.addOrUpdateAction = "add";
    }

    prepopulate(id:string):void{
      this.formBuild();
      this.addOrUpdateAction = "update";
      this.id = id;
      this.studentService.getStudentDetails(id).subscribe((data)=>{
        
        console.log(data[0]);
        
        this.studentForm.patchValue({
          name: data[0].name,
          rollNo: data[0].rollNo,
          age: data[0].age,
          date: data[0].date,
          email: data[0].email,
          gender: data[0].isMale
        });
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
