import axios from 'axios';

export const login = payload => axios.post('http://localhost:5000/api/v1/auth/login', {
  Username: payload.userName,
  Password: payload.password,
});
export const validate = token => axios.post('http://localhost:5000/api/v1/auth/google', {
  token,
});
export const getToken = () => new Promise((resolve, reject) => {
  const token = localStorage.getItem('accesToken');
  if (!token) reject(new Error('No access token found'));
  resolve(token);
});

export const login2 = payload => axios.post('http://localhost:8080/api/v1/admin/page', { Data: payload });
