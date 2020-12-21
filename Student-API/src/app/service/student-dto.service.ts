import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';

import { IStudentDTO } from "src/app/IStudentDTO"; 

@Injectable({
  providedIn: 'root'
})
export class StudentDTOService {

  // url = "http://gsmktg.azurewebsites.net/api/v1/techlabs/test/students/";
  url = "http://localhost:8080/students"

  constructor(private http: HttpClient) { }

  getStudentDetails(studentID?: string): Observable<IStudentDTO[]> {

    if(studentID == undefined) {
      return this.http.get<IStudentDTO[]>(this.url);
    } else {
      return this.http.get<IStudentDTO[]>(this.url + "/" + studentID);
    }

  }

  addNewStudent(studentDetails): Observable<any> {

    let studentJSON: string = JSON.stringify(studentDetails);
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    
    console.log(studentJSON);

    return this.http.post<any>(this.url, studentJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  }

  updateExisitingStudent(id: string, studentDetails: any): Observable<IStudentDTO> {

    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    let studentJSON: string = JSON.stringify(studentDetails); 

    console.log(studentJSON);
    
    
    return this.http.put<IStudentDTO>(this.url + "/" + id, studentJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  } 

  deleteStudent(studentID: string): Observable<IStudentDTO> {
    console.log(studentID);
    
    return this.http.delete<IStudentDTO>(this.url + "/" +studentID, {responseType:'text' as 'json'});
  }

}
