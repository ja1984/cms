import axios from 'axios';

export const createPage = payload => axios.post('http://localhost:5000/api/v1/admin/page', { Data: payload });

export const getPages = () => axios.get('http://localhost:5000/api/v1/admin/pages');
