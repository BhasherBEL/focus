<script lang="ts">
	import ModalTags from './modal_tags.svelte';
	import type { Card } from '../stores/interfaces';
	import api, { processError } from '../utils/api';
	import status from '../utils/status';

	export let show: boolean;
	export let card: Card;
	export let onCancel: () => void;
	export let onDelete: () => void;

	let tempCard: Card = { ...card };

	async function save(closeModal: boolean = true) {
		if (
			card.project_id != tempCard.project_id ||
			card.title !== tempCard.title ||
			card.content !== tempCard.content
		) {
			const response = await api.put(`/v1/cards/${card.id}`, {
				project_id: tempCard.project_id,
				title: tempCard.title,
				content: tempCard.content
			});

			if (response.status !== status.NoContent) {
				processError(response, 'Failed to update card');
				return;
			}

			card = { ...tempCard };
		}
		if (closeModal) show = false;
	}
</script>

{#if show}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="modal" on:click={() => save(true)}>
		<div class="content" on:click|stopPropagation>
			<div class="header">
				<input class="title" bind:value={tempCard.title} on:blur={() => save(false)} />
				<div class="buttons">
					<button on:click={onDelete}>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
							fill="none"
							stroke="white"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
						>
							<path d="M3 6h18"></path>
							<path
								d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m2-2h10a2 2 0 0 1 2 2v2H5V6a2 2 0 0 1 2-2z"
							></path>
							<line x1="10" y1="11" x2="10" y2="17"></line>
							<line x1="14" y1="11" x2="14" y2="17"></line>
						</svg>
					</button>
					<button on:click={onCancel}>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="white"
							viewBox="0 0 24 24"
						>
							<path d="M0 0h24v24H0z" fill="none" />
							<path
								d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 12 12z"
							/>
						</svg>
					</button>
				</div>
			</div>
			<div class="tags">
				<ModalTags bind:card />
			</div>
			<div class="body">
				<textarea
					bind:value={tempCard.content}
					placeholder="Add a description"
					on:blur={() => save(false)}
				/>
			</div>
		</div>
	</div>
{/if}
