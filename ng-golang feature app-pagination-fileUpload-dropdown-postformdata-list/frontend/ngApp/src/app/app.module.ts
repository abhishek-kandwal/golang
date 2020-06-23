import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AgmCoreModule } from '@agm/core';
import { MatGoogleMapsAutocompleteModule } from '@angular-material-extensions/google-maps-autocomplete';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatRadioModule } from '@angular/material/radio';
import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import {DemoMaterialModule} from './material.module';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { DropdownComponent } from './dropdown/dropdown.component';
import { FileUploadComponent } from './file-upload/file-upload.component';
import { UserdetailsComponent } from './userdetails/userdetails.component';
import { PostuserdetailsComponent } from './postuserdetails/postuserdetails.component';
import { ReactiveFormsModule } from '@angular/forms'; 


// const googleLoginOptions: LoginOpt = {
//   scope: 'profile email'
// };



@NgModule({
  declarations: [
    AppComponent,
    DropdownComponent,
    FileUploadComponent,
    UserdetailsComponent,
    PostuserdetailsComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatGoogleMapsAutocompleteModule,
    MatDatepickerModule,
    MatRadioModule,
    HttpClientModule,
    DemoMaterialModule,
    ReactiveFormsModule
  ],
  providers: [  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
