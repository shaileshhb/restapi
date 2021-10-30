import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { IBookIssue } from '../IBookIssue';

@Injectable({
  providedIn: 'root'
})
export class BookIssueService {

  baseURL = "http://localhost:8080/api/bookIssues"
  // baseURL = "/api/bookIssues"

  constructor(
    private http: HttpClient,
    private cookieService: CookieService
  ) { }

  getBookIssues(studentID: string): Observable<IBookIssue[]> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );
    return this.http.get<IBookIssue[]>(this.baseURL + "/" + studentID)
  }

  addNewBookIssue(bookIssueDetails): Observable<any> {

    let bookIssueJSON: string = JSON.stringify(bookIssueDetails);
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    console.log(bookIssueJSON);

    return this.http.post<any>(this.baseURL, bookIssueJSON, {'headers': httpHeaders, responseType:'text' as 'json'} )
  }

  updateBookIssue(bookID: string, bookIssueDetails: any): Observable<any> {

    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    let bookIssueJSON: string = JSON.stringify(bookIssueDetails); 

    console.log(bookID, bookIssueJSON);
  
    // return this.http.put<any>(this.baseURL + "/" + bookID, bookIssueJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );
    return this.http.put<any>(`${this.baseURL}/${bookID}`, bookIssueJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  }

  deleteStudent(bookID: string): Observable<string> {
    console.log(bookID);
    
    let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );

    return this.http.delete<string>(`${this.baseURL}/${bookID}`, {'headers': httpHeaders, responseType:'text' as 'json'});
  }
}
