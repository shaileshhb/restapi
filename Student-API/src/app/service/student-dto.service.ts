import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";
import { Observable } from 'rxjs';

import { IStudentDTO } from "src/app/IStudentDTO"; 
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class StudentDTOService {

  baseURL: string

  constructor(
    private http: HttpClient, 
    ) { 
      this.baseURL = "http://localhost:8080/students"
    }

  getStudentDetails(): Observable<IStudentDTO[]> {
    return this.http.get<IStudentDTO[]>(`${this.baseURL}`);
  }

  getStudentDetail(studentID: string): Observable<IStudentDTO> {
    return this.http.get<IStudentDTO>(`${this.baseURL}/${studentID}`);
  }

  addNewStudent(studentDetails): Observable<any> {
    let studentJSON: string = JSON.stringify(studentDetails);
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    
    return this.http.post<any>(`${this.baseURL}`, studentJSON, { headers: httpHeaders, responseType:'text' as 'json'} );
  }

  updateExisitingStudent(id: string, studentDetails: any): Observable<string> {
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    let studentJSON: string = JSON.stringify(studentDetails); 

    return this.http.put<string>(`${this.baseURL}/${id}`, studentJSON, { headers: httpHeaders, responseType:'text' as 'json'} );
  } 

  deleteStudent(studentID: string): Observable<string> {
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    return this.http.delete<string>(`${this.baseURL}/${studentID}`, { headers: httpHeaders, responseType:'text' as 'json'});
  }

  searchStudent(data: any): Observable<IStudentDTO[]> {
    let params = new HttpParams()
    for (let key of Object.keys(data)) {
      
      if(data[key] == "" || data[key] == null) {
        continue
      }
      params = params.set(key, data[key])
    }

    return this.http.get<IStudentDTO[]>(`${this.baseURL}/search`, { params: params }); 
  }

}