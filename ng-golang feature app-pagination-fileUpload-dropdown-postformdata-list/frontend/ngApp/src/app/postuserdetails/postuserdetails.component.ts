import { Component, OnInit } from '@angular/core';
import { UserdetailsService } from '../../services/userdetails.service';
import { FormBuilder , FormGroup , Validators } from '@angular/forms';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-postuserdetails',
  templateUrl: './postuserdetails.component.html',
  styleUrls: ['./postuserdetails.component.css']
})
export class PostuserdetailsComponent implements OnInit {
  
  userForm : FormGroup;
  submitted = false;
  isExist: boolean = false;
  
  minDate = new Date(2000, 0, 1);
  maxDate = new Date(Date());

  constructor(
    private formBuilder: FormBuilder,
    private userdetailService : UserdetailsService,
    private _snackBar: MatSnackBar
    ) { }

  ngOnInit(){
    this.userForm = this.formBuilder.group({
      name: ['', Validators.required],
      email: ['', Validators.email],
      dob: ['', Validators.required]
    })
  }

  get f() {return this.userForm.controls;}

  onSubmit(){
    this.submitted = true;
    
    if( this.isExist){
      return
    }

    if(this.userForm.invalid){
      return;
    }

    this.userdetailService.postUserDetails(this.userForm.value).subscribe(Response => {
      console.log(Response);
      this.openSnackBar(Response.message , "Okay");
    })
  }

  checkEmail( value: Event){
    this.userdetailService.getvalidateEmail(value).subscribe(response => {
      this.isExist = response.isExist;
    })
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      duration: 2000,
    });
  }
}