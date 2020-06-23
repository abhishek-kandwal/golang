import { Component, OnInit } from '@angular/core';
import { UserdetailsService } from '../../services/userdetails.service';

export interface userDetail{
  ID : number,
  Name : string,
  Email: string,
  Dob: Date 
}

@Component({
  selector: 'app-userdetails',
  templateUrl: './userdetails.component.html',
  styleUrls: ['./userdetails.component.css']
})
export class UserdetailsComponent implements OnInit {
  dataSource : userDetail[];
  displayedColumns: string[] = ['ID', 'Name', 'Email', 'Dob'];
  constructor(private userdetailService : UserdetailsService) { }

  ngOnInit(): void {
    this.userdetailService.getUserDetails().subscribe( results => {
      this.dataSource = results.userdata;
    });
  }

}
