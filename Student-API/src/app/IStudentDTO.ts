import { IBookIssue } from "./IBookIssue";

export interface IStudentDTO {
    id: string;
    name: string;
    rollNo: number;
    age: number;
    email: string;
    phone: string;
    date: string;
    // dateTime: string;
    isMale: boolean;
    bookIssues: IBookIssue[];
}