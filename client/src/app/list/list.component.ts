import { Component, OnInit } from '@angular/core';
import { ApiService } from '../common/api.service';
import { pb } from '../../pb';

@Component({
	selector: 'app-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

	li: pb.ITank[] = [];

	constructor(
		public api: ApiService,
	) {
	}

	async ngOnInit(): Promise<void> {
		this.li = await this.api.list();
		this.li.sort((a: pb.ITank, b: pb.ITank) => {
			const at = a?.base?.tier || 1;
			const bt = b?.base?.tier || 1;
			return bt - at;
		})
	}

	getDmg(t: pb.ITankStats|undefined|null): number {
		if (!t) {
			return 0;
		}
		const d = t?.damageDealt;
		if (!d) {
			return 0;
		}

		const b = t?.battles;
		if (!b) {
			return 0;
		}

		return Math.round(d / b);
	}
}
