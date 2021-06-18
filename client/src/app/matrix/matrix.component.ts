import { Component, OnChanges, Input, ElementRef, ViewChild } from '@angular/core';
import { ApiService, TankMap } from '../common/api.service';
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

	@ViewChild('box') box!: ElementRef;

	ready = false;
	lastDot: any = null;

	size = 700;
	border = 50;
	radius = 6;

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

	ngAfterViewInit() {
		if (this.ready) {
			return;
		}
		this.ready = true;
		this.ngOnChanges();
	}

	drawInit() {
		if (this.svg) {
			return;
		}
		if (!this.box) {
			return;
		}
		this.ready = true;

		const svg = d3.select(this.box.nativeElement).append('svg');
		this.svg = svg;

		this.svg.attr('height', this.size + this.border * 3)
		.attr('width', this.size + this.border * 2);
	}

	ngOnChanges() {
		if (!this.list.length) {
			return;
		}
		this.drawInit();
		this.draw(this.list);
	}

	resetDot() {
		const dot = this.lastDot;
		if (!dot) {
			return;
		}
		dot.attr('r', this.radius);
		this.tank = {};
		this.lastDot = null;
	}

	draw(list: Data[]) {
		if (!this.box) {
			return;
		}

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
		self.resetDot();

		svg.selectAll('circle')
		.data(list)
		.enter()
		.append('svg:circle')
		.attr('cx', v => v.x)
		.attr('cy', v => v.y)
		.attr('r', this.radius)
		.attr('fill', v => d3.schemeCategory10[v.type])
		.on('mouseover', function (v, tank) {

			self.resetDot();

			const dot = d3.select(this)
			self.lastDot = dot;
			dot.attr('r', self.radius * 2);
			const t = self.srcMap[tank.id]
			if (t?.base) {
				self.tank = t.base;
			}
		})

		const x = d3.scaleLinear().range([this.border, this.border + this.size]).domain([xMin, xMax])
		const ax = d3.axisBottom(x).ticks(12);
		svg.append('g').call(ax)

		const y = d3.scaleLinear().range([this.border, this.border + this.size]).domain([yMax, yMin])
		const ay = d3.axisRight(y).ticks(12);
		svg.append('g').call(ay);
	}
}
