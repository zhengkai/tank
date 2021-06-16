import { Component } from '@angular/core';
import { environment }  from '../../environments/environment';

@Component({
	selector: 'app-root',
	templateUrl: './bootstrap.component.html',
	styleUrls: ['./bootstrap.component.scss']
})
export class BootstrapComponent {
	title = 'Tank';
	prod = false;

	constructor() {
		this.prod = environment.production;
	}
}
