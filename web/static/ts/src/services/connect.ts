import axios, { AxiosError, AxiosResponse } from 'axios';

export class ConnectSchool {
	status = $('#connectStatus');
	send(schoolID: string) {
		axios
			.post('/connect/school', {
				schoolID,
			})
			.then((r: AxiosResponse) => this.status.html(`<p>Connected</p>`))
			.catch((e: AxiosError) => {
				console.log(e.response);
				this.handleError(e.response.data['error']);
			});
	}
	handleError(message: string) {
		this.status.html(`<p>${message}</p>`);
	}
}
