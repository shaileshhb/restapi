import { Routes } from "@angular/router";

import { StudentCrudComponent } from "./student-crud/student-crud.component";
import { StudentDetailsComponent } from "./student-details/student-details.component";
// import { PageNotFoundComponent } from "./page-not-found/page-not-found.component";

export const routeArgs: Routes = [
    { path: 'home', component: StudentDetailsComponent },
    { path: 'students', component: StudentCrudComponent },
    { path: '', redirectTo: '/home', pathMatch: 'full' },
]