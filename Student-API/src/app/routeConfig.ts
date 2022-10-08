import { Routes } from "@angular/router";

import { LoginComponent } from "./component/login/login.component";
import { StudentCrudComponent } from "./component/student-crud/student-crud.component";
import { StudentDetailsComponent } from "./component/student-details/student-details.component";

export const routeArgs: Routes = [
    { path: 'home', component: StudentDetailsComponent },
    { path: 'students', component: StudentCrudComponent },
    { path: 'login', component: LoginComponent },
    { path: 'login/:id', component: LoginComponent },
    { path: '', redirectTo: '/home', pathMatch: 'full' },
]