import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';

import { IStudentDTO } from "src/app/IStudentDTO"; 
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class StudentDTOService {

  // url = "http://gsmktg.azurewebsites.net/api/v1/techlabs/test/students/";
  baseURL = "http://localhost:8080/students"

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  getStudentDetails(): Observable<IStudentDTO[]> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );

    return this.http.get<IStudentDTO[]>(this.baseURL);

  }

  getStudentDetail(studentID: string): Observable<IStudentDTO> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );

    return this.http.get<IStudentDTO>(this.baseURL + "/" + studentID);

  }

  addNewStudent(studentDetails): Observable<any> {

    let studentJSON: string = JSON.stringify(studentDetails);
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    
    console.log(studentJSON);

    return this.http.post<any>(this.baseURL, studentJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  }

  updateExisitingStudent(id: string, studentDetails: any): Observable<string> {

    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    let studentJSON: string = JSON.stringify(studentDetails); 

    console.log(studentJSON);
    
    
    return this.http.put<string>(this.baseURL + "/" + id, studentJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  } 

  deleteStudent(studentID: string): Observable<string> {
    console.log(studentID);
    
    let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );

    return this.http.delete<string>(this.baseURL + "/" +studentID, {'headers': httpHeaders, responseType:'text' as 'json'});
  }

}