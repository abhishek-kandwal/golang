import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, from } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class FileuploadService {

  constructor(private http: HttpClient) { }
  
  private apiUrl = "http://localhost:8080/";

  getEmployeeList():Observable<any>{
    return this.http.get(this.apiUrl + "employee");
  }
  
  postFile(fileToUpload: File, id: number): Observable<any> {
    

    const formData: FormData = new FormData();
    
    formData.append('file', fileToUpload, fileToUpload.name);
    
    return this.http
      .post(this.apiUrl + "fileUpload" + `?id=${id}`, formData);
  }

}

