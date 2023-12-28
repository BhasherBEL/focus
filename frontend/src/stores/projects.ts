import axios from "axios";
import { writable } from "svelte/store";

export const projects = getProjects();

const backend = 'http://127.0.0.1:3000'

interface Project {
    id: number,
    title: string,
}

function getProjects() {
    const { subscribe, set, update } = writable([]);

    return {
        subscribe,
        init: async () => {
            const response = await axios.get(`${backend}/api/projects`)
            console.log(response.data)

            if(response.status < 303) {
                const projects: Project[] = response.data;

                return projects;
            }
        }
    }
}