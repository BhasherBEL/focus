import Project, { projects } from '$lib/types/Project';
import { hasPendingRequests } from '$lib/utils/api';
import { toastWarning } from '$lib/utils/toasts';
import { get } from 'svelte/store';

let socket: WebSocket;

export function connectWebSocket() {
	socket = new WebSocket('ws://localhost:3000/api/v1/ws');

	socket.onopen = () => {
		console.log('WebSocket connected');
	};

	socket.onclose = () => {
		console.log('WebSocket disconnected');
		toastWarning(
			'WebSocket disconnected',
			'You may experience sync issues. You can try to reload the page.'
		);

		const reconnectTimer = setTimeout(() => {
			connectWebSocket();
			clearTimeout(reconnectTimer);
		}, 5000);
	};

	socket.onerror = (err) => {
		console.error('WebSocket error:', err);
	};

	socket.onmessage = async (event) => {
		const parsed = JSON.parse(event.data);
		while (hasPendingRequests()) {
			await new Promise((resolve) => setTimeout(resolve, 100));
		}
		applyMessage(parsed);
	};
}

function applyMessage(data: any) {
	switch (data.object) {
		case 'project':
			applyProject(data);
			break;
		default:
			console.log('Unknown object:', data);
	}
}

function applyProject(data: any) {
	switch (data.action) {
		case 'create':
			if (get(projects).find((p) => p.id === data.id)) break;
		case 'update':
			Project.parse(data.value);
			break;
		case 'delete':
			projects.set(get(projects).filter((p) => p.id !== data.id));
			break;
	}
}
