import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './common/routing.module';
import { BootstrapComponent } from './common/bootstrap.component';
import { ListComponent } from './list/list.component';

@NgModule({
	declarations: [
		BootstrapComponent,
  ListComponent,
	],
	imports: [
		BrowserModule,
		AppRoutingModule,
    	HttpClientModule,
	],
	providers: [],
	bootstrap: [
		BootstrapComponent,
	],
})
export class AppModule { }
