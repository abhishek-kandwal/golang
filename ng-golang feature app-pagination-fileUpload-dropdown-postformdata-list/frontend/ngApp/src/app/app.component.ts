import { Component, OnInit } from '@angular/core';
import {WeatherdataService } from '../services/weatherdata.service';

export interface PeriodicElement {
  City: string;
  Conds: string;
  Humidity: number;
  Temp: number;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit{
  
  isDropDownSelected: boolean = false;
  isPaginationSelected: boolean = false;
  isfilupload: boolean = false;
  isGetDetails:boolean = false;
  isPostDetails:boolean = false;

  start = 0;
  stop = 5;
  isnext= true;
  isValid = false;
  dataSource:PeriodicElement[];
  displayedColumns: string[] = ['city', 'conds', 'humidity', 'temp'];
  constructor(private weatherService: WeatherdataService) {}


  ngOnInit() {
    this.weatherService.getdata("next", this.start , this.stop).subscribe( result => {
      this.dataSource = result.data;
      this.start = result.start;
      this.stop = result.stop;
      if (this.stop == result.total){
        this.isnext = false
      }
    })
  }


  next(){
    this.start = this.start + 5
    this.stop = this.stop + 5
    this.weatherService.getdata("next", this.start , this.stop).subscribe( result => {
        this.dataSource = result.data;
        this.start = result.start;
        this.stop = result.stop;
        if (this.start != 0){
          this.isValid = true;
        }
        if (this.stop > result.total){
          this.isnext = false
        }else{
          this.isnext = true
        }
    })
  }

  previous(){
    this.start = this.start - 5
    this.stop = this.stop - 5
    this.weatherService.getdata("previous", this.start , this.stop).subscribe( result => {
      this.dataSource = result.data;
      this.start = result.start;
      this.stop = result.stop;
      if (this.start == 0){
        this.isValid = false;
      }
      if (this.stop > result.total){
        this.isnext = false
      }else{
        this.isnext = true
      }
    })
  }

  buttonhit(value: string){
    if (value == "upload"){
      this.isDropDownSelected = false;
      this.isPaginationSelected = false;
      this.isGetDetails= false;
      this.isfilupload = true;
      this.isPostDetails = false;
    }else if (value == "pagination"){
      this.isDropDownSelected = false;
      this.isPaginationSelected = true;
      this.isfilupload = false;
      this.isPostDetails = false;
      this.isGetDetails= false;
    }else if (value == "dropdown"){
      this.isDropDownSelected= true;
      this.isPaginationSelected = false;
      this.isfilupload = false;
      this.isPostDetails = false;
      this.isGetDetails= false;
    }else if (value == "getDetails"){
      this.isDropDownSelected= false;
      this.isPaginationSelected = false;
      this.isfilupload = false;
      this.isPostDetails = false;
      this.isGetDetails= true;
    }else if (value == "postDetails"){
      this.isDropDownSelected= false;
      this.isPaginationSelected = false;
      this.isfilupload = false;
      this.isGetDetails= false;
      this.isPostDetails = true;
    }
  }

}

