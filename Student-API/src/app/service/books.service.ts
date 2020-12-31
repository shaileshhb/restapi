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

}
