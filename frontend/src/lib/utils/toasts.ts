import { toast } from '@zerodevx/svelte-toast';

export function toastAlert(title: string, subtitle: string = '', persistant: boolean = false) {
    toast.push(
        `<strong>${title}</strong><br>${subtitle}`,
        {
            theme: {
                '--toastBackground': '#ff4d4f',
                '--toastBarBackground': '#d32f2f',
                '--toastColor': '#fff',
            },
            initial: persistant ? 0 : 1,
            next: 0,
            duration: 10000,
            pausable: true,
        },
    ) 
}