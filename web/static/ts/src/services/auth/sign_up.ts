import axios, { AxiosError } from 'axios';

export class SignUp {
	form = $('#signUpForm');
	boxerr = $('#signUpBoxerr');
	constructor() {
		this.form.submit((e) => this.send(e));
	}
	send(e: JQuery.SubmitEvent) {
		const returnTo = new URLSearchParams(window.location.search).get(
			'return_to'
		);
		const data = new FormData(e.target);
		const url = returnTo ? `/users?return_to=${returnTo}` : '/users';
		axios
			.post(url, data, {
				headers: {
					'Content-Type': 'multipart/form-data',
				},
			})
			.then((r) => (window.location.href = r.data['return_to'] || '/'))
			.catch((e: AxiosError) => {
				console.log(e.response);
				this.handleError(e.response.data['error']);
			});

		e.preventDefault();
	}
	handleError(message: string) {
		this.boxerr.html(`<p>${message}</p>`);
	}
}
