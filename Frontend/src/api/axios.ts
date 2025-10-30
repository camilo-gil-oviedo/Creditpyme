import axios from 'axios';

// Aquí se configura la URL base del backend en Go
const api = axios.create({
  baseURL: 'http://localhost:8080', //  cambia el puerto si el backend usa otro
});

export default api;