import axios from 'axios';

export const login = payload => axios.post('http://localhost:8080/api/v1/auth/login', {
  Username: payload.userName,
  Password: payload.password,
});
export const login2 = payload => axios.post('http://localhost:8080/api/v1/admin/page', { Data: payload });
