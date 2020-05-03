import { teamsRequest } from './ApiServices';

export const getAllTeamsAPI = async() => {
    const res = await teamsRequest.get("/")
    return res;
}