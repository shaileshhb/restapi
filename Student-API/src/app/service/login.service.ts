import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { ITokenResponses } from '../ITokenResponse';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  baseURL: string

  constructor(private http: HttpClient) {
    this.baseURL = `${environment.BASEURL}/students`
    // this.baseURL = "http://localhost:8080/students"
   }

  userLogin(userDetails: any): Observable<any> {
    
    let userJSON = JSON.stringify(userDetails)
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    
    // return this.http.post<ITokenResponses>(this.baseURL + "login", userJSON, {'headers': httpHeaders, responseType: 'text' as 'json'})
    return this.http.post<any>(`${this.baseURL}/login`, userJSON, 
      { headers: httpHeaders, responseType: 'text' as 'json'})
    
  }

  register(userDetails: any): Observable<ITokenResponses>{

    let userJSON = JSON.stringify(userDetails)
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    
    console.log(userJSON)

    // return this.http.post<ITokenResponses>(this.baseURL + "register", userJSON, {headers: httpHeaders, responseType: 'text' as 'json'})
    return this.http.post<ITokenResponses>(`${this.baseURL}/register`, userJSON, 
      { headers: httpHeaders, responseType: 'text' as 'json'})
  }
}
