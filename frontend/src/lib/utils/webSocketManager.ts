import Card, { cards } from '$lib/types/Card';
import Project, { projects } from '$lib/types/Project';
import ProjectTag, { projectTags } from '$lib/types/ProjectTag';
import View, { views } from '$lib/types/View';
import { getBackendWsUrl, hasPendingRequests, randomID } from '$lib/utils/api';
import { toastAlert, toastWarning } from '$lib/utils/toasts';
import { get } from 'svelte/store';

export default class WebSocketManager {
	_socket: WebSocket | null;
	_queue: any[];
	_reconnectAttempts: number;
	_maxReconnectAttempts: number;
	_hasBeenConnected: boolean;

	constructor() {
		this._socket = null;
		this._queue = [];
		this._reconnectAttempts = 0;
		this._maxReconnectAttempts = 5;
		this._hasBeenConnected = false;
	}

	connect() {
		if (
			this._socket &&
			(this._socket.readyState === WebSocket.OPEN ||
				this._socket.readyState === WebSocket.CONNECTING)
		) {
			return;
		}

		this._socket = new WebSocket(getBackendWsUrl() + '/api/v1/ws');

		this._socket.onopen = () => {
			this._reconnectAttempts = 0;
			this._hasBeenConnected = true;
			console.log('WebSocket connected');
		};

		this._socket.onclose = () => {
			console.log('WebSocket disconnected');

			if (this._reconnectAttempts < this._maxReconnectAttempts) {
				toastWarning(
					'WebSocket disconnected',
					`You may experience sync issues. Trying to reconnect... (${this._reconnectAttempts + 1}/${
						this._maxReconnectAttempts
					})`
				);
				this._reconnectAttempts++;
				setTimeout(() => {
					this._socket?.close();
					this.connect();
				}, 5000);
			} else {
				toastAlert('Failed to connect to WebSocket', 'Please refresh the page to try again.');
			}
		};

		this._socket.onerror = (err) => {
			console.warn('WebSocket error:', err);
		};

		this._socket.onmessage = async (event) => {
			const data = JSON.parse(event.data);
			const source = data['X-Request-Source'];
			// console.log(source);
			if (!source || source === randomID) {
				return;
			}
			while (hasPendingRequests()) {
				await new Promise((resolve) => setTimeout(resolve, 100));
			}
			applyMessage(data);
		};
	}
}

function applyMessage(data: any) {
	switch (data.object) {
		case 'project':
			applyProject(data);
			break;
		case 'card':
			applyCard(data);
			break;
		case 'view':
			applyView(data);
			break;
		case 'projectTag':
			applyProjectTag(data);
			break;
		default:
			console.log('Unknown object:', data);
	}
}

function applyProject(data: any) {
	switch (data.action) {
		case 'create':
			Project.parse(data.data);
		case 'update':
			get(projects)
				.find((p) => p.id === data.id)
				?.updateFromDict(data.changes);
			projects.reload();
			break;
		case 'delete':
			projects.set(get(projects).filter((p) => p.id !== data.id));
			break;
	}
}

function applyCard(data: any) {
	switch (data.action) {
		case 'create':
			Card.parse(data.data);
			break;
		case 'update':
			get(cards)
				.find((c) => c.id === data.id)
				?.updateFromDict(data.changes);
			cards.reload();
			break;
		case 'delete':
			cards.set(get(cards).filter((c) => c.id !== data.id));
			break;
	}
}

function applyView(data: any) {
	switch (data.action) {
		case 'create':
		case 'update':
			View.parse(data.value);
			break;
		case 'delete':
			views.set(get(views).filter((v) => v.id !== data.id));
			break;
	}
}

function applyProjectTag(data: any) {
	switch (data.action) {
		case 'create':
		case 'update':
			ProjectTag.parse(data.value);
			break;
		case 'delete':
			projectTags.set(get(projectTags).filter((t) => t.id !== data.id));
			break;
	}
}
