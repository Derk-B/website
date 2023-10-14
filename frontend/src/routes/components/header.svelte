<script lang="ts">
    import { AppBar } from "@skeletonlabs/skeleton";
    import Terminal from '~icons/ic/outline-terminal'
</script>

<AppBar>
	<svelte:fragment slot="lead">
        <Terminal font-size="30"/>
    </svelte:fragment>
	<h1 class="text-2xl">Derk Blom</h1>
	<svelte:fragment slot="trail">
        <nav>
			<ul class="flex flex-row gap-3">
				<li><a href="/">Home</a></li>
				<li><a href="/about">About</a></li>
			</ul>
		</nav>
    </svelte:fragment>
</AppBar>