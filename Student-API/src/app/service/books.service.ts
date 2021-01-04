import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { IBooks } from '../IBooks';

@Injectable({
  providedIn: 'root'
})
export class BooksService {

  baseURL = "http://localhost:8080/books"

  constructor(
    private http: HttpClient,
    private cookieService: CookieService,
  ) { }


  getBooks(): Observable<IBooks[]> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );
    return this.http.get<IBooks[]>(this.baseURL)
  }

  
  addNewBookIssue(bookDetails): Observable<any> {

    let bookJSON: string = JSON.stringify(bookDetails);
    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    console.log(bookJSON);

    return this.http.post<any>(this.baseURL, bookJSON, {'headers': httpHeaders, responseType:'text' as 'json'} )
  }

  updateBookIssue(id: string, bookDetails: any): Observable<string> {

    let httpHeaders = new HttpHeaders( { 'Content-type': 'application/json; charset=utf-8', 'Token': this.cookieService.get("Token") } );
    let bookJSON: string = JSON.stringify(bookDetails); 

    console.log(bookJSON);
    
    
    return this.http.put<string>(this.baseURL + "/" + id, bookJSON, {'headers': httpHeaders, responseType:'text' as 'json'} );

  }

  deleteStudent(id: string): Observable<string> {
    console.log(id);
    
    let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );

    return this.http.delete<string>(this.baseURL + "/" +id, {'headers': httpHeaders, responseType:'text' as 'json'});
  }

}
