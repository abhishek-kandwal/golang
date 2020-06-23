import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class UserdetailsService {

  apiUrl = 'http://localhost:8080/'
  constructor( private http : HttpClient) { }

  getUserDetails():Observable<any>{
    return this.http.get(this.apiUrl + "guserdetails");
  }

  getvalidateEmail(email:any):Observable<any>{
    return this.http.get(this.apiUrl + "guserdetails" + `?email="${email}"`);
  }

  postUserDetails(data : any):Observable<any>{
    console.log(data);
    var body = {
      name : data.name,
      email : data.email,
      dob: data.dob
    }
    console.log(body);
    return this.http.post(this.apiUrl + "puserdetails", body);
  }
}
