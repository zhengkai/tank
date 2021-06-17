import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { ApiService, TankMap  } from '../common/api.service';
import { pb } from '../../pb';
import * as d3 from 'd3';

export interface Data {
	nx: number;
	ny: number;
	type: number;
	id: number;
	x: number;
	y: number;
}

@Component({
	selector: 'app-matrix',
	templateUrl: './matrix.component.html',
	styleUrls: ['./matrix.component.scss']
})
export class MatrixComponent implements OnChanges {

	@Input() list: Data[] = [];

	size = 700;
	border = 50;

	src: pb.ITank[] = [];
	srcMap: TankMap = {};

	tank: any = {};

	svg: d3.Selection<SVGSVGElement, any, any, any>|null = null;

	constructor(
		public api: ApiService,
	) {
		(async () => {
			this.srcMap = await api.data();
		})();
	}

	drawInit() {
		if (this.svg) {
			return;
		}

		const svg = d3.select('#matrix').append('svg');
		this.svg = svg;

		this.svg.attr('height', this.size + this.border * 3)
		.attr('width', this.size + this.border * 2);
	}

	ngOnChanges() {
		this.drawInit();
		this.draw(this.list);
	}

	demo() {
		const list: Data[] = [];
		this.src.map(v => {
			const st = v?.statsHigher;
			if (v?.base?.tier !== 10 && v?.base?.tier !== 9) {
				return;
			}
			if (v?.base?.shop === 2 || !v?.base?.ID || !st?.battles || !st.damageDealt || !st.wins || !st.survivedBattles) {
				return
			}
			list.push({
				nx: st.wins / st.battles,
				ny: st.damageDealt / st.battles,
				type: v?.base?.type || 0,
				id: v.base.ID,
				x: 0,
				y: 0,
			})
		});

		this.draw(list);
	}

	draw(list: Data[]) {

		const svg = this.svg;
		if (!svg) {
			return;
		}

		const xList = list.map(v => v.nx);
		const yList = list.map(v => v.ny);

		const xMin = Math.min(...xList);
		const xMax = Math.max(...xList);
		const xRange = xMax - xMin;

		const yMin = Math.min(...yList);
		const yMax = Math.max(...yList);
		const yRange = yMax - yMin;

		list.forEach(v => {
			v.x =  this.border + (v.nx - xMin) / xRange * this.size;
			v.y =  this.border + (yMax - v.ny) / yRange * this.size;
		});

		const radius = 6;

		const self = this;

		svg.selectAll('*').remove();

		svg.selectAll('circle')
		.data(list)
		.enter()
		.append('svg:circle')
		.attr('cx', v => v.x)
		.attr('cy', v => v.y)
		.attr('r', radius)
		.attr('fill', v => d3.schemeCategory10[v.type])
		.on("mouseover", function (v, tank) {
			d3.select(this).attr('r', radius * 2);
			const t = self.srcMap[tank.id]
			if (t?.base) {
				self.tank = t.base;
			}
		})
		.on("mouseout", function (v, tank) {
			d3.select(this).attr('r', radius);
			// self.tank = {};
		})

		const x = d3.scaleLinear().range([this.border, this.border + this.size]).domain([xMin, xMax])
		const ax = d3.axisBottom(x).ticks(12);
		svg.append('g').call(ax)

		const y = d3.scaleLinear().range([this.border, this.border + this.size]).domain([yMax, yMin])
		const ay = d3.axisRight(y).ticks(12);
		svg.append('g').call(ay);
	}
}
