import { Component } from '@angular/core';
import { environment }  from '../../environments/environment';
import { ApiService } from '../common/api.service';

@Component({
	selector: 'app-root',
	templateUrl: './bootstrap.component.html',
	styleUrls: ['./bootstrap.component.scss']
})
export class BootstrapComponent {
	title = 'Tank';
	prod = false;
	init = false;

	constructor(
		public api: ApiService,
	) {
		this.prod = environment.production;

		(async () => {
			await api.data();
			this.init = true;
		})();
	}
}
