import { Component, OnInit } from '@angular/core';
import {FileuploadService} from '../../services/fileupload.service';

export interface employeeInterface{
  id: number,
  name: string
}


@Component({
  selector: 'app-file-upload',
  templateUrl: './file-upload.component.html',
  styleUrls: ['./file-upload.component.css']
})


export class FileUploadComponent implements OnInit {
  fileToUpload: File = null;
  employeeList:employeeInterface;
  selectedID: number;
  isOptSetected: boolean = true;
  isfileSelected: boolean = true;
  fileStatus: any;

  constructor(private fileUplaadService:FileuploadService) { }

  ngOnInit(): void {
    this.fileUplaadService.getEmployeeList()
    .subscribe( result => {
      this.employeeList = result.employeeList
    });
  }

  handleFileInput(files: FileList) {
    this.fileToUpload = files.item(0);
    this.isFile()
}
selectedOptions(id: number){
  this.selectedID = id;
   this.isOptSetected = false;
}

isFile(){
  this.isfileSelected = false;
}

uploadFileToActivity() {
  this.fileUplaadService.postFile(this.fileToUpload , this.selectedID).subscribe(data => {
      this.fileStatus = data.isUpload;

    }, error => {
      console.log(error);
    });
}

}
