import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

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
	],
	providers: [],
	bootstrap: [
		BootstrapComponent,
	],
})
export class AppModule { }
