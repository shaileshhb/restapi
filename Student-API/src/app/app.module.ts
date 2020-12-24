import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from "@angular/router";
import { HttpClientModule } from "@angular/common/http";
import { FormsModule } from "@angular/forms";
import { ReactiveFormsModule } from "@angular/forms";

import { routeArgs } from "./routeConfig";
import { AppComponent } from './app.component';
import { StudentDetailsComponent } from './component/student-details/student-details.component';
import { PageNotFoundComponent } from './component/page-not-found/page-not-found.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { EmtpyToNullDirective } from './directives/emtpy-to-null.directive';
import { LoginComponent } from './component/login/login.component';
import { MasterNavbarComponent } from './component/master-navbar/master-navbar.component';
import { StudentCrudComponent } from './component/student-crud/student-crud.component';

@NgModule({
  declarations: [
    AppComponent,
    StudentDetailsComponent,
    StudentCrudComponent,
    PageNotFoundComponent,
    EmtpyToNullDirective,
    LoginComponent,
    MasterNavbarComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routeArgs),
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    NgbModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
