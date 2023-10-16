<script>
	import { fly } from 'svelte/transition';
	import { cubicOut } from "svelte/easing";
	import { onMount } from 'svelte';

	let options = {duration: 1000, easing: cubicOut, y: 80};
	let pageLoaded = false;

	onMount(() => {
		pageLoaded = true;
	});
</script>

{#if pageLoaded}
<div in:fly={{...options}} class="h-full flex items-center justify-center">
	
	<div class="w-fit text-center">
		<h1 class="text-4xl font-bold mb-3 bg-gradient-to-br from-red-500 to-yellow-500 bg-clip-text text-transparent box-decoration-clone">Hello, my name is Derk!</h1>
		<h2 class="text-2xl">A Dutch software developer and computer science student</h2>
		<div class="mt-5">
			<button type="button" class="btn variant-filled-primary m-5 w-32">Projects</button>
			<button type="button" class="btn variant-filled-primary m-5 w-32">Blog</button>
		</div>
	</div>
	<div class="background_anim"/>
</div>


{/if}

<style>
	@keyframes background_anim {
		0% { transform: rotate(30deg); }
		25% { transform: rotate(80deg); }
		50% { transform: rotate(120deg); }
		75% { transform: rotate(200deg); }
		100% { transform: rotate(270deg); }
	}

	.background_anim {
		position: absolute;
		z-index: -1;
		background-color: rgba(86, 86, 233, 0.153);
		width: 200px;
		height: 200px;
		transform: rotate(30deg);
		animation-name: background_anim;
		animation-duration: 10s;
		animation-iteration-count: infinite;
		animation-timing-function: linear;
		animation-direction: alternate;
	}
</style>
