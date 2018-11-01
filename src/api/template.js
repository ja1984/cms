import axios from 'axios';

export const createTemplate = payload => axios.post('http://localhost:5000/api/v1/admin/template', { Data: payload });

export const getTemplates = () => axios.get('http://localhost:5000/api/v1/admin/templates');
