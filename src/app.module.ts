
import {CommonModule} from '@angular/common';
import {HttpClientModule} from '@angular/common/http';
import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import * as AdditionService from 'angular_bazel_GRPC/addition_service/swagger_gen';
import {BASE_PATH} from 'angular_bazel_GRPC/addition_service/swagger_gen';

import {AppComponent} from './app.component';

@NgModule({
  imports: [
    AdditionService.ApiModule,
    CommonModule,
    BrowserModule,
    HttpClientModule,
  ],
  providers:
      [{provide: BASE_PATH, useValue: window.location.protocol + '//' + window.location.host}],
  declarations: [AppComponent],
  bootstrap: [AppComponent],
})
export class AppModule {
}
