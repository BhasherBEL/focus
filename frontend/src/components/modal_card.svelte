<script lang="ts">
	import ModalTags from './modal_tags.svelte';
	import type { Card } from '../stores/interfaces';
	import api, { processError } from '../utils/api';
	import status from '../utils/status';
	import { cards, currentModalCard } from '../stores/smallStore';
	import TrashIcon from './icons/trashIcon.svelte';
	import CloseIcon from './icons/closeIcon.svelte';

	export let card: Card;

	let tempCard: Card = { ...card };

	async function save(closeModal: boolean = true) {
		if (
			card.projectId != tempCard.projectId ||
			card.title !== tempCard.title ||
			card.content !== tempCard.content
		) {
			const response = await api.put(`/v1/cards/${card.id}`, {
				project_id: tempCard.projectId,
				title: tempCard.title,
				content: tempCard.content
			});

			if (response.status !== status.NoContent) {
				processError(response, 'Failed to update card');
				return;
			}

			card = { ...tempCard };
		}
		if (closeModal) currentModalCard.set(-1);
	}
</script>

{#if $currentModalCard == card.id}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="modal" on:click={() => save(true)}>
		<div class="content" on:click|stopPropagation>
			<div class="header">
				<input class="title" bind:value={tempCard.title} on:blur={() => save(false)} />
				<div class="buttons">
					<button on:click={() => cards.remove(card)}>
						<TrashIcon />
					</button>
					<button on:click={() => currentModalCard.set(-1)}>
						<CloseIcon />
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
