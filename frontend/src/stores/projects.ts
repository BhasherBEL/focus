import axios from "axios";
import { writable } from "svelte/store";

export const projects = getProjects();

const backend = 'http://127.0.0.1:3000'

interface Project {
    id: number | undefined,
    title: string,
}

function getProjects() {
    const { subscribe, set, update } = writable([] as Project[]);

    return {
        subscribe,
        init: async () => {
            const response = await axios.get(`${backend}/api/projects`);

            if(response.status < 303) {
                const data: Project[] = response.data;

                set(data);

                return data;
            }
        },
        add: async (project: Project) => {
            const response = await axios.post(`${backend}/api/project`, project);

            if(response.status < 303) {
                project.id = response.data["id"];
                update((oldProjects) => {
                    oldProjects.push(project)
                    return oldProjects;
                });
            }
        },
        edit: async (project: Project) => {
            const response = await axios.put(`${backend}/api/project/${project.id}`, project)
        

            if(response.status < 303) {
                update((oldProjects: Project[]) => {
                    for(let p of oldProjects){
                        if(p.id === project.id){
                            p.title = project.title;
                        }
                    }

                    return oldProjects;
                });
            }
        }
    }
}