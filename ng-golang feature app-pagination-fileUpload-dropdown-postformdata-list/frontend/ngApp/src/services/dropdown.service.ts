import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DropdownService {

  apiUrl = "http://localhost:8080/location";
  constructor(private http : HttpClient) { }

  getdropvalue():Observable<any>{
    return this.http.get(this.apiUrl);
  }

  hitdropvalue(stateValue:number):Observable<any>{
    return this.http.get(this.apiUrl + `?state=${stateValue}`);
  }
}
