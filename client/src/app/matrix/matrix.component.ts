import { Component, OnInit } from '@angular/core';
import { ApiService } from '../common/api.service';
import { pb } from '../../pb';
import * as d3 from 'd3';

@Component({
	selector: 'app-matrix',
	templateUrl: './matrix.component.html',
	styleUrls: ['./matrix.component.scss']
})
export class MatrixComponent implements OnInit {

	src: pb.ITank[] = [];

	tank: any = {};

	constructor(
		public api: ApiService,
	) {
	}

	async ngOnInit(): Promise<void> {
		this.src = await this.api.list();
		this.draw();
	}

	draw() {

		const size = 700;
		const border = 50;

		const svg = d3.select('#matrix').append('svg');

		const list = this.src.map(v => {
			const st = v?.statsHigher;
			if ((v?.base?.tier || 0) < 9 || v?.base?.type === 0 ||  !st?.battles || !st.damageDealt || !st.wins) {
				return {
					base: {},
					dmg: 0,
					win: 0,
					ty: 0,
				}
			}
			return {
				base: v?.base || {},
				dmg: st.damageDealt / st.battles,
				win: st.wins / st.battles,
				ty: v?.base?.type || 0,
			}
		}).filter(v => !!v.dmg);

		const winList = list.map(v => v.win);
		const dmgList = list.map(v => v.dmg);

		const winMin = Math.min(...winList);
		const winMax = Math.max(...winList);
		const winRange = winMax - winMin;

		const dmgMin = Math.min(...dmgList);
		const dmgMax = Math.max(...dmgList);
		const dmgRange = dmgMax - dmgMin;

		console.log([winMin, winMin + winRange]);

		const radius = 10;

		const self = this;

		svg.attr('height', size + border * 3)
			.attr('width', size + border * 2)
			// .style('background-color', '#eee')

		svg.selectAll('circle')
			.data(list)
			.enter()
			.append('svg:circle')
			.attr('cx', (v) => {
				return border + (v.win - winMin) / winRange * size;
			})
			.attr('cy', (v) => {
				return border + (dmgMax - v.dmg) / dmgRange * size;
			})
			.attr('r', radius)
			.attr('fill', v => d3.schemeCategory10[v.ty])
            .on("mouseover", function (v, tank) {
            	d3.select(this).attr('r', radius * 2);
				self.tank = tank.base;
			})
            .on("mouseout", function (v, tank) {
            	d3.select(this).attr('r', radius);
				self.tank = {};
			})

		const x = d3.scaleLinear().range([border, border + size]).domain([winMin, winMax])
		const ax = d3.axisBottom(x).ticks(12);
		svg.append('g').call(ax)

		const y = d3.scaleLinear().range([border, border + size]).domain([dmgMax, dmgMin])
		const ay = d3.axisRight(y).ticks(12);
		svg.append('g').call(ay);

		console.log(svg);
	}
}
