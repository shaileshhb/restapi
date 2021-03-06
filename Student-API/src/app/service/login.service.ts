import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ITokenResponses } from '../ITokenResponse';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  baseURL = "http://localhost:8080/students/"

  constructor(private http: HttpClient) { }

  userLogin(userDetails: any): Observable<ITokenResponses> {
    
    let userJSON = JSON.stringify(userDetails)
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    
    console.log(userJSON)

    return this.http.post<ITokenResponses>(this.baseURL + "login", userJSON, {'headers': httpHeaders, responseType: 'text' as 'json'})
    
  }

  register(userDetails: any): Observable<ITokenResponses>{

    let userJSON = JSON.stringify(userDetails)
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8'} );
    
    console.log(userJSON)

    return this.http.post<ITokenResponses>(this.baseURL + "register", userJSON, {'headers': httpHeaders, responseType: 'text' as 'json'})
  }
}
