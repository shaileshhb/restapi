import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BooksService {

  baseURL: string

  constructor(
    private http: HttpClient,
  ) { 
    // this.baseURL = "/api/books"
    this.baseURL = `${environment.BASEURL}/books`
  }


  getBooks(): Observable<any> {
    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    return this.http.get(this.baseURL, { headers: httpHeaders })
  }
  
  addNewBookIssue(bookDetails): Observable<any> {

    let bookJSON: string = JSON.stringify(bookDetails);
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    console.log(bookJSON);

    return this.http.post<any>(this.baseURL, bookJSON, { headers: httpHeaders, responseType:'text' as 'json'} )
  }

  updateBookIssue(id: string, bookDetails: any): Observable<string> {

    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );
    let bookJSON: string = JSON.stringify(bookDetails); 

    console.log(bookJSON);
    
    
    // return this.http.put<string>(this.baseURL + "/" + id, bookJSON, { headers: httpHeaders, responseType:'text' as 'json'} );
    return this.http.put<string>(`${this.baseURL}/${id}`, bookJSON, 
      { headers: httpHeaders, responseType:'text' as 'json'} );

  }

  deleteStudent(id: string): Observable<string> {
    console.log(id);
    
    let httpHeaders = new HttpHeaders( { 'Token': localStorage.getItem("token") } );

    // return this.http.delete<string>(this.baseURL + "/" +id, { headers: httpHeaders, responseType:'text' as 'json'});
    return this.http.delete<string>(`${this.baseURL}/${id}`, { headers: httpHeaders, responseType:'text' as 'json'});
  }

}
