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
	@Input() name = false;

	@ViewChild('box') box!: ElementRef;

	ready = false;
	lastDot: any = null;

	size = 700;
	border = 50;
	radius = 6;

	srcMap: TankMap = {};

	tank: any = {};

	svg: d3.Selection<SVGSVGElement, any, any, any>|null = null;

	constructor(
		public api: ApiService,
	) {
		(async () => {
			this.srcMap = await api.list();
			if (!this.ready) {
				this.ngOnChanges();
			}
		})();
	}

	ngAfterViewInit() {
		if (this.ready) {
			return;
		}
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

		this.svg = d3.select(this.box.nativeElement).append('svg');
		this.svg.attr('height', this.size + this.border * 2)
		.attr('width', this.size + this.border * 3);
	}

	ngOnChanges() {
		if (!this.list.length || !Object.keys(this.srcMap).length) {
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

		if (!this.svg) {
			return;
		}

		this.svg.selectAll('*').remove();
		this.resetDot();

		this.calcData(list);

		const li = this.svg.selectAll('circle')
		.data(list)
		.enter()

		if (this.name) {
			li.append('text')
			.attr('x', v => v.x + 8)
			.attr('y', v => v.y + 4)
			.style('font-size', 12)
			.style('fill', 'gray')
			.style('pointer-events', 'none')
			.text(d => {
				const t = this.srcMap[d.id]?.base;
				if (!t) {
					return '';
				}
				return t.name || '';
			});
		}

		const self = this;

		li.append('svg:circle')
		.attr('cx', v => v.x)
		.attr('cy', v => v.y)
		.attr('r', this.radius)
		.attr('fill', v => d3.schemeCategory10[v.type])
		.attr('stroke', 'white')
		.on('mouseover', function (ev, tank) {

			self.resetDot();

			const dot = d3.select(this)
			self.lastDot = dot;
			dot.attr('r', self.radius * 2);
			const t = self.srcMap[tank.id]
			if (t?.base) {
				self.tank = t.base;
			}
		});
	}

	calcData(list: Data[]) {

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

		const x = d3.scaleLinear().range([this.border, this.border + this.size]).domain([xMin, xMax])
		const ax = d3.axisBottom(x).ticks(12);
		this.svg?.append('g').call(ax);

		const y = d3.scaleLinear().range([this.border, this.border + this.size]).domain([yMax, yMin])
		const ay = d3.axisRight(y).ticks(12);
		this.svg?.append('g').call(ay);
	}
}
