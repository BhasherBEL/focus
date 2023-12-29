<script lang="ts">
	import axios from 'axios';
	import ModalTags from './modal_tags.svelte';
	import type { Card } from '../stores/interfaces';
	import { backend } from '../stores/config';

	export let show: boolean;
	export let card: Card;
	export let onCancel: () => void;

	let tempCard: Card = { ...card };

	function save(closeModal: boolean = true) {
		if (card != tempCard) {
			axios.put(`${backend}/api/card/${card.id}`, {
				project_id: tempCard.project_id,
				title: tempCard.title,
				content: tempCard.content
			});

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
				<ModalTags {tempCard} />
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
