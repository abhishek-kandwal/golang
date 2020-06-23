import { Component, OnInit } from '@angular/core';
import { DropdownService } from '../../services/dropdown.service';
 
export interface cities {
  Name: string;
  Id: number;
}

export interface states {
  Name: string;
  Id: number;
}

@Component({
  selector: 'app-dropdown',
  templateUrl: './dropdown.component.html',
  styleUrls: ['./dropdown.component.css']
})


export class DropdownComponent implements OnInit {

  constructor(private dropdownService: DropdownService) { }
  states: states[];

  cities: cities[];
  
  disableSCity = true;

  ngOnInit(): void {
     this.dropdownService.getdropvalue().subscribe(result => {
       this.states = result.data
      })
  }

  getCities(stateid: number){
    this.dropdownService.hitdropvalue(stateid).subscribe( results => {
      this.cities = results.data
      this.disableSCity= false;
    })
  }
}
