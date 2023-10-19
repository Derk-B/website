<script>
    import { fly } from 'svelte/transition';
	import { cubicOut } from "svelte/easing";
	import { onMount } from 'svelte';

    let options = {duration: 1000, easing: cubicOut, y: -80};
	let pageLoaded = false;

	onMount(() => {
		pageLoaded = true;
	});
</script>

<div class="flex flex-col justify-center text-center h-full">
    {#if pageLoaded}
    <div in:fly={{...options}}>
        <div class="flex justify-center m-5">
            <img width={100} src="404-error.png" alt="An icon indicating that this is a 404 error"/>
        </div>
        <h1 class="text-4xl font-bold mb-3 gradient-heading">Error</h1>
        <h2 class="text-4xl mb-3 gradient-heading">page not found</h2>
    </div>
    {/if}
</div>


<div class="iconref__bar">
    <a href="https://www.flaticon.com/free-icons/webpage" title="webpage icons">Webpage icons created by Freepik - Flaticon</a>
</div>

<style>
    .iconref__bar {
        font-size: small;
        margin-top: auto;
        padding: 5px;
        background-color: rgba(0, 0, 0, 0.171)
    }
</style>
