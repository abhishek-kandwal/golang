import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WeatherdataService {

  apiUrl = "http://localhost:8080/weather";
  constructor(private http : HttpClient) { }

  getdata(action: string, start: number , stop: number):Observable<any>{
    if (action === "next"){
      // + "?action='next'"
      return this.http.get(this.apiUrl + `?action="next"&startIndex=${start}&stopIndex=${stop}`);
    }else if (action === "previous"){
      // + "?action='previous'"
      return this.http.get(this.apiUrl + `?action="previous"&startIndex=${start}&stopIndex=${stop}`);
    }   
  }
}
