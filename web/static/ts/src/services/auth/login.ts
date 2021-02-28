import axios, { AxiosError, AxiosResponse } from 'axios';

export class Login {
	form = $('#loginForm');
	boxerr = $('#loginBoxerr');
	constructor() {
		this.form.submit((e) => this.send(e));
	}
	send(e: JQuery.SubmitEvent) {
		e.preventDefault();
		const returnTo = new URLSearchParams(window.location.search).get(
			'return_to'
		);
		const data = new FormData(e.target);
		const url = returnTo ? `/sessions?return_to=${returnTo}` : '/sessions';
		axios
			.post(url, data, {
				headers: {
					'Content-Type': 'multipart/form-data',
				},
			})
			.then((r: AxiosResponse) => {
				window.location.href = r.data['return_to'] || '/';
			})
			.catch((e: AxiosError) => {
				console.log(e.response);
				this.handleError(e.response.data['error']);
			});
	}
	handleError(message: string) {
		this.boxerr.html(`<p>${message}</p>`);
	}
}
