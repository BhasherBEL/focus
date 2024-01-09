<script lang="ts">
	import AddIcon from '$lib/components/icons/AddIcon.svelte';
	import Menu from '$lib/components/menu/Menu.svelte';
	import { tick } from 'svelte';
	import ModalTagTypes from './ModalTagTypes.svelte';
	import type Project from '$lib/types/Project';
	import ProjectTag from '$lib/types/ProjectTag';

	export let project: Project;

	let isOpen = false;

	let titleInput: HTMLInputElement;

	let title: string = '';
	let typeId: number = 0;

	async function openMenu() {
		isOpen = !isOpen;
		await tick();
		if (isOpen && titleInput) titleInput.focus();
	}

	async function createTag() {
		if (title == '') return;
		const res = await ProjectTag.create(project, title, typeId);

		if (!res) return;

		isOpen = false;
	}
</script>

<tr>
	<td>
		<button
			class="add-button"
			on:click={openMenu}
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					openMenu();
				}
			}}
		>
			<AddIcon />
		</button>
		<Menu
			{isOpen}
			onLeave={() => {
				isOpen = false;
			}}
		>
			<div class="menu-items">
				<input bind:this={titleInput} bind:value={title} />
				<ModalTagTypes type={typeId} onChoice={(id) => (typeId = id)} />
				<button on:click={createTag}>Create</button>
			</div>
		</Menu>
	</td>
</tr>

<style lang="less">
	.add-button {
		margin-top: 10px;
		padding: 3px 5px;
		border-radius: 5px;
		width: 100%;
		height: 100%;
	}

	.menu-items {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

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
	}
</style>
