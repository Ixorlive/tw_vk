// TODO: how to read it from enviroment?
// I tried using process.env but it didn't work
const AUTH_API_URL = "http://localhost:8082"
const NOTE_API_URL = "http://localhost:8081";

export const AUTH_API_LOGIN = `${AUTH_API_URL}/login`;
export const AUTH_API_TOKEN = `${AUTH_API_URL}/token`;
export const AUTH_API_REGISTER = `${AUTH_API_URL}/register`;

export const NOTE_API_GETALL = `${NOTE_API_URL}/notes`;
export const NOTE_API_CREATE = `${NOTE_API_URL}/notes`;
export const NOTE_API_EDIT = `${NOTE_API_URL}/notes`;
export const NOTE_API_DELETE = `${NOTE_API_URL}/notes`;
