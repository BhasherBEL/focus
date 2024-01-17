<script lang="ts">
	import CloseIcon from '$lib/components/icons/CloseIcon.svelte';
	import TrashIcon from '$lib/components/icons/TrashIcon.svelte';
	import ModalTags from '$lib/components/tags/ModalTags.svelte';
	import type Card from '$lib/types/Card';
	import { cards } from '$lib/types/Card';
	import SvelteMarkdown from 'svelte-markdown';

	export let card: Card;

	let newTitle: string = card.title;
	let newContent: string = card.content;
	let editDescription: boolean = false;

	async function save(closeModal: boolean = true) {
		if (card.title !== newTitle || card.content !== newContent) {
			if (!(await card.update(newTitle, newContent))) return;
		}
		if (closeModal) card.showModal = false;

		cards.reload();
	}
</script>

<svelte:window
	on:keydown|once={(e) => {
		if (e.key === 'Escape') return save(true);
	}}
/>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="modal" on:click|self={() => save(true)}>
	<div class="content">
		<div class="header">
			<input class="title" bind:value={newTitle} on:blur={() => save(false)} />
			<div class="buttons">
				<button
					on:click|once={async () => {
						await card.delete();
						card.showModal = false;
					}}
				>
					<TrashIcon />
				</button>
				<button
					on:click|once={() => {
						card.showModal = false;
						cards.reload();
					}}
				>
					<CloseIcon />
				</button>
			</div>
		</div>
		<div class="tags">
			<ModalTags {card} />
		</div>
		<div class="body">
			<div class="toggleEdit" on:click|preventDefault={() => (editDescription = !editDescription)}>
				👁
			</div>
			{#if editDescription}
				<textarea
					bind:value={newContent}
					placeholder="Add a description"
					on:blur={() => save(false)}
				/>
			{:else}
				<SvelteMarkdown source={card.content} />
			{/if}
		</div>
	</div>
</div>

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
		font-size: 2rem;
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
		// border: 1px solid red;
		font-size: 1.5rem;

		border: 1px solid #444;
		min-height: 300px;
		border-radius: 5px;
		padding: 5px;
		position: relative;

		&:before {
			transition: color 0.2s ease-in-out;
		}

		&:before:hover {
			color: #333;
		}
	}

	.toggleEdit {
		position: absolute;
		top: 10px;
		right: 10px;
		cursor: pointer;
	}

	.modal textarea {
		width: 100%;
		min-height: 300px;
		height: fit-content;
		border: 0;
		resize: none;
		font-family: inherit;
	}
</style>
