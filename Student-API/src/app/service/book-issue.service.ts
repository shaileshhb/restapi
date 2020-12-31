import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IBookIssue } from '../IBookIssue';

@Injectable({
  providedIn: 'root'
})
export class BookIssueService {

  baseURL = "http://localhost:8080/bookIssues"

  constructor(
    private http: HttpClient
  ) { }

  getBookIssues(): Observable<IBookIssue[]> {

    // let httpHeaders = new HttpHeaders( { 'Token': this.cookieService.get("Token") } );
    return this.http.get<IBookIssue[]>(this.baseURL)
  }

}
