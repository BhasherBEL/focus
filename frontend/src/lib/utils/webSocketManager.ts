import Card, { cards } from '$lib/types/Card';
import CardTag from '$lib/types/CardTag';
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
	console.debug('WebSocket message:', data);
	if (!data.object) {
		toastWarning('Failed to parse WebSocket message: missing object');
		return;
	}

	if (!data.action) {
		toastWarning('Failed to parse WebSocket message: missing action');
		return;
	}

	if (data.object === 'project') applyProject(data);
	else if (data.object === 'cardTag') applyCardTag(data);
	else if (data.object === 'card') applyCard(data);
	else if (data.object === 'view') applyView(data);
	else if (data.object === 'projectTag') applyProjectTag(data);
	else if (data.object === 'filter') applyFilter(data);
	else toastWarning('Failed to parse WebSocket message: unknown object');
}

function applyProject(data: any) {
	if (data.action === 'create') {
		Project.parse(data.data);
		return;
	}
	if (data.action === 'delete') {
		Project.parseDelete(data.id);
		return;
	}

	const project = Project.fromId(data.id);
	if (!project) {
		toastWarning('Failed to parse project update: project not found');
		return;
	}

	if (data.action !== 'update') {
		toastWarning('Failed to parse project update: unknown action');
		return;
	}

	project.parseUpdate(data.changes);
	projects.reload();
}

function applyCard(data: any) {
	if (data.action === 'create') {
		Card.parse(data.data);
		return;
	}
	if (data.action === 'delete') {
		Card.parseDelete(data.id);
		return;
	}

	const card = Card.fromId(data.id);
	if (!card) {
		toastWarning('Failed to parse card update: card not found');
		return;
	}

	if (data.action !== 'update') {
		toastWarning('Failed to parse card update: unknown action');
		return;
	}

	card.parseUpdate(data.changes);
	cards.reload();
}

function applyCardTag(data: any) {
	const card = Card.fromId(data.card_id);
	if (!card) {
		toastWarning('Failed to parse card tag update: card not found');
		return;
	}

	if (data.action === 'create') {
		card.parseTag(data.data);
	} else if (data.action === 'update') {
		card.parseTagUpdate(data.changes);
	} else if (data.action === 'delete') {
		card.parseTagDelete(data.tag_id);
	} else {
		toastWarning('Failed to parse card tag update: unknown action');
		return;
	}

	cards.reload();
}

function applyView(data: any) {}

function applyProjectTag(data: any) {}

function applyFilter(data: any) {}
