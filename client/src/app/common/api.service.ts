import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { pb } from '../../pb';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  urlGateway = '/data/list.pb';

  constructor(
    private httpClient: HttpClient,
  ) {
  }

  async list(): Promise<pb.ITank[]> {

	const uri = this.urlGateway + '?' + (new Date()).toISOString().substring(0, 10);

    const post = this.httpClient.get(uri.replace(/-/g, ''), {
      observe: 'response',
      responseType: 'arraybuffer',
    }).toPromise();
    const res = await post;

    let rsp: pb.TankList;
    if (res.status !== 200 || !res.body) {
		return [];
    }

    const re = new Uint8Array(res.body);
    rsp = pb.TankList.decode(re);

    return rsp?.list || [];
  }
}
