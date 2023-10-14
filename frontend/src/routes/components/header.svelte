<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import SvelteLogo from '~icons/logos/svelte-icon'
</script>

<AppBar>
	<svelte:fragment slot="lead">
        <SvelteLogo/>
    </svelte:fragment>
	<h1 class="text-2xl">Derk Blom</h1>
	<svelte:fragment slot="trail">(actions)</svelte:fragment>
</AppBar>