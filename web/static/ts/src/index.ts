import { Logout } from './services/auth/logout';
import { Login } from './services/auth/login';
import { SignUp } from './services/auth/sign_up';

$(document).ready(() => {
	const login = new Login();
	const signup = new SignUp();
	const logout = new Logout();
});
