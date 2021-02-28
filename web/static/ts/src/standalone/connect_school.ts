import { ConnectSchool } from '../services/connect';

$(document).ready(() => {
	const connectSchool = new ConnectSchool();
	const schoolID = window.location.pathname.replace('/connect/school/', '');
	connectSchool.send(schoolID);
});
