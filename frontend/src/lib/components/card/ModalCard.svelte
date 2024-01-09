<script lang="ts">
	import CloseIcon from '$lib/components/icons/CloseIcon.svelte';
	import TrashIcon from '$lib/components/icons/TrashIcon.svelte';
	import ModalTags from '$lib/components/tags/ModalTags.svelte';
	import currentModalCard from '$lib/stores/currentModalCard';
	import type Card from '$lib/types/Card';
	import { cards } from '$lib/types/Card';

	export let card: Card;

	let newTitle = card.title;
	let newContent = card.content;

	async function save(closeModal: boolean = true) {
		if (card.title !== newTitle || card.content !== newContent) {
			if (!(await card.update(newTitle, newContent))) return;
		}
		if (closeModal) currentModalCard.set(null);

		cards.reload();
	}
</script>

{#if $currentModalCard == card.id}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="modal" on:click={() => save(true)}>
		<div class="content" on:click|stopPropagation>
			<div class="header">
				<input class="title" bind:value={newTitle} on:blur={() => save(false)} />
				<div class="buttons">
					<button on:click={() => card.delete()}>
						<TrashIcon />
					</button>
					<button on:click={() => currentModalCard.set(null)}>
						<CloseIcon />
					</button>
				</div>
			</div>
			<div class="tags">
				<ModalTags {card} />
			</div>
			<div class="body">
				<textarea
					bind:value={newContent}
					placeholder="Add a description"
					on:blur={() => save(false)}
				/>
			</div>
		</div>
	</div>
{/if}

<style lang="less">
	.modal {
		background-color: rgba(0, 0, 0, 0.8);
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;

		.content {
			background: #1e1e1e;
			padding: 20px;
			border-radius: 8px;
			max-width: 1000px;
			width: 100%;
			display: flex;
			justify-content: center;
			flex-direction: column;
			gap: 30px;
		}
	}

	.modal input,
	.modal textarea {
		background: none;
		color: inherit;
		border: 1px solid #333;
		border-radius: 7px;
		padding: 4px;
	}

	.modal .title {
		font-size: 24px;
		font-weight: bold;
		width: 100%;
	}

	.modal .header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.modal .buttons {
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.modal button {
		margin-left: 5px;
		height: 50px;
		width: 50px;
		background: none;
		border: none;
		border-radius: 10px;
	}

	.modal button:hover {
		background-color: #333;
		cursor: pointer;
	}

	.modal .buttons button:first-child:hover {
		background-color: #433;
	}

	.modal .body {
		margin-bottom: 20px;
	}

	.modal textarea {
		width: 100%;
		min-height: 200px;
		resize: vertical;
		font-family: inherit;
	}

	.modal td button {
		width: 30px;
		height: 30px;
	}
</style>
