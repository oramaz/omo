import axios, { AxiosError, AxiosResponse } from 'axios';

export class Logout {
	btn = $('#logoutButton');
	constructor() {
		this.btn.click((e) => this.send(e));
	}
	send(e: JQuery.ClickEvent) {
		e.preventDefault();
		axios
			.delete('/sessions')
			.then((r: AxiosResponse) => (window.location.href = '/login'))
			.catch((e: AxiosError) => console.log(e));
	}
}
