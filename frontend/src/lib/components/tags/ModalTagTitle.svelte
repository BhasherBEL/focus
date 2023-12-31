<script lang="ts">
	import Menu from '$lib/components/menu/Menu.svelte';
	import { toastAlert } from '$lib/utils/toasts';
	import { tick } from 'svelte';
	import ModalTagTypes from './ModalTagTypes.svelte';
	import type ProjectTag from '$lib/types/ProjectTag';

	export let projectTag: ProjectTag;

	let askConfirm: boolean = false;

	let titleInput: HTMLInputElement;
	let isMenuOpen: boolean = false;
	let newTitle: string = projectTag.title;
	let newType: number = projectTag.type;

	async function openMenu() {
		isMenuOpen = !isMenuOpen;
		await tick();
		if (isMenuOpen && titleInput) titleInput.focus();
	}

	async function saveProjectTag() {
		if (newTitle === projectTag.title) return;

		if (newTitle === '') {
			toastAlert('Tag title cannot be empty');
			return;
		}

		return await projectTag.update(newTitle, newType);
	}
</script>

<td>
	<div
		class="title"
		on:click={openMenu}
		tabindex="0"
		role="button"
		on:keydown={(e) => {
			if (e.key === 'Enter') {
				openMenu();
			}
		}}
	>
		{projectTag.title}
	</div>
	<Menu bind:isOpen={isMenuOpen} onLeave={() => (askConfirm = false)}>
		<div class="menu-items">
			<input
				bind:this={titleInput}
				bind:value={newTitle}
				on:blur={saveProjectTag}
				on:keydown={(e) => {
					if (e.key === 'Enter') {
						saveProjectTag();
					}
				}}
			/>
			<ModalTagTypes
				type={newType}
				onChoice={async (id) => {
					saveProjectTag();
				}}
			/>
			{#if askConfirm}
				<div class="askconfirm">
					<span>Confirm?</span>
					<div>
						<button
							on:click={async () => {
								if (!(await projectTag.delete())) return;
								isMenuOpen = false;
							}}>✓</button
						>
						<button on:click={() => (askConfirm = false)}>✗</button>
					</div>
				</div>
			{:else}
				<button on:click={() => (askConfirm = true)}>Delete</button>
			{/if}
		</div>
	</Menu>
</td>

<style lang="less">
	td {
		min-width: 120px;
	}

	.title {
		padding: 3px 5px;
		margin: 2px;
		border-radius: 5px;
		height: 28px;

		&:hover {
			background-color: #fff2;
			cursor: pointer;
		}
	}

	input {
		margin-bottom: 5px;
		width: 70%;
		background-color: inherit;
		color: inherit;
		border: 1px solid #666;
		padding: 3px 5px;
		font-size: inherit;
		border-radius: 5px;
	}

	button {
		padding: 3px 5px;
		text-align: left;
		width: 80%;
	}

	.menu-items {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.askconfirm {
		width: 75%;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		text-align: center;

		div {
			display: flex;
			flex-direction: row;
			align-items: center;

			button {
				margin-left: 5px;
			}

			:first-child {
				background-color: #0f02;
			}

			:last-child {
				background-color: #f002;
			}
		}
	}
</style>
