import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { IBookIssue } from '../IBookIssue';

@Injectable({
  providedIn: 'root'
})
export class BookIssueService {

  baseURL: string

  constructor(
    private http: HttpClient,
  ) {
    // this.baseURL = "/api/bookIssues"
    this.baseURL = `${environment.BASEURL}/bookIssues`
   }

  getBookIssues(studentID: string): Observable<IBookIssue[]> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );

    return this.http.get<IBookIssue[]>(this.baseURL + "/" + studentID, { headers: httpHeaders })
  }

  addNewBookIssue(bookIssueDetails): Observable<any> {

    let bookIssueJSON: string = JSON.stringify(bookIssueDetails);
    
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    console.log(bookIssueJSON);

    return this.http.post<any>(this.baseURL, bookIssueJSON, 
      { headers: httpHeaders, responseType:'text' as 'json' } )
  }

  updateBookIssue(bookID: string, bookIssueDetails: any): Observable<any> {

    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    let bookIssueJSON: string = JSON.stringify(bookIssueDetails); 

    console.log(bookID, bookIssueJSON);
  
    // return this.http.put<any>(this.baseURL + "/" + bookID, bookIssueJSON, { headers: httpHeaders, responseType:'text' as 'json'} );
    return this.http.put<any>(`${this.baseURL}/${bookID}`, bookIssueJSON, { headers: httpHeaders, responseType:'text' as 'json'} );

  }

  deleteStudent(bookID: string): Observable<string> {
    console.log(bookID);
    
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    return this.http.delete<string>(`${this.baseURL}/${bookID}`, { headers: httpHeaders, responseType:'text' as 'json'});
  }
}
